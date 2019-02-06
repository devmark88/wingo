package utils

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"gitlab.com/mt-api/wingo/handlers"
)

//Start => Start Server
func Start(r *gin.Engine) {
	ctx := handlers.AppContext{}
	handlers.Setup(r, ctx)
	Debug("SERVER RUNNING")
	r.Run(viper.Get("server.address").(string))
}
