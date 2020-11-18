package routers

import (
	"goweixin/conf"
	"goweixin/middleware"
	"net/http"

	"github.com/gin-gonic/gin"
)

// InitRouter 路由初始华
func InitRouter() *gin.Engine {
	gin.SetMode(conf.App.RunMode)

	router := gin.New()
	router.Use(middleware.LoggerToFile())

	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"data": 111,
		})
	})

	return router
}
