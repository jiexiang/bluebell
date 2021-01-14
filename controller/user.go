package controller

import (
	"bluebell/dao/mysql"
	"bluebell/logic"
	"bluebell/models"
	"errors"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
	"net/http"
)

// Login 登录界面
func Login(c *gin.Context) {
	c.HTML(http.StatusOK, "login.html", nil)
}

// DoLogin 登录
func DoLogin(c *gin.Context) {
	// 参数校验
	var loginParams models.ParamLogin
	if err := c.ShouldBind(&loginParams); err != nil {
		zap.L().Error("DoLogin 参数错误", zap.Error(err))
		errs, ok := err.(validator.ValidationErrors)
		if ok {
			err := removeTopStruct(errs.Translate(trans))
			c.HTML(http.StatusOK, "login.html", gin.H{
				"error": true,
				"msg":   err,
			})
			return
		}
		c.HTML(http.StatusOK, "login.html", "系统错误，请稍后再试")
		return
	}

	// 登录逻辑
	user, err := logic.Login(loginParams)
	if err != nil {
		zap.L().Error("logic.Login", zap.Error(err))
		if errors.Is(err, mysql.ErrorPasswordInvalid) {
			c.HTML(http.StatusOK, "login.html", gin.H{
				"error": true,
				"msg":   err.Error(),
			})
			return
		}
		c.HTML(http.StatusOK, "login.html", gin.H{
			"error": true,
			"msg":   CodeServerBusy.Msg(),
		})
		return
	}

	// 保存用户信息
	session := sessions.Default(c)
	session.Set(UserInfo, user)
	session.Save()

	// 跳转首页
	c.Redirect(http.StatusMovedPermanently, "/")
}

// Signup 注册页面
func Signup(c *gin.Context) {
	c.HTML(http.StatusOK, "signup.html", nil)
}

// DoSignup 注册
func DoSignup(c *gin.Context) {
	// 参数校验
	var signupParams models.ParamSignUp
	if err := c.ShouldBind(&signupParams); err != nil {
		zap.L().Error("DoSignup 参数错误", zap.Error(err))
		errs, ok := err.(validator.ValidationErrors)
		if ok {
			err := removeTopStruct(errs.Translate(trans))
			c.HTML(http.StatusOK, "signup.html", gin.H{
				"error": true,
				"msg":   err,
			})
			return
		}
		c.HTML(http.StatusOK, "signup.html", "系统错误，请稍后再试")
		return
	}
	// 注册逻辑
	if err := logic.SignUp(signupParams); err != nil {
		zap.L().Error("logic.SignUp", zap.Error(err))
		if errors.Is(err, mysql.ErrorUserExist) {
			c.HTML(http.StatusOK, "signup.html", gin.H{
				"error": true,
				"msg":   err,
			})
			return
		}
		c.HTML(http.StatusOK, "signup.html", gin.H{
			"error": true,
			"msg":   CodeServerBusy.Msg(),
		})
		return
	}
	// 注册成功，跳转登录界面
	c.Redirect(http.StatusMovedPermanently, "/login")
}
