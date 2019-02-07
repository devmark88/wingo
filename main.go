package main

import (
	"github.com/gin-gonic/gin"
	"gitlab.com/mt-api/wingo/connectors"
	"gitlab.com/mt-api/wingo/middleware"
	"gitlab.com/mt-api/wingo/utils"
)

func main() {

	utils.InitConfig("config.yaml", "WINGO")
	r := gin.New()
	db := connectors.ConnectDatabase()
	cn := connectors.Connections{Database: db}
	middleware.ApplyGin(r)
	// r.Use(middleware.Errors())

	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	utils.Start(r, &cn)
	defer db.Close()
}
