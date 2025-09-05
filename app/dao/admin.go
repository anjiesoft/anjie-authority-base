package dao

import (
	"base-service/utils"
)

// Admin 公司内部人员表
type Admin struct {
	Id         int              `gorm:"column:id;type:int(11);primary_key;AUTO_INCREMENT;comment:ID" json:"id"`
	Username   string           `gorm:"column:username;type:varchar(30);comment:账号;NOT NULL" json:"username"`
	Avatar     string           `gorm:"column:avatar;type:varchar(255);comment:头像" json:"avatar"`
	Name       string           `gorm:"column:name;type:varchar(30);comment:真实姓名;NOT NULL" json:"name"`
	Phone      string           `gorm:"column:phone;type:char(12);comment:电话" json:"phone"`
	Email      string           `gorm:"column:email;type:varchar(255);comment:邮箱" json:"email"`
	Password   string           `gorm:"column:password;type:char(61);comment:密码;NOT NULL" json:"password"`
	AdminName  string           `gorm:"column:admin_name;type:varchar(100);comment:操作人姓名" json:"admin_name"`
	Salt       string           `gorm:"column:salt;type:char(5);comment:与密码混合搭配加密;NOT NULL" json:"salt"`
	Roles      string           `gorm:"column:roles;type:varchar(255);comment:角色" json:"roles"`
	Status     uint8            `gorm:"column:status;type:tinyint(1);default:1;comment:状态，1为正常，2为禁用;NOT NULL" json:"status"`
	LastTime   utils.CustomTime `gorm:"column:last_time;type:datetime;default:NULL;comment:最后一次登录时间" json:"last_time"`
	FailNumber int              `gorm:"column:fail_number;type:tinyint(2);default:0;comment:连续登录失败次数" json:"fail_number"`
	Reason     string           `gorm:"column:reason;type:varchar(255);comment:原因" json:"reason"`
	AdminId    int              `gorm:"column:admin_id;type:int(11);default:0;comment:操作人;NOT NULL" json:"admin_id"`
	CreateTime utils.CustomTime `gorm:"column:create_time;type:datetime;default:NULL;comment:添加时间" json:"create_time"`
	UpdateTime utils.CustomTime `gorm:"column:update_time;type:datetime;default:NULL;comment:更新时间" json:"update_time"`
}

// AdminPassword 修改管理员密码
type AdminPassword struct {
	Id         int              `json:"id"`
	Password   string           `json:"password"`
	Salt       string           `json:"salt"`
	AdminId    int              `json:"admin_id"`
	AdminName  string           `gorm:"column:admin_name;type:varchar(100);comment:操作人姓名" json:"admin_name"`
	UpdateTime utils.CustomTime `gorm:"column:update_time;type:datetime;default:NULL;comment:更新时间" json:"update_time"`
}

// AdminFail 登录失败
type AdminFail struct {
	Id         int `json:"id"`
	FailNumber int `gorm:"column:fail_number;type:tinyint(2);default:0;comment:连续登录失败次数" json:"fail_number"`
}

// AdminStatus 修改管理员状态
type AdminStatus struct {
	Id         int              `json:"id"`
	Status     uint8            `json:"status"`
	Reason     string           `json:"reason"`
	AdminId    int              `json:"admin_id"`
	AdminName  string           `gorm:"column:admin_name;type:varchar(100);comment:操作人姓名" json:"admin_name"`
	UpdateTime utils.CustomTime `gorm:"column:update_time;type:datetime;default:NULL;comment:更新时间" json:"update_time"`
}

// AdminInfo 管理员详情
type AdminInfo struct {
	Id         int              `gorm:"column:id;type:int(11);primary_key;AUTO_INCREMENT;comment:ID" json:"id"`
	Username   string           `gorm:"column:username;type:varchar(30);comment:账号;NOT NULL" json:"username"`
	Avatar     string           `gorm:"column:avatar;type:varchar(255);comment:头像" json:"avatar"`
	Name       string           `gorm:"column:name;type:varchar(30);comment:真实姓名;NOT NULL" json:"name"`
	Phone      string           `gorm:"column:phone;type:char(12);comment:电话" json:"phone"`
	Status     uint8            `gorm:"column:status;type:tinyint(1);default:1;comment:状态，1为正常，2为禁用;NOT NULL" json:"status"`
	Email      string           `gorm:"column:email;type:varchar(255);comment:邮箱" json:"email"`
	Roles      string           `gorm:"column:roles;type:varchar(255);comment:角色" json:"roles"`
	LastTime   utils.CustomTime `gorm:"column:last_time;type:datetime;comment:最后一次登录时间" json:"last_time"`
	FailNumber int              `gorm:"column:fail_number;type:tinyint(2);default:0;comment:连续登录失败次数" json:"fail_number"`
}

type AdminInfoRet struct {
	AdminInfo
	Roles []int `json:"roles"`
}

type AdminItemsRet struct {
	AdminInfo
	RoleInfo []RoleNameItems `json:"role_info"`
}

type AdminItems struct {
	Items []AdminItemsRet `json:"items"`
	Count int64           `json:"count"`
}

type AdminLastTime struct {
	Id       int              `gorm:"column:id;type:int(11);primary_key;AUTO_INCREMENT;comment:ID" json:"id"`
	LastTime utils.CustomTime `gorm:"column:last_time;type:datetime;default:NULL;comment:最后一次登录时间" json:"last_time"`
}
