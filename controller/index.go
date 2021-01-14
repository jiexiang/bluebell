package controller

import (
	"bluebell/models"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Index(c *gin.Context) {
	user := sessions.Default(c).Get(UserInfo)
	if user, ok := user.(models.User); ok {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"login": true,
			"user":  user,
		})
		return
	}
	c.HTML(http.StatusOK, "index.html", nil)
}
