package v1

import (
	"fmt"
	"net/http"
	"time"

	"gitlab.com/mt-api/wingo/helpers"

	"gitlab.com/mt-api/wingo/model"

	"gitlab.com/mt-api/wingo/response"

	"github.com/gin-gonic/gin"
	"gitlab.com/mt-api/wingo/repository"
)

func (h *V1Handlers) GetContestMeta(c *gin.Context) {
	r := repository.Connections{DB: h.Context.Connections.Database}
	c.Header("Content-Type", "application/json; charset=utf-8")
	err, m := r.GetMeta(true)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":  err.Error(),
			"status": http.StatusBadRequest,
		})
		return
	}
	
	_, err = createMetaResponse(m)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":  err.Error(),
			"status": http.StatusBadRequest,
		})
		return
	}
	c.JSON(http.StatusOK, m)
}
func createMetaResponse(m []*model.ContestMeta) (response.GetMetaResponse, error) {
	res := response.GetMetaResponse{}
	next := findNextContestInRange(m)
	if next == nil {
		res.NextContestInHours = "09"
		res.NextContestInMinutes = "30"
		res.NextContestInSeconds = "00"
	} else {
		h, m, s := helpers.GetTime(helpers.TimeInTehran(next.BeginTime))
		res.NextContestInHours = fmt.Sprintf("%v", h)
		res.NextContestInMinutes = fmt.Sprintf("%v", m)
		res.NextContestInSeconds = fmt.Sprintf("%v", s)
		res.Prize = next.Prize
		res.NextContestNeededTickets = next.NeededTickets
		// res.Tickets = next.NeededTickets
		// res.Correctors = next.NeededCorrectors
	}

	return res, nil
}
func findNextContestInRange(m []*model.ContestMeta) *model.ContestMeta {
	var next model.ContestMeta
	now := helpers.TimeInTehran(time.Now())
	for _, c := range m {
		if helpers.TimeInTehran(c.BeginTime).Sub(now) > 0 {
			next = *c
			break
		}
	}
	if next.ID == 0 {
		return nil
	}
	return next
}
