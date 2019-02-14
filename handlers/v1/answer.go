package v1

import (
	"net/http"

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
	var track model.UserTrack
	track.UserID = uid
	if qidx == 0 {
		// it is the first question
	} else {

	}
	
	
}
