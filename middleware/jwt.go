package middleware

import (
	"crypto/rsa"
	"fmt"
	"net/http"
	"strings"

	"gitlab.com/mt-api/wingo/context"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type appClaim struct {
	ID          string     `json:"id"`
	Avatar      string     `json:"avatar"`
	AccessLevel string     `json:"accessLevel"`
	IssuedAt    float64    `json:"iat"`
	Audience    string     `json:"aud"`
	Issuer      string     `json:"iss"`
	Cellphone   string     `json:"sub"`
	Username    string     `json:"name"`
	App         labelClaim `json:"app"`
}
type labelClaim struct {
	EnglishName string `json:"enName"`
	PersianName string `json:"faName"`
}

func Auth(app *context.AppContext, rsaPub *rsa.PublicKey) gin.HandlerFunc {
	var cls appClaim
	return func(c *gin.Context) {
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

			claims := token.Claims.(jwt.MapClaims)
			err = extractClaims(claims, &cls)
			if err != nil {
				return nil, fmt.Errorf("Error while extracting claims: %v", err)
			}
			return rsaPub, nil
		})
		if err != nil {
			c.AbortWithError(401, err)
		} else {
			app.AuthUser.ID = cls.ID
			app.AuthUser.Avatar = cls.Avatar
			app.AuthUser.Cellphone = cls.Cellphone
			if len(cls.Username) > 0 {
				app.AuthUser.Username = cls.Username
			}
		}
	}
}

func extractClaims(claims jwt.MapClaims, cls *appClaim) error {
	var lbl labelClaim
	l := claims["app"].(map[string]interface{})
	lbl.EnglishName = l["enName"].(string)
	lbl.PersianName = l["faName"].(string)
	cls.ID = claims["id"].(string)
	cls.Avatar = claims["avatar"].(string)
	cls.AccessLevel = claims["accessLevel"].(string)
	cls.IssuedAt = claims["iat"].(float64)
	cls.Audience = claims["aud"].(string)
	cls.Issuer = claims["iss"].(string)
	cls.Cellphone = claims["sub"].(string)
	if claims["name"] != nil && claims["name"] != "unknown" {
		cls.Username = claims["name"].(string)
	}
	cls.App = lbl
	return nil
}
