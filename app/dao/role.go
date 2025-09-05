package dao

import (
	"base-service/utils"
)

// Role 岗位表
type Role struct {
	Id         int              `gorm:"column:id;type:int(11);primary_key;AUTO_INCREMENT;comment:ID" json:"id"`
	Name       string           `gorm:"column:name;type:varchar(255);comment:名称;NOT NULL" json:"name"`
	ProjectId  int              `gorm:"column:project_id;type:int(11);comment:项目ID;NOT NULL" json:"project_id"`
	Status     uint8            `gorm:"column:status;type:tinyint(1);default:1;comment:状态，1为正常，2为禁用;NOT NULL" json:"status"`
	Rules      string           `gorm:"column:rules;type:text;comment:权限内容" json:"rules"`
	Reason     string           `gorm:"column:reason;type:varchar(255);comment:原因" json:"reason"`
	AdminId    int              `gorm:"column:admin_id;type:int(11)" json:"admin_id"`
	AdminName  string           `gorm:"column:admin_name;type:varchar(255);comment:操作人员姓名" json:"admin_name"`
	CreateTime utils.CustomTime `gorm:"column:create_time;default:NULL;type:datetime" json:"create_time"`
	UpdateTime utils.CustomTime `gorm:"column:update_time;default:NULL;type:datetime" json:"update_time"`
}

// RoleStatus 修改状态
type RoleStatus struct {
	Id         int              `json:"id"`
	Status     uint8            `json:"status"`
	Reason     string           `gorm:"column:reason;type:varchar(255);comment:原因" json:"reason"`
	AdminId    int              `json:"admin_id"`
	AdminName  string           `gorm:"column:admin_name;type:varchar(100);comment:操作人姓名" json:"admin_name"`
	UpdateTime utils.CustomTime `gorm:"column:update_time;type:datetime;default:NULL;comment:更新时间" json:"update_time"`
}

// RoleRule 修改状态
type RoleRule struct {
	Id         int              `json:"id"`
	Rules      string           `json:"rules"`
	AdminId    int              `json:"admin_id"`
	AdminName  string           `gorm:"column:admin_name;type:varchar(100);comment:操作人姓名" json:"admin_name"`
	UpdateTime utils.CustomTime `gorm:"column:update_time;type:datetime;default:NULL;comment:更新时间" json:"update_time"`
}

type RoleItems struct {
	Role
	RuleItems []string `json:"rule_items"`
}

type RoleNameItems struct {
	Id        int    `gorm:"column:id;type:int(11);primary_key;AUTO_INCREMENT" json:"id"`
	Name      string `gorm:"column:name;type:varchar(255);comment:项目描述;NOT NULL" json:"name"`
	ProjectId int    `json:"project_id"`
}

type RoleProjectNameItems struct {
	Id    int             `gorm:"column:id;type:int(11);primary_key;AUTO_INCREMENT" json:"id"`
	Name  string          `gorm:"column:name;type:varchar(255);comment:项目描述;NOT NULL" json:"name"`
	Items []RoleNameItems `json:"items"`
}

type RoleRulesItems struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Rules string `json:"rules"`
}

type RoleInfo struct {
	Id        int    `gorm:"column:id;type:int(11);primary_key;AUTO_INCREMENT;comment:ID" json:"id"`
	Name      string `gorm:"column:name;type:varchar(255);comment:名称;NOT NULL" json:"name"`
	ProjectId int    `gorm:"column:project_id;type:int(11);comment:项目ID;NOT NULL" json:"project_id"`
	Status    uint8  `gorm:"column:status;type:tinyint(1);default:1;comment:状态，1为正常，2为禁用;NOT NULL" json:"status"`
	Rules     []int  `gorm:"column:rules;type:text;comment:权限内容" json:"rules"`
	Reason    string `gorm:"column:reason;type:varchar(255);comment:原因" json:"reason"`
}
