package mysql

import "bluebell/models"

// Login 登录
func Login(username string) (user models.User, err error) {
	sql := `select user_id, username, password from user where username = ?`
	err = db.Get(&user, sql, username)
	return
}

// IsUserExits 检查用户是否存在
func IsUserExits(username string) (err error) {
	count := 0
	sql := `select count(username) from user where username = ?`
	err = db.Get(&count, sql, username)
	if err != nil {
		return
	}
	if count != 0 {
		err = ErrorUserExist
	}
	return
}

// InsertUser 新增用户
func InsertUser(userId int64, username, password string) (err error) {
	sql := `insert into user (user_id, username, password) values (?, ?, ?)`
	_, err = db.Exec(sql, userId, username, password)
	return
}
