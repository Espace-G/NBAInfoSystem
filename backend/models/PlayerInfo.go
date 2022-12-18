package models

import "backend/dao"

type PlayerInfo struct {
	Pid     int     `json:"pid" form:"pid" gorm:"primary_key"`
	CnName  string  `json:"cnname" form:"cnname"`
	EnName  string  `json:"enname" form:"enname"`
	Height  float32 `json:"height" form:"height"`
	Weight  float32 `json:"weight" form:"weight"`
	Nation  string  `json:"nation" form:"nation"`
	ImgPath string  `json:"imgpath" form:"imgpath"`
	Age     int16   `json:"age" form:"age"`
	Tid     int16   `json:"tid" form:"tid"`
}

func SelectPlayers(params map[string]string) (players []PlayerInfo, err error) {
	//查询有三种情况，无参数直接查询所有球员，
	//有orderby参数进行排序查询，有search参数进行字符串匹配查询
	if params["orderby"] == "" && params["search"] == "" {
		err = dao.DB.Find(&players).Error
	} else if params["orderby"] != "" {
		dao.DB.Order(params["orderby"]).Find(&players)
	} else {
		dao.DB.Where("cn_name like ? or en_name like ?",
			"%"+params["search"]+"%", "%"+params["search"]+"%").Find(&players)
	}
	return
}

//根据球员ID查询信息
func SelectPlayerById(pid int) (player PlayerInfo, err error) {
	err = dao.DB.Where("pid = ?", pid).First(&player).Error
	return
}

//插入新的球员
func InsertPlayer(player *PlayerInfo) (err error) {
	err = dao.DB.Create(player).Error
	return
}

//删除某个球员
func RemovePlayer(pid int) (err error) {
	err = dao.DB.Delete(&PlayerInfo{}, pid).Error
	return
}

func UpdatePlayer(player PlayerInfo) (err error) {
	err = dao.DB.Model(&PlayerInfo{}).Where("pid = ?", player.Pid).Updates(player).Error
	return
}

//查找在同一只球队的球员
func SelectPlayersInOneTeam(tid int) (playerSlice []PlayerInfo, err error) {
	err = dao.DB.Where("tid = ?", tid).Find(&playerSlice).Error
	return
}
