package dao

import (
	"base-service/utils"
)

const (
	TypeDir    = 1
	TypeWeb    = 2
	TypeButton = 3
	TypeUrl    = 4
	NoData     = "#"

	QuerySortLIstWeb       = 1
	QuerySortListAll       = 2
	QuerySortListNotButton = 3
)

// Authority 权限表
type Authority struct {
	Id             int              `gorm:"column:id;type:int(11);primary_key;AUTO_INCREMENT;comment:ID" json:"id"`
	Name           string           `gorm:"column:name;type:varchar(255);comment:菜单;NOT NULL" json:"name"`
	ParentId       int              `gorm:"column:parent_id;type:int(11);default:0;comment:父级菜单;NOT NULL" json:"parent_id"`
	ParentIds      string           `gorm:"column:parent_ids;type:varchar(255);comment:多级父级菜单;NOT NULL" json:"parent_ids"`
	Level          uint8            `gorm:"column:level;type:tinyint(2);default:0;comment:级别;NOT NULL" json:"level"`
	Type           uint8            `gorm:"column:type;type:tinyint(2);comment:类型，1目录，2页面，3按钮;NOT NULL" json:"type"`
	ProjectId      int              `gorm:"column:project_id;type:int(11);comment:项目ID;NOT NULL" json:"project_id"`
	Path           string           `gorm:"column:path;type:varchar(255);comment:页面请求地址;NOT NULL" json:"path"`
	Api            string           `gorm:"column:api;type:varchar(255);comment:接口地址" json:"api"`
	ViewPath       string           `gorm:"column:view_path;type:varchar(255);comment:页面文件路径" json:"view_path"`
	Identification string           `gorm:"column:identification;type:varchar(255);comment:权限标识" json:"identification"`
	Icon           string           `gorm:"column:icon;type:varchar(255);comment:图标" json:"icon"`
	Status         uint8            `gorm:"column:status;type:tinyint(1);default:1;comment:状态，1为正常，2为禁用;NOT NULL" json:"status"`
	Sort           uint8            `gorm:"column:sort;type:smallint(3);default:1;comment:排序" json:"sort"`
	IsShow         uint8            `gorm:"column:is_show;type:tinyint(1);default:1;comment:菜单中是否显示，1为显示，2为不显示;NOT NULL" json:"is_show"`
	Remarks        string           `gorm:"column:remarks;type:varchar(255);comment:备注;NOT NULL" json:"remarks"`
	Reason         string           `gorm:"column:reason;type:varchar(255);comment:原因" json:"reason"`
	AdminId        int              `gorm:"column:admin_id;type:int(11);comment:操作人ID" json:"admin_id"`
	AdminName      string           `gorm:"column:admin_name;type:varchar(255);comment:操作人员姓名" json:"admin_name"`
	CreateTime     utils.CustomTime `gorm:"column:create_time;default:NULL;type:datetime" json:"create_time"`
	UpdateTime     utils.CustomTime `gorm:"column:update_time;default:NULL;type:datetime" json:"update_time"`
}

// AuthorityStatus 修改状态
type AuthorityStatus struct {
	Id         int              `json:"id"`
	Status     uint8            `json:"status"`
	Reason     string           `gorm:"column:reason;type:varchar(255);comment:原因" json:"reason"`
	ProjectId  int              `gorm:"column:project_id;type:int(11);comment:项目ID;NOT NULL" json:"project_id"`
	AdminId    int              `json:"admin_id"`
	AdminName  string           `gorm:"column:admin_name;type:varchar(100);comment:操作人姓名" json:"admin_name"`
	UpdateTime utils.CustomTime `gorm:"column:update_time;type:datetime;default:NULL;comment:更新时间" json:"update_time"`
}

type AuthorityCheck struct {
	Id  int    `json:"id"`
	Api string `json:"api"`
}

type AuthorityItems struct {
	Id             int              `json:"id"`
	Name           string           `json:"name"`
	ParentId       int              `json:"parent_id"`
	ParentIds      string           `json:"parent_ids"`
	Level          uint8            `json:"level"`
	Type           uint8            `json:"type"`
	ProjectId      int              `json:"project_id"`
	Path           string           `json:"path"`
	Api            string           `json:"api"`
	ViewPath       string           `json:"view_path"`
	Identification string           `json:"identification"`
	Icon           string           `json:"icon"`
	Status         uint8            `json:"status"`
	Sort           uint8            `json:"sort"`
	IsShow         uint8            `json:"is_show"`
	Remarks        string           `json:"remarks"`
	Reason         string           `json:"reason"`
	AdminId        int              `json:"admin_id"`
	AdminName      string           `json:"admin_name"`
	CreateTime     utils.CustomTime `json:"create_time"`
	UpdateTime     utils.CustomTime `json:"update_time"`
	Children       []AuthorityItems `json:"children"`
}

type AuthorityNames struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Level    uint8  `json:"level"`
	Type     uint8  `json:"type"`
	ParentId int    `json:"parent_id"db:"parent_id"`
}

type AuthorityNamesItems struct {
	Id       int                   `json:"value"`
	Name     string                `json:"label"`
	Level    uint8                 `json:"level"`
	Type     uint8                 `json:"type"`
	ParentId int                   `json:"pid"`
	Children []AuthorityNamesItems `json:"children"`
}

type AuthorityInfo struct {
	Id             int    `gorm:"column:id;type:int(11);primary_key;AUTO_INCREMENT;comment:ID" json:"id"`
	Name           string `gorm:"column:name;type:varchar(255);comment:菜单;NOT NULL" json:"name"`
	ParentId       int    `gorm:"column:parent_id;type:int(11);default:0;comment:父级菜单;NOT NULL" json:"parent_id"`
	Type           uint8  `gorm:"column:type;type:tinyint(2);comment:类型，1目录，2页面，3按钮;NOT NULL" json:"type"`
	Path           string `gorm:"column:path;type:varchar(255);comment:页面请求地址;NOT NULL" json:"path"`
	Api            string `gorm:"column:api;type:varchar(255);comment:接口地址" json:"api"`
	ViewPath       string `gorm:"column:view_path;type:varchar(255);comment:页面文件路径" json:"view_path"`
	Identification string `gorm:"column:identification;type:varchar(255);comment:权限标识" json:"identification"`
	Icon           string `gorm:"column:icon;type:varchar(255);comment:图标" json:"icon"`
	IsShow         uint8  `gorm:"column:is_show;type:tinyint(1);default:1;comment:菜单中是否显示，1为显示，2为不显示;NOT NULL" json:"is_show"`
	Remarks        string `gorm:"column:remarks;type:varchar(255);comment:备注;NOT NULL" json:"remarks"`
	Reason         string `gorm:"column:reason;type:varchar(255);comment:原因" json:"reason"`
}
