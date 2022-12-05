package router

import (
	"backend/controllers"

	"github.com/gin-gonic/gin"
)

func SetRouter() *gin.Engine {
	r := gin.Default()

	user := r.Group("user")

	{
		user.POST("checkLogin", controllers.CheckLogin)
		user.POST("register", controllers.UserReg)
		user.GET("selectAll", controllers.SelectAll)
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
