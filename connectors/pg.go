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
	connStr := createConnectionString()
	logger.Debug("connecting to " + connStr)
	db, err := gorm.Open("postgres", connStr)
	logger.CheckOrFatal(err)
	logger.Debug("pinging database")
	err = db.DB().Ping()
	logger.CheckOrFatal(err)
	logger.Debug("connected to postgres > migrating...")
	migrate(db)
	logger.Debug("migrated")
	setLogger(db)
	return db
}
func createConnectionString() string {
	h := viper.GetString("database.host")
	p := viper.GetInt("database.port")
	pwd := viper.GetString("database.password")
	u := viper.GetString("database.user")
	dbn := viper.GetString("database.dbname")
	ssl := viper.GetString("database.ssl")

	return fmt.Sprintf("host=%s port=%v user=%s dbname=%s password=%s sslmode=%s", h, p, u, dbn, pwd, ssl)
}
func migrate(db *gorm.DB) {
	db.AutoMigrate(&model.ContestMeta{}, &model.Contest{}, &model.Question{}, &model.UserInfo{}, &model.UserTrack{})
	db.Model(&model.Question{}).AddForeignKey("contest_id", "contests(id)", "RESTRICT", "RESTRICT")
	db.Model(&model.Contest{}).AddForeignKey("contest_meta_id", "contest_meta(id)", "RESTRICT", "RESTRICT")
}
func setLogger(db *gorm.DB) {
	l := true
	if logger.LogLevel == "trace" || logger.LogLevel == "debug" {
		l = true
	}
	db.LogMode(l)
}
