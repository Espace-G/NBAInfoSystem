package main

import (
	"backend/dao"
	"backend/router"
	"log"
)

func main() {
	//初始化gorm数据库连接
	err := dao.InitMysql()
	if err != nil {
		log.Fatal("Database Connect Failed")
		return
	}

	r := router.SetRouter()
	r.Run("localhost:8083")

}
