package v1

import (
	"net/http"

	"gitlab.com/mt-api/wingo/repository"

	"github.com/gin-gonic/gin"
	"gitlab.com/mt-api/wingo/context"
	"gitlab.com/mt-api/wingo/request"
)

type V1Handlers struct {
	Context *context.AppContext
}

func (h *V1Handlers) AddMetaContest(c *gin.Context) {
	var m request.AddMetaRequest
	r := repository.Connections{DB: h.Context.Connections.Database}
	c.Header("Content-Type", "application/json; charset=utf-8")
	if err := c.Bind(&m); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":  err.Error(),
			"status": http.StatusBadRequest,
		})
		return
	}
	meta := m.ToModel()
	r.AddMeta(&meta)
	c.JSON(http.StatusOK, meta)
}
