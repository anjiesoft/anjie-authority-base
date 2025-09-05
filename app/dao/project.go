package dao

import "base-service/utils"

// Project 项目表
type Project struct {
	Id         int              `gorm:"column:id;type:int(11);primary_key;AUTO_INCREMENT" json:"id"`
	Name       string           `gorm:"column:name;type:varchar(255);comment:项目描述;NOT NULL" json:"name"`
	Status     uint8            `gorm:"column:status;type:tinyint(1);default:1;comment:状态，1为正常，2为禁用;NOT NULL" json:"status"`
	Remarks    string           `gorm:"column:remarks;type:varchar(255);default:'';comment:描述;NOT NULL" json:"remarks"`
	Logo       string           `gorm:"column:logo;type:varchar(255);default:'';comment:logo;NOT NULL" json:"logo"`
	Reason     string           `gorm:"column:reason;type:varchar(255);comment:原因" json:"reason"`
	AdminId    int              `gorm:"column:admin_id;type:int(11);comment:操作人员ID" json:"admin_id"`
	AdminName  string           `gorm:"column:admin_name;type:varchar(255);comment:操作人员姓名" json:"admin_name"`
	CreateTime utils.CustomTime `gorm:"column:create_time;type:datetime" json:"create_time"`
	UpdateTime utils.CustomTime `gorm:"column:update_time;type:datetime" json:"update_time"`
}

type ProjectItems struct {
	Items []Project `json:"items"`
	Count int64     `json:"count"`
}

// ProjectStatus 修改状态
type ProjectStatus struct {
	Id         int              `json:"id"`
	Status     uint8            `json:"status"`
	Reason     string           `gorm:"column:reason;type:varchar(255);comment:原因" json:"reason"`
	AdminId    int              `json:"admin_id"`
	AdminName  string           `gorm:"column:admin_name;type:varchar(100);comment:操作人姓名" json:"admin_name"`
	UpdateTime utils.CustomTime `gorm:"column:update_time;type:datetime" json:"update_time"`
}

type ProjectNameItems struct {
	Id   int    `gorm:"column:id;type:int(11);primary_key;AUTO_INCREMENT" json:"id"`
	Name string `gorm:"column:name;type:varchar(255);comment:项目描述;NOT NULL" json:"name"`
	Logo string `gorm:"column:logo;type:varchar(255);default:'';comment:logo;NOT NULL" json:"logo"`
}
