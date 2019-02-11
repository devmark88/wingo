package v1

import (
	"fmt"
	"net/http"
	"time"

	"github.com/spf13/viper"

	"gitlab.com/mt-api/wingo/helpers"

	"gitlab.com/mt-api/wingo/model"

	"gitlab.com/mt-api/wingo/response"

	"github.com/gin-gonic/gin"
	"gitlab.com/mt-api/wingo/repository"
)

func (h *V1Handlers) GetContestMeta(c *gin.Context) {
	r := repository.Connections{DB: h.Context.Connections.Database, Redis: h.Context.Connections.Cache, Queue: h.Context.Q.Server}
	c.Header("Content-Type", "application/json; charset=utf-8")
	err, m := r.GetMeta(true)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":  err.Error(),
			"status": http.StatusBadRequest,
		})
		return
	}
	err, u := r.GetUserInfo(h.Context.AuthUser.ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":  err.Error(),
			"status": http.StatusBadRequest,
		})
		return
	}
	res, err := createMetaResponse(m, u, h.Context.AuthUser.Username)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":  err.Error(),
			"status": http.StatusBadRequest,
		})
		return
	}
	c.JSON(http.StatusOK, res)
}
func createMetaResponse(m *[]model.ContestMeta, u *model.UserInfo, un string) (response.GetMetaResponse, error) {
	res := response.GetMetaResponse{}
	next := findNextContestInRange(*m)
	if next == nil {
		res.NextContestInHours = viper.GetString("app.next_contest.default_hour")
		res.NextContestInMinutes = viper.GetString("app.next_contest.default_min")
		res.NextContestInSeconds = viper.GetString("app.next_contest.default_sec")
	} else {
		h, m, s := helpers.GetTime(helpers.TimeInTehran(next.BeginTime))
		res.NextContestInHours = fmt.Sprintf("%v", h)
		res.NextContestInMinutes = fmt.Sprintf("%v", m)
		res.NextContestInSeconds = fmt.Sprintf("%v", s)
		res.Prize = next.Prize
		res.NextContestNeededTickets = next.NeededTickets
		res.Tickets = u.Tickets
		res.Correctors = u.Correctors
	}
	res.ShareData = response.ShareData{
		Title:   "بیا بازی کنیم",
		Message: fmt.Sprintf(`سلامَلِکُم – استاد چند روزه درگیره یه بازی شدم که حالم خرابه – میخوای حالت خراب بشه بکوف روی لینک زیر و با کد معرف %s ثبت نام کن `, un),
	}
	res.ShareData.DialogTitle = res.ShareData.Title
	res.Timeline = generateTimelineResponse(*m)
	return res, nil
}
func generateTimelineResponse(m []model.ContestMeta) []response.Timeline {
	var tl []response.Timeline
	idx := -1
	for i, c := range m {
		t := response.Timeline{}
		t.Currency = "تومان"
		t.Prize = c.Prize
		t.StartTime = helpers.TimeInTehran(c.BeginTime).String()
		t.Text = c.Title
		t.IsCurrent = false
		t.IsPast = helpers.IsPast(c.BeginTime)
		if t.IsPast == true {
			idx = i
		}
		tl = append(tl, t)
	}
	if idx != -1 {
		tl[idx+1].IsCurrent = true
	} else {
		tl[0].IsCurrent = true
	}
	return tl
}
func findNextContestInRange(m []model.ContestMeta) *model.ContestMeta {
	var next model.ContestMeta
	now := helpers.TimeInTehran(time.Now())
	for _, c := range m {
		if helpers.TimeInTehran(c.BeginTime).Sub(now) > 0 {
			next = c
			break
		}
	}
	if next.ID == 0 {
		return nil
	}
	return &next
}
