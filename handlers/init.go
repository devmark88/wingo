package handlers

import (
	"github.com/gin-gonic/gin"
	"gitlab.com/mt-api/wingo/context"
	m "gitlab.com/mt-api/wingo/handlers/v1"
)

//Setup => Setup application handlers
func Setup(r *gin.Engine, appCtx *context.AppContext) {
	vh := m.V1Handlers{Context: appCtx}
	v1 := r.Group("/v1")

	// Admin Routes
	admin := v1.Group("/admin")
	{
		admin.POST("contest/meta", vh.AddMetaContest)
	}
}
