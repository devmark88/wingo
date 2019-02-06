package v1
func AddMetaContest(r *gin.Engine) {
	
}

func Setup(r *gin.Engine, appCtx AppContext) {
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
}
