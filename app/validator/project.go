package validator

// ProjectEditValidator 创建/编辑
type ProjectEditValidator struct {
	Id      int    `form:"id" json:"id"`
	Name    string `form:"name" json:"name" binding:"required"`
	Remarks string `form:"remarks" json:"remarks"`
	Logo    string `form:"logo" json:"logo"`
}

// GetMessage 创建 - 提示消息
func (p ProjectEditValidator) GetMessage() ValidatorMessages {
	return ValidatorMessages{
		"Name.required": "真实姓名不能为空",
	}
}

// ProjectInfoValidator 查看信息
type ProjectInfoValidator struct {
	Id int `form:"id" json:"id" binding:"required"`
}

// GetMessage 查看信息 - 提示消息
func (p ProjectInfoValidator) GetMessage() ValidatorMessages {
	return ValidatorMessages{
		"Id.required": "ID不能为空",
	}
}

// ProjectItemsValidator 查看列表
type ProjectItemsValidator struct {
	Page   int    `form:"page" json:"page"`
	Size   int    `form:"size" json:"size"`
	Name   string `form:"name" json:"name"`
	Status uint8  `form:"status" json:"status"`
}

// GetMessage 查看列表 - 提示消息
func (p ProjectItemsValidator) GetMessage() ValidatorMessages {
	return ValidatorMessages{}
}

// ProjectStatusValidator 修改状态
type ProjectStatusValidator struct {
	Id     int    `form:"id" json:"id" binding:"required"`
	Status uint8  `form:"status" json:"status" binding:"required,oneof=1 2"`
	Reason string `form:"reason" json:"reason"`
}

// GetMessage 修改状态 - 提示消息
func (p ProjectStatusValidator) GetMessage() ValidatorMessages {
	return ValidatorMessages{
		"Id.required":     "ID不能为空",
		"Status.required": "状态不能为空",
		"Status.oneof":    "状态传值不对",
	}
}

// ProjectNameItemsValidator 名称列表
type ProjectNameItemsValidator struct {
	Name string `form:"name" json:"name"`
}

// GetMessage 修改状态 - 提示消息
func (p ProjectNameItemsValidator) GetMessage() ValidatorMessages {
	return ValidatorMessages{}
}
