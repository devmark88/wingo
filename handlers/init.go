package handlers

import (
	"github.com/gin-gonic/gin"
	"gitlab.com/mt-api/wingo/connectors"
)

type AppContext struct {
	Connections *connectors.Connections
}

//Setup => Setup application handlers
func Setup(r *gin.Engine, appCtx *AppContext) {
}
