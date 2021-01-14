package mysql

import "errors"

var (
	ErrorUserExist       = errors.New("用户已经存在")
	ErrorPasswordInvalid = errors.New("密码错误")
)
