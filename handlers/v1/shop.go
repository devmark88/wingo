package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gitlab.com/mt-api/wingo/repository"
	"gitlab.com/mt-api/wingo/request"
)

// UpdateUserInfo : add ticket and corrector to user
func (h *Handlers) UpdateUserInfo(c *gin.Context) {
	var m request.UpdateUserInfoRequest
	r := repository.Connections{DB: h.Context.Connections.Database, Redis: h.Context.Connections.Cache, Queue: h.Context.Q.Server}
	c.Header("Content-Type", "application/json; charset=utf-8")
	if err := c.Bind(&m); err != nil {
		attachErrorToContext(c, err, http.StatusBadRequest)
		return
	}
	uinfo, err := r.GetUserInfo(h.Context.AuthUser.ID)
	if err != nil {
		attachErrorToContext(c, err, http.StatusBadRequest)
		return
	}
	if m.Tickets > 0 {
		uinfo.Tickets = m.Tickets
		uinfo.CanPlay = true
	}
	if m.Correctors > 0 {
		uinfo.Correctors = m.Correctors
	}
	err = r.AddUserInfo(uinfo)
	if err != nil {
		attachErrorToContext(c, err, http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"error": nil,
		"data":  true,
	})
}
