package v1

import (
	"net/http"

	"gitlab.com/mt-api/wingo/helpers"

	"gitlab.com/mt-api/wingo/model"

	"github.com/gin-gonic/gin"
	"gitlab.com/mt-api/wingo/repository"
	"gitlab.com/mt-api/wingo/request"
)

// PostAnswer : an answer posted by clients
func (h *Handlers) PostAnswer(c *gin.Context) {
	var m request.PostAnswer
	r := repository.Connections{DB: h.Context.Connections.Database, Redis: h.Context.Connections.Cache, Queue: h.Context.Q.Server}
	c.Header("Content-Type", "application/json; charset=utf-8")
	if err := c.Bind(&m); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":  err.Error(),
			"status": http.StatusBadRequest,
		})
		return
	}
	uid := h.Context.AuthUser.ID

	contest, err := r.GetContest(m.ContestID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":  err.Error(),
			"status": http.StatusBadRequest,
		})
		return
	}
	if contest.IsPast() || !contest.IsQuestionInTime(m.QuestionID) {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":  err.Error(),
			"status": http.StatusPreconditionFailed,
		})
		return
	}
	qidx := contest.GetQuestionIndex(m.QuestionID)
	ca, err := helpers.StringToIntArray(contest.CorrectAnswersIndices, ",")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":  err.Error(),
			"status": http.StatusBadRequest,
		})
		return
	}
	correctAnswer := ca[qidx]
	var track model.UserTrack
	uinfo, err := r.GetUserInfo(uid)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":  err.Error(),
			"status": http.StatusBadRequest,
		})
		return
	}
	track.UserID = uid
	track.ContestID = contest.ID
	track.QuestionID = m.QuestionID
	track.QuestionIndex = qidx

	if qidx == 0 {
		// it is the first question
		if uinfo.CanPlay == false {
			c.JSON(http.StatusBadRequest, gin.H{
				"error":  "Cannot answer this question",
				"status": http.StatusPreconditionFailed,
			})
			return
		}
		if m.SelectedIndex == correctAnswer {
			track.CanPlay = true
			track.CanUseCorrector = true
			track.IsSelectCorrectAnswer = true
			track.State = model.PostAnswer
		} else {
			if contest.Meta.AllowedCorrectorUsageTimes > 0 && uinfo.Correctors >= contest.Meta.NeededCorrectors {
				// reduce user correctors
				uinfo.Correctors -= contest.Meta.NeededCorrectors
				r.AddUserInfo(uinfo)
				track.CanPlay = true
				track.IsSelectCorrectAnswer = false
				track.State = model.PostAnswer
				if contest.Meta.AllowedCorrectorUsageTimes > 1 {
					track.CanUseCorrector = true
				} else {
					track.CanUseCorrector = false
				}
			}
		}
	} else {
		tracks, err := r.GetUserTracks(uid, m.ContestID)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error":  err.Error(),
				"status": http.StatusBadRequest,
			})
			return
		}
		latestTrack := tracks[len(tracks)-1]

	}

}
func returnError(err error, code int) {

}
