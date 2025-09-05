package validator

// AuthorityEditValidator 创建/编辑
type AuthorityEditValidator struct {
	Id             int    `form:"id" json:"id"`
	Name           string `form:"name" json:"name" binding:"required"`
	ParentId       int    `form:"parent_id" json:"parent_id"`
	Type           uint8  `form:"type" json:"type" binding:"required,oneof=1 2 3 4"`
	ProjectId      int    `form:"project_id" json:"project_id" binding:"required"`
	Path           string `form:"path" json:"path"`
	Api            string `form:"api" json:"api"`
	ViewPath       string `form:"view_path" json:"view_path"`
	Identification string `form:"identification" json:"identification" binding:"required"`
	Icon           string `form:"icon" json:"icon"`
	IsShow         uint8  `form:"is_show" json:"is_show" binding:"required,oneof=1 2"`
	Remarks        string `form:"remarks" json:"remarks"`
}

// GetMessage 创建/编辑 - 提示消息
func (p AuthorityEditValidator) GetMessage() ValidatorMessages {
	return ValidatorMessages{
		"Name.required":           "名称不能为空",
		"Type.required":           "类型不能为空",
		"ProjectId.required":      "项目不能为空",
		"Identification.required": "权限标识不能为空",
		"IsShow.required":         "是否显示不能为空",
		"Type.oneof":              "类型传值不对",
		"IsShow.oneof":            "显示传值不对",
	}
}

// AuthorityInfoValidator 查看信息
type AuthorityInfoValidator struct {
	Id int `form:"id" json:"id" binding:"required"`
}

// GetMessage 查看信息 - 提示消息
func (p AuthorityInfoValidator) GetMessage() ValidatorMessages {
	return ValidatorMessages{
		"Id.required": "ID不能为空",
	}
}

// AuthorityItemsValidator 查看列表
type AuthorityItemsValidator struct {
	ProjectId int `form:"project_id" json:"project_id" binding:"required"`
}

// GetMessage 查看列表 - 提示消息
func (p AuthorityItemsValidator) GetMessage() ValidatorMessages {
	return ValidatorMessages{
		"ProjectId.required": "项目不能为空",
	}
}

// AuthorityStatusValidator 修改状态
type AuthorityStatusValidator struct {
	Id     int    `form:"id" json:"id" binding:"required"`
	Status uint8  `form:"status" json:"status" binding:"required,oneof=1 2"`
	Reason string `form:"reason" json:"reason"`
}

// GetMessage 修改状态 - 提示消息
func (p AuthorityStatusValidator) GetMessage() ValidatorMessages {
	return ValidatorMessages{
		"Id.required":     "ID不能为空",
		"Status.required": "状态不能为空",
		"Status.oneof":    "状态传值不对",
	}
}

// AuthorityNameItemsValidator 名称列表
type AuthorityNameItemsValidator struct {
	ProjectId int   `form:"project_id" json:"project_id" binding:"required"`
	Type      uint8 `form:"type" json:"type"`
}

// GetMessage 修改状态 - 提示消息
func (p AuthorityNameItemsValidator) GetMessage() ValidatorMessages {
	return ValidatorMessages{
		"ProjectId.required": "项目不能为空",
	}
}

// AuthorityValidator 查看信息
type AuthorityValidator struct {
	Id int `form:"id" json:"id" binding:"required"`
}

// GetMessage 查看信息 - 提示消息
func (p AuthorityValidator) GetMessage() ValidatorMessages {
	return ValidatorMessages{
		"Id.required": "ID不能为空",
	}
}
