package connectors

import (
	"database/sql"
	"fmt"

	"gitlab.com/mt-api/wingo/utils"

	_ "github.com/lib/pq"

	"github.com/spf13/viper"
)

// ConnectDatabase => Connect to Postgres Database
func ConnectDatabase() *sql.DB {
	h := viper.GetString("Database.host")
	p := viper.GetInt("database.port")
	pwd := viper.GetString("database.password")
	u := viper.GetString("database.user")
	db := viper.GetString("database.name")

	connStr := fmt.Sprintf("postgres://%s:%s@%s:%v/%s?sslmode=disable",
		u, pwd, h, p, db)
	pdb, err := sql.Open("postgres", connStr)
	utils.Debug("Connecting to " + connStr)
	utils.CheckOrFatal(err)
	defer pdb.Close()

	err = pdb.Ping()
	utils.CheckOrFatal(err)
	utils.Debug("Connected to PSQL => " + connStr)
	return pdb
}
