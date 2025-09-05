package models

import (
	"base-service/app/dao"
)

type ActionLogModel struct {
	base
}

// Create 添加数据
func (m ActionLogModel) Create(data *dao.ActionLog) error {
	ret := db.Table(m.table).Create(&data)
	return ret.Error
}

func (m ActionLogModel) TableName() string {
	return m.baseTable() + "action_log"
}

func (m ActionLogModel) Init() ActionLogModel {
	m.table = m.TableName()
	return m
}
