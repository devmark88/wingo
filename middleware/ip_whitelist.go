package middleware

import (
	"log"
	"net"
	"net/http"

	"github.com/spf13/viper"

	"github.com/gin-gonic/gin"
)

// IPCheck : check user ip is in valid range
func IPCheck() gin.HandlerFunc {
	return func(c *gin.Context) {
		ipstr := viper.GetString("server.cidr_ip_whitelist")
		if ipstr == "*" {
			c.Next()
		} else {
			cip := net.ParseIP(c.ClientIP())
			_, net, err := net.ParseCIDR(ipstr)
			if err != nil {
				log.Fatal(err)
			}
			if !net.Contains(cip) {
				c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
					"code":  http.StatusUnauthorized,
					"error": "You have not access to this resource",
				})
			}
		}
	}
}
