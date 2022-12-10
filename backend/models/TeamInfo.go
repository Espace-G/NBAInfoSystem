package models

import "backend/dao"

type TeamInfo struct {
	Tid   int    `json:"tid" gorm:"primary_key" form:"tid"`
	Tname string `json:"tname" form:"tname"`
	Tcity string `json:"tcity" form:"tcity"`
	Tzone string `json:"tzone" form:"tzone"`
	Arena string `json:"arena" form:"arena"`
	Logo  string `json:"logo" form:"logo_img"`
}

// 可选条件查询球队
func SelectTeams(search string) (teams []TeamInfo, err error) {
	if search == "" {
		err = dao.DB.Find(&teams).Error
	} else {
		err = dao.DB.Where("tname like ?", "%"+search+"%").Find(&teams).Error
	}
	return
}

func CreateTeam(team *TeamInfo) (err error) {
	err = dao.DB.Create(team).Error
	return
}

func SelectTeamById(tid int) (team TeamInfo, err error) {
	err = dao.DB.Where("tid = ?", tid).First(&team).Error
	return
}

func UpdateTeam(team TeamInfo) (err error) {
	err = dao.DB.Model(&TeamInfo{}).Where("tid = ?", team.Tid).Updates(team).Error
	return
}

func RemoveTeam(tid int) (err error) {
	err = dao.DB.Delete(&TeamInfo{}, tid).Error
	return
}
