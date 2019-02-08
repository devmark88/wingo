package utils

import (
	"io/ioutil"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"gitlab.com/mt-api/wingo/connectors"
	"gitlab.com/mt-api/wingo/context"
	"gitlab.com/mt-api/wingo/handlers"
	"gitlab.com/mt-api/wingo/logger"
)

//Start => Start Server
func Start(r *gin.Engine, cn *connectors.Connections) {
	ctx := context.AppContext{
		Connections: cn,
		AuthUser:    context.AuthUser{},
	}
	kp := viper.GetString("secrets.admin")
	s, e := ioutil.ReadFile(kp)
	logger.CheckOrFatal(e)
	rsaPub, e := jwt.ParseRSAPublicKeyFromPEM(s)
	logger.CheckOrFatal(e)
	ctx.AdminKey = rsaPub

	kp = viper.GetString("secrets.user")
	s, e = ioutil.ReadFile(kp)
	logger.CheckOrFatal(e)
	rsaPub, e = jwt.ParseRSAPublicKeyFromPEM(s)
	logger.CheckOrFatal(e)
	ctx.UserKey = rsaPub

	handlers.Setup(r, &ctx)
	p := viper.Get("server.address").(string)
	logger.Debug("server running on " + p)
	r.Run(p)
}
