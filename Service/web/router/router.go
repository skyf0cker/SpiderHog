package router

import (
	v1 "SpiderHog/Service/web/api/v1"
	"SpiderHog/Service/web/middleware"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.Default()
	r.Use(middleware.LoggerToFile())

	r.GET("api/v1/log", v1.Getlog)
	r.GET("api/v1/chan", v1.Querychan)

	return r
}
