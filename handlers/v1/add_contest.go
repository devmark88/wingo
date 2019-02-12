package v1

import (
	"net/http"

	"gitlab.com/mt-api/wingo/model"
	"gitlab.com/mt-api/wingo/repository"
	"gitlab.com/mt-api/wingo/response"

	"github.com/gin-gonic/gin"
	"gitlab.com/mt-api/wingo/request"
)

// AddMetaContest : Add meta data of contest
func (h *V1Handlers) AddMetaContest(c *gin.Context) {
	var m request.AddMetaRequest
	r := repository.Connections{DB: h.Context.Connections.Database, Redis: h.Context.Connections.Cache, Queue: h.Context.Q.Server}
	c.Header("Content-Type", "application/json; charset=utf-8")
	if err := c.Bind(&m); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":  err.Error(),
			"status": http.StatusBadRequest,
		})
		return
	}
	meta := m.ToModel()
	err := r.AddMeta(meta)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":  err.Error(),
			"status": http.StatusBadRequest,
		})
		return
	}

	c.JSON(http.StatusOK, mapMetaModelToResponse(meta))
}

// AttachQuestion : Add question and answer to the contest meta data
func (h *V1Handlers) AttachQuestion(c *gin.Context) {
	var m request.AttachQuestion
	r := repository.Connections{DB: h.Context.Connections.Database, Redis: h.Context.Connections.Cache, Queue: h.Context.Q.Server}
	c.Header("Content-Type", "application/json; charset=utf-8")
	if err := c.Bind(&m); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":  err.Error(),
			"status": http.StatusBadRequest,
		})
		return
	}
	md, err := m.ToModel()
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	if err = r.AddContest(md); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":  err.Error(),
			"status": http.StatusBadRequest,
		})
		return
	}

	c.JSON(http.StatusOK, mapContestToResponse(md))
}

func mapMetaModelToResponse(meta *model.ContestMeta) response.AddMeta {
	res := response.AddMeta{
		ID:                         meta.ID,
		AppID:                      meta.AppID,
		Title:                      meta.Title,
		Prize:                      meta.Prize,
		Duration:                   meta.Duration,
		NeededCorrectors:           meta.NeededCorrectors,
		AllowedCorrectorUsageTimes: meta.AllowedCorrectorUsageTimes,
		AllowCorrectTilQuestion:    meta.AllowCorrectTilQuestion,
		NeededTickets:              meta.NeededTickets,
	}
	res.BeginTime = meta.BeginTime.UTC().String()
	return res
}
func mapContestToResponse(contest *model.Contest) response.AttachQuestion {
	res := response.AttachQuestion{}
	res.Result = true
	return res
}
