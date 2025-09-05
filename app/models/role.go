package models

import (
	"base-service/app/dao"
)

type RoleModel struct {
	base
}

// Create 添加数据
//
//	user := dao.User{Openid: "323232"}
//	err := models.User{}.Init().Create(&user)
func (m RoleModel) Create(data *dao.Role) error {
	ret := db.Table(m.table).Create(&data)
	return ret.Error
}

func (m RoleModel) TableName() string {
	return m.baseTable() + "Role"
}

func (m RoleModel) Init() RoleModel {
	m.table = m.TableName()
	return m
}
