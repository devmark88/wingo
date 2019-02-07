package main

import (
	"github.com/gin-gonic/gin"
	"gitlab.com/mt-api/wingo/connectors"
	"gitlab.com/mt-api/wingo/middlewares"
	"gitlab.com/mt-api/wingo/utils"
)

func main() {

	utils.InitConfig("config.yaml", "WINGO")
	r := gin.New()
	db := connectors.ConnectDatabase()
	cn := connectors.Connections{db}
	middlewares.ApplyGin(r)

	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	utils.Start(r, &cn)
	// r := gin.Default()
	// r.GET("/ping", func(c *gin.Context) {
	// 	c.JSON(200, gin.H{
	// 		"message": "pong",
	// 	})
	// })
	// r.Run() // listen and serve on 0.0.0.0:8080
}
