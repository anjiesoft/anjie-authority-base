package models

import (
	"base-service/app/dao"
)

type ProjectModel struct {
	base
}

// Create 添加数据
func (m ProjectModel) Create(data *dao.Project) error {
	ret := db.Table(m.table).Create(&data)
	return ret.Error
}

func (m ProjectModel) TableName() string {
	return m.baseTable() + "project"
}

func (m ProjectModel) Init() ProjectModel {
	m.table = m.TableName()
	return m
}
