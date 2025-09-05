package models

import (
	"base-service/app/dao"
)

type AuthorityModel struct {
	base
}

// Create 添加数据
func (m AuthorityModel) Create(data *dao.Authority) error {
	ret := db.Debug().Table(m.table).Create(&data)
	return ret.Error
}

func (m AuthorityModel) TableName() string {
	return m.baseTable() + "Authority"
}

func (m AuthorityModel) Init() AuthorityModel {
	m.table = m.TableName()
	return m
}
