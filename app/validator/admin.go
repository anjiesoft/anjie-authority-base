package validator

// AdminEditValidator 创建/编辑
type AdminEditValidator struct {
	Id       int    `form:"id" json:"id"`
	Username string `form:"username" json:"username" binding:"required"`
	Name     string `form:"name" json:"name" binding:"required"`
	Password string `form:"password" json:"password"`
	Avatar   string `form:"avatar" json:"avatar"`
	Phone    string `form:"phone" json:"phone" binding:"required"`
	Email    string `form:"email" json:"email" binding:"required,email"`
	Roles    []int  `form:"roles" json:"roles"`
}

// GetMessage 创建/编辑 - 提示消息
func (p AdminEditValidator) GetMessage() ValidatorMessages {
	return ValidatorMessages{
		"Name.required":     "真实姓名不能为空",
		"Password.required": "密码不能为空",
		"Phone.required":    "手机号不能为空",
		"Email.email":       "邮箱格式不对",
		"Email.required":    "邮箱不能为空",
		"Username.required": "账号不能为空",
	}
}

// AdminPasswordValidator 修改所有管理员密码
type AdminPasswordValidator struct {
	Id       int    `form:"id" json:"id" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

// GetMessage 修改所有管理员密码 - 提示消息
func (p AdminPasswordValidator) GetMessage() ValidatorMessages {
	return ValidatorMessages{
		"Id.required":       "ID不能为空",
		"Password.required": "密码不能为空",
	}
}

// AdminOwnPasswordValidator 修改所有管理员密码
type AdminOwnPasswordValidator struct {
	Password string `form:"password" json:"password" binding:"required"`
}

// GetMessage 修改所有管理员密码 - 提示消息
func (p AdminOwnPasswordValidator) GetMessage() ValidatorMessages {
	return ValidatorMessages{
		"Password.required": "密码不能为空",
	}
}

// AdminInfoValidator 查看信息
type AdminInfoValidator struct {
	Id int `form:"id" json:"id" binding:"required"`
}

// GetMessage 查看信息 - 提示消息
func (p AdminInfoValidator) GetMessage() ValidatorMessages {
	return ValidatorMessages{
		"Id.required": "ID不能为空",
	}
}

// AdminItemsValidator 查看列表
type AdminItemsValidator struct {
	Page     int    `form:"page" json:"page"`
	Size     int    `form:"size" json:"size"`
	Phone    string `form:"phone" json:"phone"`
	Username string `form:"username" json:"username"`
	Name     string `form:"name" json:"name"`
	Email    string `form:"email" json:"email"`
	Status   uint8  `form:"status" json:"status"`
}

// GetMessage 查看列表 - 提示消息
func (p AdminItemsValidator) GetMessage() ValidatorMessages {
	return ValidatorMessages{}
}

// AdminStatusValidator 修改状态
type AdminStatusValidator struct {
	Id     int    `form:"id" json:"id" binding:"required"`
	Status uint8  `form:"status" json:"status" binding:"required,oneof=1 2"`
	Reason string `form:"reason" json:"reason"`
}

// GetMessage 修改状态 - 提示消息
func (p AdminStatusValidator) GetMessage() ValidatorMessages {
	return ValidatorMessages{
		"Id.required":     "ID不能为空",
		"Status.required": "状态不能为空",
		"Status.oneof":    "状态传值不对",
	}
}
