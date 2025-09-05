package validator

import (
	"base-service/utils"
	"github.com/go-playground/validator/v10"
)

// ValidatorMessages 校验信息
type ValidatorMessages map[string]string

// Validator 存放GetMessages()方法
type Validator interface {
	GetMessage() ValidatorMessages
}

// GetErrorMsg GetErrorMsg方法， 获取错误信息
func GetErrorMsg(request interface{}, err error) string {
	if _, isValidatorErrors := err.(validator.ValidationErrors); isValidatorErrors {
		_, isValidator := request.(Validator)

		for _, v := range err.(validator.ValidationErrors) {
			// 若 request 结构体实现 Validator 接口即可实现自定义错误信息
			if isValidator {
				if message, exist := request.(Validator).GetMessage()[v.Field()+"."+v.Tag()]; exist {
					return message
				}
			}
			return v.Error()
		}
	}

	return "Parameter error"
}

func Out(request interface{}, err error) utils.Error {
	return utils.NewError(utils.ErrorMissingParams.GetCode(), GetErrorMsg(request, err))
}
