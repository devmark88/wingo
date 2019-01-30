package middlewares

import (
	"github.com/gin-gonic/gin"
)

//ApplyGin => Apply gin middlewares
func ApplyGin(r *gin.Engine) {
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
}
