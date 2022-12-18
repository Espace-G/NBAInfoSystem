package router

import (
	"backend/controllers"
	"backend/models"
	"github.com/gin-gonic/gin"
)

func SetRouter() *gin.Engine {
	r := gin.Default()

	user := r.Group("user")

	{
		user.POST("checkLogin", controllers.CheckLogin)
		user.POST("register", controllers.UserReg)
		user.GET("selectAll", controllers.SelectAll)
		user.GET("getLogining", controllers.GetLoginingUser)
		user.GET("logout", controllers.Logout)
	}

	player := r.Group("player")
	{
		player.GET("selectPlayers", controllers.SelectPlayers)
		player.GET("selectPlayerById", controllers.SelectPlayerById)
		player.POST("insertPlayer", controllers.InsertPlayer)
		player.GET("remove", controllers.Remove)
		player.POST("update", controllers.UpdatePlayer)
		player.GET("teamPlayer", controllers.TeamPlayers)
	}

	team := r.Group("team")
	{
		team.GET("query", controllers.SelectTeams)
		team.POST("insertTeam", controllers.CreateTeam)
		team.GET("selectById", controllers.SelectTeamById)
		team.POST("updateTeam", controllers.UpdateTeam)
		team.GET("deleteTeam", controllers.RemoveTeam)

		team.POST("insertTeam2", func(ctx *gin.Context) {
			team := models.TeamInfo{}
			err := ctx.ShouldBind(&team)
			tname := ctx.PostForm("tname")
			if err != nil {
				ctx.String(200, err.Error())
				return
			}

			ctx.JSON(200, gin.H{
				"team name": tname,
				"obj":       team,
			})
		})
	}

	favor := r.Group("favor")
	{
		favor.GET("create", controllers.CreateFavor)
		favor.GET("undo", controllers.DeleteFavor)
		favor.GET("player", controllers.QueryFavorPlayerByUid)
		favor.GET("team", controllers.QueryFavorTeamByUid)
	}
	return r
}
