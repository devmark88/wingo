package middleware

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		s, e := ioutil.ReadFile("/Users/mark/Devs/me/moneyket/wingo/secrets/pub.user.pem")
		if e != nil {
			c.AbortWithError(http.StatusInternalServerError, fmt.Errorf("error in reading file."))
			return
		}
		rsaPub, e := jwt.ParseRSAPublicKeyFromPEM(s)
		if e != nil {
			c.AbortWithError(http.StatusInternalServerError, fmt.Errorf("Cannot parse pub key"))
			return
		}
		authHeader := strings.Split(c.Request.Header.Get("Authorization"), " ")
		if len(authHeader) != 2 {
			c.AbortWithError(http.StatusUnauthorized, fmt.Errorf("No Authorization found"))
			return
		}
		_, err := jwt.Parse(authHeader[1], func(token *jwt.Token) (interface{}, error) {
			method := jwt.GetSigningMethod("RS256")
			parts := strings.Split(authHeader[1], ".")
			err := method.Verify(strings.Join(parts[0:2], "."), parts[2], rsaPub)

			if err != nil {
				return nil, fmt.Errorf("Unexpected signing method: %v => %v", token.Header["alg"], err)
			}
			return rsaPub, nil
		})
		if err != nil {
			c.AbortWithError(401, err)
		}
	}
}
