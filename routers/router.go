package routers

import (
	v1 "goweixin/api/v1"
	"goweixin/conf"
	"goweixin/middleware"

	"github.com/gin-gonic/gin"
)

// InitRouter 路由初始华
func InitRouter() *gin.Engine {
	gin.SetMode(conf.App.RunMode)

	router := gin.New()

	// 加载中间件
	router.Use(middleware.LoggerToFile())
	router.Use(middleware.Cors())

	router.GET("/", v1.Index)

	return router
}
