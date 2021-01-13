package controller

import (
	"bluebell/models"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
)

func Signup(c *gin.Context) {
	c.HTML(http.StatusOK, "signup.html", nil)
}

func DoSignup(c *gin.Context) {
	var signupParams models.ParamSignUp
	if err := c.ShouldBind(&signupParams); err != nil {
		errs, ok := err.(validator.ValidationErrors)
		if ok {
			err := removeTopStruct(errs.Translate(trans))
			c.HTML(http.StatusOK, "signup.html", gin.H{
				"error": true,
				"msg":   err,
			})
			//c.HTML(http.StatusOK, "signup.html", err)
			return
		}
		c.HTML(http.StatusOK, "signup.html", "系统错误，请稍后再试")
		return
	}
}
