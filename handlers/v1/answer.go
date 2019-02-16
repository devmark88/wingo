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
		attachErrorToContext(c, err, http.StatusBadRequest)
		return
	}
	uid := h.Context.AuthUser.ID

	contest, err := r.GetContest(m.ContestID)
	if err != nil {
		attachErrorToContext(c, err, http.StatusBadRequest)
		return
	}
	if contest.IsPast() || !contest.IsQuestionInTime(m.QuestionID) {
		attachErrorToContext(c, err, http.StatusPreconditionFailed)
		return
	}
	qidx := contest.GetQuestionIndex(m.QuestionID)
	if err != nil {
		attachErrorToContext(c, err, http.StatusBadRequest)
		return
	}
	var track model.UserTrack
	uinfo, err := r.GetUserInfo(uid)
	if err != nil {
		attachErrorToContext(c, err, http.StatusBadRequest)
		return
	}
	track.UserID = uid
	track.ContestID = contest.ID
	track.QuestionID = m.QuestionID
	track.QuestionIndex = qidx

	if contest.IsItFirstQuestion(qidx) {
		if uinfo.CanPlay == false {
			attachErrorToContext(c, err, http.StatusPreconditionFailed)
			return
		}
		if contest.IsItCorrectAnswer(m.SelectedIndex, m.QuestionID) {
			track.CanPlay = true
			track.CanUseCorrector = true
			track.IsSelectCorrectAnswer = true
			track.State = model.PostAnswer
		} else {
			if contest.CaneYetUserCorrector(uinfo, nil, qidx) {
				// reduce user correctors
				uinfo.Correctors -= contest.Meta.NeededCorrectors
				r.AddUserInfo(uinfo)
				track.CanPlay = true
				track.IsSelectCorrectAnswer = false
				track.CorrectorUsageTimes++
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
			attachErrorToContext(c, err, http.StatusBadRequest)
			return
		}
		cl := *tracks
		lst := cl[len(cl)-1]
		if lst.QuestionIndex+1 != qidx || lst.CanPlay == false {
			attachErrorToContext(c, err, http.StatusPreconditionFailed)
			return
		}
		if contest.IsItCorrectAnswer(m.SelectedIndex, m.QuestionID) {
			track.CanPlay = true
			track.IsSelectCorrectAnswer = true
			track.State = model.PostAnswer
			if contest.IsItLastQuestion(qidx) {
				track.State = model.WinTheGame
			}
		} else {
			if contest.CaneYetUserCorrector(nil, &lst, qidx) {
				track.CanPlay = true
				track.IsSelectCorrectAnswer = false
				track.CorrectorUsageTimes++
				track.State = model.PostAnswer
				if contest.Meta.AllowedCorrectorUsageTimes < track.CorrectorUsageTimes {
					track.CanUseCorrector = true
				} else {
					track.CanUseCorrector = false
				}
			}
		}
	}
	go r.SaveUserTrackAsync(&track)
	c.JSON(http.StatusOK, gin.H{
		"error": nil,
		"data":  true,
	})
}
func attachErrorToContext(c *gin.Context, err error, code int) {
	c.JSON(http.StatusBadRequest, gin.H{
		"error":  err.Error(),
		"status": code,
	})
	return
}
