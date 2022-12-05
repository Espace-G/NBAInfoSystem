package models

import "backend/dao"

// 定义用户对象
type UserInfo struct {
	Uid      int    `json:"uid" gorm:"primary_key"`
	Username string `json:"username"`
	Password string `json:"password"`
	Utype    int16  `json:"utype"`
	Headshot string `json:"headshot"`
}

//登录验证,根据用户名密码查询用户，返回用户对象，无则返回nil
func CheckLogin(username string, password string) UserInfo {
	loginUser := UserInfo{}
	dao.DB.Where("username = ? and password = ?", username, password).First(&loginUser)
	return loginUser
}

//插入用户信息
func InsertUser(insertUser UserInfo) (err error) {
	err = dao.DB.Create(&insertUser).Error
	return
}

//查询所有用户信息
func SelectAll() (userSlice []UserInfo, err error) {
	userSlice = make([]UserInfo, 0)
	err = dao.DB.Find(&userSlice).Error

	return
}
