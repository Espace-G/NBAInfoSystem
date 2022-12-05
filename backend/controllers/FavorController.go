package controllers

import (
	"backend/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

// 收藏一个球员或一支球队
func CreateFavor(ctx *gin.Context) {
	res := models.Dto{}
	t := ctx.Query("type")
	if t == "player" {
		//球员记录
		favor := models.FavorPlayer{}
		s := ctx.Query("uid")
		uid, _ := strconv.Atoi(s)
		s = ctx.Query("id")
		pid, _ := strconv.Atoi(s)
		favor.Uid = uid
		favor.Pid = pid
		err := models.CPFavor(favor)
		if err != nil {
			res.Status = 0
			res.Message = "Create Player Favor Failed!"
			ctx.JSON(200, res)
		} else {
			res.Status = 1
			res.Message = "Success!"
			ctx.JSON(200, res)
		}
	} else if t == "team" {
		//球队记录
		favor := models.FavorTeam{}
		s := ctx.Query("uid")
		uid, _ := strconv.Atoi(s)
		s = ctx.Query("id")
		tid, _ := strconv.Atoi(s)
		favor.Uid = uid
		favor.Tid = tid
		err := models.CTFavor(favor)
		if err != nil {
			res.Status = 0
			res.Message = "Create Team Favor Failed!"
			ctx.JSON(200, res)
		} else {
			res.Status = 1
			res.Message = "Success!"
			ctx.JSON(200, res)
		}
	} else {
		res.Status = 0
		res.Message = "Error Parameter"
		ctx.JSON(200, res)
	}
}

// 取消收藏一个球员或一支球队
func DeleteFavor(ctx *gin.Context) {
	res := models.Dto{}
	t := ctx.Query("type")
	if t == "player" {
		//球员记录
		favor := models.FavorPlayer{}
		s := ctx.Query("uid")
		uid, _ := strconv.Atoi(s)
		s = ctx.Query("id")
		pid, _ := strconv.Atoi(s)
		favor.Uid = uid
		favor.Pid = pid
		err := models.DPFavor(favor)
		if err != nil {
			res.Status = 0
			res.Message = "Delete Player Favor Failed!"
			ctx.JSON(200, res)
		} else {
			res.Status = 1
			res.Message = "Success!"
			ctx.JSON(200, res)
		}
	} else if t == "team" {
		//球队记录
		favor := models.FavorTeam{}
		s := ctx.Query("uid")
		uid, _ := strconv.Atoi(s)
		s = ctx.Query("id")
		tid, _ := strconv.Atoi(s)
		favor.Uid = uid
		favor.Tid = tid
		err := models.DTFavor(favor)
		if err != nil {
			res.Status = 0
			res.Message = "Delete Team Favor Failed!"
			ctx.JSON(200, res)
		} else {
			res.Status = 1
			res.Message = "Success!"
			ctx.JSON(200, res)
		}
	} else {
		res.Status = 0
		res.Message = "Error Parameter"
		ctx.JSON(200, res)
	}
}

// 查询用户收藏的所有球员
func QueryFavorPlayerByUid(ctx *gin.Context) {
	s := ctx.Query("id")
	uid, _ := strconv.Atoi(s)
	playerSlice, err := models.QueryById(uid)
	res := models.Dto{}
	if err != nil {
		res.Status = 0
		res.Message = "Query Failed!Try Again"
	} else {
		res.Status = 1
		res.Message = "Success!"
		res.Obj = playerSlice
	}
	ctx.JSON(200, res)
}

// 查询用户收藏的所有球队
func QueryFavorTeamByUid(ctx *gin.Context) {
	s := ctx.Query("id")
	uid, _ := strconv.Atoi(s)
	teamSlice, err := models.QueryById2(uid)
	res := models.Dto{}
	if err != nil {
		res.Status = 0
		res.Message = "Query Failed!Try Again"
	} else {
		res.Status = 1
		res.Message = "Success!"
		res.Obj = teamSlice
	}
	ctx.JSON(200, res)
}
