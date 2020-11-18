package main

import (
	"github.com/gin-gonic/gin"
	"goweixin/middleware"
	"net/http"
)

func main()  {
	engine := gin.Default()

	engine.Use(middleware.LoggerToFIle())

	engine.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"data": 111,
		})
	})
	engine.Run(":3000")
}
