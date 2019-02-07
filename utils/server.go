package utils

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"gitlab.com/mt-api/wingo/connectors"
	"gitlab.com/mt-api/wingo/handlers"
	"gitlab.com/mt-api/wingo/logger"
)

//Start => Start Server
func Start(r *gin.Engine, cn *connectors.Connections) {
	ctx := handlers.AppContext{}
	ctx.Connections = cn
	handlers.Setup(r, &ctx)
	p := viper.Get("server.address").(string)
	logger.Debug("server running on " + p)
	r.Run(p)
}
