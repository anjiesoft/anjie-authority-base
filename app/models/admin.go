package models

import (
	"base-service/app/dao"
)

type AdminModel struct {
	base
}

// Create 添加数据
func (m AdminModel) Create(data *dao.Admin) error {
	ret := db.Debug().Table(m.table).Create(&data)
	return ret.Error
}

func (m AdminModel) TableName() string {
	return m.baseTable() + "admin"
}

func (m AdminModel) Init() AdminModel {
	m.table = m.TableName()
	return m
}
