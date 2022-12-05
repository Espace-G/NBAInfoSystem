package controllers

import (
	"backend/models"
	"github.com/gin-gonic/gin"
	"strconv"
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
		res.Message = "Register Failed!Try Again."
		res.Status = 0
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
