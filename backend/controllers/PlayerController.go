package controllers

import (
	"backend/models"
	"backend/utils"
	"github.com/gin-gonic/gin"
	"os"
	"strconv"
)

func SelectPlayers(ctx *gin.Context) {
	res := models.Dto{}
	params := make(map[string]string, 2)
	orderby := ctx.Query("orderby")
	search := ctx.Query("search")
	if orderby != "" {
		params["orderby"] = orderby
	}
	if search != "" {
		params["search"] = search
	}
	players, err := models.SelectPlayers(params)
	if err != nil {
		res.Status = 0
		res.Message = "Failed!!"
	} else {
		res.Message = "Success!!"
		res.Status = 1
		res.Obj = players
	}
	ctx.JSON(200, res)
}

func SelectPlayerById(ctx *gin.Context) {
	res := models.Dto{}
	s := ctx.Query("id")
	pid, _ := strconv.Atoi(s)
	player, err := models.SelectPlayerById(pid)

	if err != nil {
		res.Message = "NO Player"
		res.Status = 0
	} else {
		res.Message = "Success!"
		res.Status = 1
		res.Obj = player
	}
	ctx.JSON(200, res)
}

func InsertPlayer(ctx *gin.Context) {
	res := models.Dto{}
	player := models.PlayerInfo{}
	err := ctx.ShouldBind(&player)
	if err != nil {
		res.Status = 0
		res.Message = "Create Player Failed!!Try Again"
		ctx.JSON(200, res)
		return
	}

	fh, err2 := ctx.FormFile("img")
	if err2 != nil {
		res.Status = 0
		res.Message = "Upload Img Failed!!Try Again"
		ctx.JSON(200, res)
		return
	}
	filename := fh.Filename
	filename = utils.GetOnlyFileName(filename)
	dst := "../frontend/public/img/player_img/" + filename
	err2 = ctx.SaveUploadedFile(fh, dst)
	if err2 != nil {
		res.Status = 0
		res.Message = "Save Img Failed!!Try Again"
		ctx.JSON(200, res)
		return
	}
	player.ImgPath = filename

	err2 = models.InsertPlayer(&player)
	if err2 != nil {
		res.Message = "Sava Data Failed!!Try Again"
		res.Status = 0
		ctx.JSON(200, res)
		return
	}
	res.Message = "Create Success!!"
	res.Status = 1
	ctx.JSON(200, res)
}

func Remove(ctx *gin.Context) {
	res := models.Dto{}
	s := ctx.Query("id")
	pid, _ := strconv.Atoi(s)
	//先删除照片文件
	player, _ := models.SelectPlayerById(pid)
	s2 := player.ImgPath
	if s2 != "" {
		filepath := "../frontend/public/img/player_img/" + s2
		err := os.Remove(filepath)
		if err != nil {
			res.Obj = 0
			res.Message = "Delete Img Failed! Try Again"
			res.Obj = err.Error()
			ctx.JSON(200, res)
			return
		}
	}

	err := models.RemovePlayer(pid)
	if err != nil {
		res.Status = 0
		res.Message = "Remove Failed!!"
	} else {
		res.Status = 1
		res.Message = "Remove Success!!"
	}
	ctx.JSON(200, res)
}

func UpdatePlayer(ctx *gin.Context) {
	res := models.Dto{}
	player := models.PlayerInfo{}
	err := ctx.ShouldBind(&player)
	if err != nil {
		res.Message = "Get Data Failed !!"
		res.Status = 0
		ctx.JSON(200, res)
		return
	}
	err = models.UpdatePlayer(player)
	if err != nil {
		res.Message = "Update Data Failed! Try Again"
		res.Status = 0
		res.Obj = err
		ctx.JSON(200, res)
		return
	}
	res.Message = "Update Success!"
	res.Status = 1
	res.Obj = player
	ctx.JSON(200, res)
}

func TeamPlayers(ctx *gin.Context) {
	res := models.Dto{}
	s := ctx.Query("id")
	tid, _ := strconv.Atoi(s)
	playerSlice, err := models.SelectPlayersInOneTeam(tid)
	if err != nil {
		res.Status = 0
		res.Message = "Query Info Failed! Try Again"
		ctx.JSON(200, res)
		return
	}
	res.Status = 1
	res.Message = "Success!"
	res.Obj = playerSlice
	ctx.JSON(200, res)
}
