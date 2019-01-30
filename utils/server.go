package utils

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"gitlab.com/mt-api/wingo/handlers"
)

//Start => Start Server
func Start(r *gin.Engine) {
	handlers.Setup(r)

	r.Run(viper.Get("server.address").(string))
}
