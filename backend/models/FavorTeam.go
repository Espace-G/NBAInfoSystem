package models

import "backend/dao"

type FavorTeam struct {
	Uid int `json:"uid" gorm:"primary_key"`
	Tid int `json:"tid" gorm:"primary_key"`
}

// 插入一条收藏记录
func CTFavor(favor FavorTeam) (err error) {
	err = dao.DB.Create(&favor).Error
	return
}

// 删除一条收藏记录
func DTFavor(favor FavorTeam) (err error) {
	err = dao.DB.Delete(&favor).Error
	return
}

//查询某个用户的收藏球队列表
func QueryById2(uid int) (teamSlice []TeamInfo, err error) {
	//使用子查询
	//select * from team_infos
	//	where tid in (select tid from favor_team where uid = ?)
	err = dao.DB.Where("tid in (?)", dao.DB.Model(&FavorTeam{}).
		Where("uid = ?", uid).
		Select("tid")).
		Find(&teamSlice).
		Error
	return
}
