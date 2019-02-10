package utils

import (
	"io/ioutil"

	"gitlab.com/mt-api/wingo/q"

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
	ctx.Q = &context.QContext{
		Server:  nil,
		Workers: nil,
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

	srv, wrk, err := q.Start(1)
	if err != nil {
		logger.CheckOrFatal(err)
	}
	ctx.Q.Server = srv
	ctx.Q.Workers = wrk

	r.Run(p)
}
