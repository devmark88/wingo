package v1

import (
	"net/http"

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
	c.JSON(http.StatusOK, m)
}
