package router

import (
	"bluebell/controller"
	"bluebell/logger"
	"bluebell/models"
	"encoding/gob"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func InitRouter(mode string) *gin.Engine {
	if mode == gin.ReleaseMode {
		gin.SetMode(gin.ReleaseMode)
	}

	r := gin.New()

	// session
	gob.Register(models.User{})
	store := cookie.NewStore([]byte("bluebell-cookie"))
	r.Use(sessions.Sessions("mysession", store))

	r.Use(logger.GinLogger(), logger.GinRecovery(true))

	// 加载静态文件、网页
	r.LoadHTMLGlob("./templates/*")
	r.Static("/static", "./static")

	r.GET("/signup", controller.Signup)
	r.POST("/signup", controller.DoSignup)

	r.GET("/login", controller.Login)
	r.POST("/login", controller.DoLogin)

	r.GET("/", controller.Index)

	return r
}
