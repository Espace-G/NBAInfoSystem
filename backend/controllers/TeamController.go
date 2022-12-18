package controllers

import (
	"backend/models"
	"github.com/gin-gonic/gin"
	"os"
	"strconv"
)

func SelectTeams(ctx *gin.Context) {
	res := models.Dto{}
	s := ctx.Query("search")
	teams, err := models.SelectTeams(s)

	if err != nil {
		res.Status = 0
		res.Message = "Query Failed!!Try again"
	} else {
		res.Status = 1
		res.Message = "Query Success!!"
		res.Obj = teams
	}
	ctx.JSON(200, res)
}

func CreateTeam(ctx *gin.Context) {
	res := models.Dto{}
	team := models.TeamInfo{}
	err := ctx.ShouldBind(&team)

	if err != nil {
		res.Status = 0
		res.Message = "Bind Data Failed!! Try Again"
		ctx.JSON(200, res)
		return
	}
	fh, err2 := ctx.FormFile("img")
	if err2 != nil {
		res.Status = 0
		res.Message = "Upload Img Failed!! Try Again"
		ctx.JSON(200, res)
		return
	}
	filename := fh.Filename
	//filename = utils.GetOnlyFileName(filename)
	dst := "../frontend/public/img/team_img/" + filename
	err = ctx.SaveUploadedFile(fh, dst)
	if err != nil {
		res.Status = 0
		res.Message = "Save Img Failed!! Try Again"
		ctx.JSON(200, res)
		return
	}
	team.Logo = filename
	err = models.CreateTeam(&team)
	if err != nil {
		res.Status = 0
		res.Message = "Create Team Failed!!Try Again"
		ctx.JSON(200, res)
		return
	}
	res.Status = 1
	res.Message = "Create Success!"
	res.Obj = team
	ctx.JSON(200, res)
}

func SelectTeamById(ctx *gin.Context) {
	res := models.Dto{}
	s := ctx.Query("id")
	tid, _ := strconv.Atoi(s)

	team, err := models.SelectTeamById(tid)
	if err != nil {
		res.Status = 0
		res.Message = "Query Failed!! Try Again"
		ctx.JSON(200, res)
		return
	}
	res.Obj = team
	res.Status = 1
	res.Message = "Query Success!"
	ctx.JSON(200, res)
}

func UpdateTeam(ctx *gin.Context) {
	res := models.Dto{}
	team := models.TeamInfo{}
	err := ctx.ShouldBind(&team)
	if err != nil {
		res.Status = 0
		res.Message = "Bind Data Failed! Try Again"
		ctx.JSON(200, res)
		return
	}

	err = models.UpdateTeam(team)
	if err != nil {
		res.Status = 0
		res.Message = "Update Info Failed! Try Again"
		ctx.JSON(200, res)
		return
	}
	res.Status = 1
	res.Message = "Update Success!"
	ctx.JSON(200, res)
}

func RemoveTeam(ctx *gin.Context) {
	res := models.Dto{}
	s := ctx.Query("id")
	tid, _ := strconv.Atoi(s)

	team, err := models.SelectTeamById(tid)
	if err != nil {
		res.Status = 0
		res.Message = "Id error!"
		ctx.JSON(200, res)
		return
	}
	logoName := team.Logo
	if logoName != "" {
		dst := "../frontend/public/img/team_img/" + logoName
		err = os.Remove(dst)
		if err != nil {
			res.Status = 0
			res.Message = "Remove Logo Img Failed!"
			ctx.JSON(200, res)
			return
		}
	}

	err = models.RemoveTeam(tid)
	if err != nil {
		res.Status = 0
		res.Message = "Delete Data Failed!"
		ctx.JSON(200, res)
		return
	}

	res.Message = "Remove Team Success!"
	res.Status = 1
	ctx.JSON(200, res)
}
