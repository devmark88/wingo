package context

import (
	"crypto/rsa"

	"gitlab.com/mt-api/wingo/connectors"
)

type AuthUser struct {
	ID     string `json:"id"`
	Avatar string `json:"avatar"`
	Cellphone string `json:"cellphone"`
	Username string `json:"username"`
}
type AppContext struct {
	Connections *connectors.Connections
	AuthUser    AuthUser
	AdminKey    *rsa.PublicKey
	UserKey     *rsa.PublicKey
}
