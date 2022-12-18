package controllers

import (
	"backend/models"
	"backend/utils"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
)

// 用户登录
func CheckLogin(ctx *gin.Context) {
	username := ctx.PostForm("username")
	password := ctx.PostForm("password")

	loginUser := models.CheckLogin(username, password)

	res := models.Dto{}
	//将用户ID存入cookie
	//name string, value string, maxAge int, path string, domain string, secure bool, httpOnly bool
	if loginUser.Uid != 0 {
		//查询到用户，登录正确
		ctx.SetCookie("userId", strconv.Itoa(loginUser.Uid), 3600, "/", "localhost", false, true)
		res.Message = "Login Success!"
		res.Status = 1
		res.Obj = loginUser
	} else {
		res.Message = "Wrong Username Or Password!"
		res.Status = 0
	}
	ctx.JSON(200, res)
}

// 用户退出登录
func Logout(ctx *gin.Context) {
	res := models.Dto{}
	s, err := ctx.Cookie("userId")
	if err != nil || s == "" {
		res.Status = 0
		res.Message = "Didn't Login"
		ctx.JSON(200, res)
	} else {
		ctx.SetCookie("userId", "", 0, "/", "localhost", false, true)
		res.Message = "Log out Success"
		res.Status = 1
		ctx.JSON(200, res)
	}
}

// 用户注册
func UserReg(ctx *gin.Context) {
	var res models.Dto
	//form 表单提交用户名密码头像文件
	username := ctx.PostForm("username")
	password := ctx.PostForm("password")
	fh, err := ctx.FormFile("headshot")
	if err != nil {
		res.Message = "Img Upload Failed!"
		res.Status = 0
		ctx.JSON(200, res)
		return
	}
	filename := fh.Filename
	filename = utils.GetOnlyFileName(filename)
	dst := "../frontend/public/img/headshot/" + filename
	err = ctx.SaveUploadedFile(fh, dst)
	if err != nil {
		res.Message = "Img Save Failed!"
		res.Status = 0
		ctx.JSON(200, res)
		return
	}
	regUser := models.UserInfo{
		Username: username,
		Password: password,
		Headshot: filename,
		Utype:    0,
	}
	err = models.InsertUser(regUser)
	if err != nil {
		res.Message = err.Error()
		res.Status = 0
		//删除刚刚存储的文件
		os.Remove(dst)
		ctx.JSON(200, res)
	} else {
		res.Message = "Register Success!!"
		res.Status = 1
		ctx.JSON(200, res)
	}

}

// 查询所有用户
func SelectAll(ctx *gin.Context) {
	userSlice, err := models.SelectAll()
	res := models.Dto{}
	if err != nil {
		res.Status = 0
		res.Message = "Select All User Failed!! Try Again"
		ctx.JSON(200, res)
	} else {
		res.Status = 1
		res.Message = "Success!"
		res.Obj = userSlice
		ctx.JSON(200, res)
	}
}

// 获取当前的登录用户
func GetLoginingUser(ctx *gin.Context) {
	res := models.Dto{}
	s, err := ctx.Cookie("userId")
	if err != nil {
		res.Status = 0
		res.Message = "No logining!"
		ctx.JSON(200, res)
		return
	}
	uid, _ := strconv.Atoi(s)

	login, err2 := models.QueryUserByID(uid)
	if err2 != nil {
		res.Message = "Query Error !"
		res.Status = 0
		ctx.JSON(200, res)
		return
	}

	//若当前有用户登录，延长userId的Cookie时间
	ctx.SetCookie("userId", strconv.Itoa(login.Uid), 3600, "/", "localhost", false, true)
	res.Obj = login
	res.Message = "Success"
	res.Status = 1
	ctx.JSON(200, res)
}
