package connectors

import (
	"fmt"

	"gitlab.com/mt-api/wingo/model"

	"gitlab.com/mt-api/wingo/logger"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"

	"github.com/spf13/viper"
)

// ConnectDatabase => Connect to Postgres Database
func ConnectDatabase() *gorm.DB {
	h := viper.GetString("database.host")
	p := viper.GetInt("database.port")
	pwd := viper.GetString("database.password")
	u := viper.GetString("database.user")
	dbn := viper.GetString("database.dbname")
	ssl := viper.GetString("database.ssl")

	connStr := fmt.Sprintf("host=%s port=%v user=%s dbname=%s password=%s sslmode=%s",
		h, p, u, dbn, pwd, ssl)
	logger.Debug("connecting to " + connStr)
	db, err := gorm.Open("postgres", connStr)
	logger.CheckOrFatal(err)
	logger.Debug("pinging database")
	err = db.DB().Ping()
	logger.CheckOrFatal(err)
	logger.Debug("connected to postgres")
	defer db.Close()
	db.AutoMigrate(&model.ContestMeta{}, &model.Contest{}, &model.Question{}, &model.UserInfo{}, &model.UserTrack{})
	logger.Debug("migrated")

	return db
}
