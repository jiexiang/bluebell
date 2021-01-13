package router

import (
	"bluebell/controller"
	"bluebell/logger"
	"github.com/gin-gonic/gin"
)

func InitRouter(mode string) *gin.Engine {
	if mode == gin.ReleaseMode {
		gin.SetMode(gin.ReleaseMode)
	}

	r := gin.New()
	r.Use(logger.GinLogger(), logger.GinRecovery(true))

	// 加载静态文件、网页
	r.LoadHTMLGlob("./templates/*")
	r.Static("/static", "./static")

	r.GET("/signup", controller.Signup)
	r.POST("/signup", controller.DoSignup)

	return r
}
