package validator

type LoginValidator struct {
	Username string `form:"username" json:"username" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

// GetMessage 提示消息
func (p LoginValidator) GetMessage() ValidatorMessages {
	return ValidatorMessages{
		"Password.required": "密码不能为空",
		"Username.required": "账号不能为空",
	}
}

type LoginRoutesValidator struct {
	ProjectId int `form:"project_id" json:"project_id" binding:"required"`
}

// GetMessage 提示消息
func (p LoginRoutesValidator) GetMessage() ValidatorMessages {
	return ValidatorMessages{
		"ProjectId.required": "项目ID不能为空",
	}
}
