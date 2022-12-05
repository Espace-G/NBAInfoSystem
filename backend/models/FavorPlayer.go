package models

import "backend/dao"

type FavorPlayer struct {
	Uid int `json:"uid" gorm:"primary_key"`
	Pid int `json:"pid" gorm:"primary_key"`
}

// 插入一条收藏记录
func CPFavor(favor FavorPlayer) (err error) {
	err = dao.DB.Create(&favor).Error
	return
}

// 删除一条收藏记录
func DPFavor(favor FavorPlayer) (err error) {
	err = dao.DB.Delete(&favor).Error
	return
}

//查询某个用户的收藏球员列表
func QueryById(uid int) (playerSlice []PlayerInfo, err error) {
	err = dao.DB.Where("pid in (?)", dao.DB.Model(&FavorPlayer{}).Where("uid = ?", uid).Select("pid")).Find(&playerSlice).Error
	return
}
