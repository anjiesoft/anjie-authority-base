package dao

import (
	"base-service/utils"
)

// LoginLog 登录日志表
type LoginLog struct {
	Id             int              `gorm:"column:id;type:int(11);primary_key;AUTO_INCREMENT;comment:ID" json:"id"`
	AdminId        int              `gorm:"column:admin_id;type:int(11);default:0;comment:登录人ID;NOT NULL" json:"admin_id"`
	Ip             string           `gorm:"column:ip;type:varchar(50);comment:IP;NOT NULL" json:"ip"`
	Username       string           `gorm:"column:username;type:varchar(255);comment:用户名" json:"username"`
	Name           string           `gorm:"column:name;type:varchar(255);comment:真实姓名" json:"name"`
	BrowserInfo    string           `gorm:"column:browser_info;type:varchar(255);comment:浏览器详细信息" json:"browser_info"`
	BrowserName    string           `gorm:"column:browser_name;type:varchar(50);comment:浏览器名称" json:"browser_name"`
	BrowserVersion string           `gorm:"column:browser_version;type:varchar(50);comment:浏览器版本" json:"browser_version"`
	Status         uint8            `gorm:"column:status;type:tinyint(1);default:1;comment:状态，1为成功，2为失败" json:"status"`
	CreateTime     utils.CustomTime `gorm:"column:create_time;type:datetime;comment:添加时间;NOT NULL" json:"create_time"`
	Reason         string           `gorm:"column:reason;type:varchar(50);comment:失败原因" json:"reason"`
}

type LoginLogItems struct {
	Items []LoginLog `json:"items"`
	Count int64      `json:"count"`
}

type Count struct {
	Id int `gorm:"column:id;type:int(11);primary_key;AUTO_INCREMENT;comment:ID" json:"id"`
}
