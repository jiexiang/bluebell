package logic

import (
	"bluebell/dao/mysql"
	"bluebell/models"
	"bluebell/pkg/snowflake"
	"crypto/md5"
	"encoding/hex"
)

const secret = "bluebell_password_md5"

// Login 登录逻辑
func Login(loginParams models.ParamLogin) (user models.User, err error) {
	// 登录
	user, err = mysql.Login(loginParams.Username)
	if err != nil {
		return
	}

	// 判断密码
	if user.Password != encryptPassword(loginParams.Password) {
		err = mysql.ErrorPasswordInvalid
		return
	}
	return
}

// SignUp 注册逻辑
func SignUp(signupParams models.ParamSignUp) (err error) {
	// 查看账号是否已经存在了
	if err = mysql.IsUserExits(signupParams.Username); err != nil {
		return
	}

	// 雪花id
	userId := snowflake.GenID()

	// 对密码进行加密
	password := encryptPassword(signupParams.Password)
	err = mysql.InsertUser(userId, signupParams.Username, password)
	return
}

// encryptPassword 密码加密
func encryptPassword(oPassword string) string {
	h := md5.New()
	h.Write([]byte(secret))
	return hex.EncodeToString(h.Sum([]byte(oPassword)))
}
