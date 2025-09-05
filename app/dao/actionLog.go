package dao

import (
	"base-service/utils"
)

const (
	//增加、删除、修改时需要改 services.log中的返回值

	ActionTypeAdd      = 1 //增加
	ActionTypeEdit     = 2 //编辑
	ActionTypeDel      = 3 //删除
	ActionTypeOut      = 4 //强退
	ActionTypeDown     = 5 //下载
	ActionTypeStatus   = 6 //状态
	ActionTypePassword = 7 //修改密码

	//增加、删除、修改时需要改 services.log中的返回值

	ModuleNameAdmin      = "管理员管理"
	ModuleNamePost       = "岗位管理"
	ModuleNameRole       = "角色管理"
	ModuleNameProject    = "项目管理"
	ModuleNameDepartment = "部门管理"
	ModuleNameAuthority  = "权限管理"
)

// ActionLog 操作日志表
type ActionLog struct {
	Id             int              `gorm:"column:id;type:int(11);primary_key;AUTO_INCREMENT;comment:ID" json:"id"`
	Content        string           `gorm:"column:content;type:text;comment:内容;NOT NULL" json:"content"`
	AdminId        int              `gorm:"column:admin_id;type:int(11);comment:操作人;NOT NULL" json:"admin_id"`
	AdminName      string           `gorm:"column:admin_name;type:varchar(255);comment:真实姓名" json:"admin_name"`
	Type           uint8            `gorm:"column:type;type:smallint(3);comment:类型，以常量配置文件为准，目前，1为添加，2为编辑，3为删除，4为强退，5为下载;NOT NULL" json:"type"`
	ModuleName     string           `gorm:"column:module_name;type:varchar(255);comment:操作模块名，如管理员;NOT NULL" json:"module_name"`
	Ip             string           `gorm:"column:ip;type:varchar(20);comment:IP;NOT NULL" json:"ip"`
	BrowserInfo    string           `gorm:"column:browser_info;type:varchar(255);comment:浏览器详细信息" json:"browser_info"`
	BrowserName    string           `gorm:"column:browser_name;type:varchar(50);comment:浏览器名称" json:"browser_name"`
	BrowserVersion string           `gorm:"column:browser_version;type:varchar(50);comment:版本" json:"browser_version"`
	ProjectId      int              `gorm:"column:project_id;type:int(11);default:0;comment:项目ID" json:"project_id"`
	Reason         string           `gorm:"column:reason;type:varchar(255);comment:原因" json:"reason"`
	CreateTime     utils.CustomTime `gorm:"column:create_time;type:datetime;default:NULL;comment:添加时间;NOT NULL" json:"create_time"`
}

type ActionItems struct {
	Items []ActionLog  `json:"items"`
	Count int64        `json:"count"`
	Type  []ActionType `json:"type"`
}

type ActionAdd struct {
	NewContent interface{} `json:"new_content"`
	OldContent interface{} `json:"old_content"`
	ProjectId  int         `json:"project_id"`
	Reason     string      `json:"reason"`
	Type       uint8       `json:"type"`
	ModuleName string      `json:"module_name"`
}

type ActionType struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}
