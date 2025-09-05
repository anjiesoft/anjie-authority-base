package models

import (
	"base-service/app/dao"
)

type LoginLogModel struct {
	base
}

// Create 添加数据
func (m LoginLogModel) Create(data *dao.LoginLog) error {
	ret := db.Table(m.table).Create(&data)
	return ret.Error
}

func (m LoginLogModel) TableName() string {
	return m.baseTable() + "login_log"
}

func (m LoginLogModel) Init() LoginLogModel {
	m.table = m.TableName()
	return m
}
