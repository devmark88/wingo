package v1

import (
	"net/http"

	"github.com/spf13/viper"

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
	code, err := updateUserInfo(h.Context.AuthUser.ID, m.Tickets, m.Correctors, &r)
	if err != nil {
		c.JSON(code, gin.H{
			"error":  err.Error(),
			"status": code,
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"error": nil,
		"data":  true,
	})
}

// AddReferral : add referral to the user
func (h *Handlers) AddReferral(c *gin.Context) {
	r := repository.Connections{DB: h.Context.Connections.Database, Redis: h.Context.Connections.Cache, Queue: h.Context.Q.Server}
	referralTicket := viper.GetInt("app.referralTicket")
	referralCorrector := viper.GetInt("app.referralTicket")
	code, err := updateUserInfo(h.Context.AuthUser.ID, uint(referralTicket), uint(referralCorrector), &r)
	if err != nil {
		c.JSON(code, gin.H{
			"error":  err.Error(),
			"status": code,
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"error": nil,
		"data":  true,
	})
}

func updateUserInfo(uID string, tickets, correctors uint, r *repository.Connections) (int, error) {
	uinfo, err := r.GetUserInfo(uID)
	if err != nil {
		return http.StatusBadRequest, err
	}
	if tickets > 0 {
		uinfo.Tickets = tickets
		uinfo.CanPlay = true
	}
	if correctors > 0 {
		uinfo.Correctors = correctors
	}
	err = r.AddUserInfo(uinfo)
	if err != nil {
		return http.StatusInternalServerError, nil
	}
	return 200, nil
}
