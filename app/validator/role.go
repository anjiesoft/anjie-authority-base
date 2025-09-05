package validator

// RoleEditValidator 创建/编辑
type RoleEditValidator struct {
	Id        int    `form:"id" json:"id"`
	Name      string `form:"name" json:"name" binding:"required"`
	ProjectId int    `form:"project_id" json:"project_id"`
	Rules     []int  `form:"rules" json:"rules"`
}

// GetMessage 创建 - 提示消息
func (p RoleEditValidator) GetMessage() ValidatorMessages {
	return ValidatorMessages{
		"Name.required": "名称不能为空",
	}
}

// RoleInfoValidator 查看信息
type RoleInfoValidator struct {
	Id int `form:"id" json:"id" binding:"required"`
}

// GetMessage 查看信息 - 提示消息
func (p RoleInfoValidator) GetMessage() ValidatorMessages {
	return ValidatorMessages{
		"Id.required": "ID不能为空",
	}
}

// RoleItemsValidator 查看列表
type RoleItemsValidator struct {
	Page      int    `form:"page" json:"page"`
	Size      int    `form:"size" json:"size"`
	Name      string `form:"name" json:"name"`
	Status    uint8  `form:"status" json:"status"`
	ProjectId int    `form:"project_id" json:"project_id" binding:"required"`
}

// GetMessage 查看列表 - 提示消息
func (p RoleItemsValidator) GetMessage() ValidatorMessages {
	return ValidatorMessages{
		"ProjectId.required": "项目不能为空",
	}
}

// RoleStatusValidator 修改状态
type RoleStatusValidator struct {
	Id     int    `form:"id" json:"id" binding:"required"`
	Status uint8  `form:"status" json:"status" binding:"required,oneof=1 2"`
	Reason string `form:"reason" json:"reason"`
}

// GetMessage 修改状态 - 提示消息
func (p RoleStatusValidator) GetMessage() ValidatorMessages {
	return ValidatorMessages{
		"Id.required":     "ID不能为空",
		"Status.required": "状态不能为空",
		"Status.oneof":    "状态传值不对",
	}
}

// RoleRuleValidator 编辑权限
type RoleRuleValidator struct {
	Id   int    `form:"id" json:"id" binding:"required"`
	Rule string `form:"rule" json:"rule" binding:"required"`
}

// GetMessage 修改状态 - 提示消息
func (p RoleRuleValidator) GetMessage() ValidatorMessages {
	return ValidatorMessages{
		"Id.required":   "ID不能为空",
		"Rule.required": "权限不能为空",
	}
}

// RoleNameItemsValidator 名称列表
type RoleNameItemsValidator struct {
	Name      string `form:"name" json:"name"`
	ProjectId int    `form:"project_id" json:"project_id"`
}

// GetMessage 修改状态 - 提示消息
func (p RoleNameItemsValidator) GetMessage() ValidatorMessages {
	return ValidatorMessages{}
}
