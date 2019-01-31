package connectors

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

// ConnectDatabase => Connect to Postgres Database
func ConnectDatabase(ctx *gin.Context) {
	h := viper.GetString("Database.host")
	p := viper.GetInt("database.port")
	pwd := viper.GetString("database.password")
	u := viper.GetString("database.username")
	db := viper.GetString("database.name")
	log.WithFields(log.Fields{
		"animal": "walrus",
	}).Info("A walrus appears")
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		h, p, u, pwd, db)
	ctx.Strin
}
