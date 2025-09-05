package utils

import "encoding/json"

var _ Error = (*err)(nil)

type Error interface {
	// i 为了避免被其他包实现
	i()
	// WithData 设置成功时返回的数据
	WithData(data interface{}) Error

	GetCode() int
	GetMsg() string
	GetData() interface{}
	// ToString 返回 JSON 格式的错误详情
	ToString() string
}

type err struct {
	Code int         `json:"code"`    // 业务编码
	Msg  string      `json:"message"` // 错误描述
	Data interface{} `json:"data"`    // 成功时返回的数据
}

func NewError(code int, msg string) Error {
	return &err{
		Code: code,
		Msg:  msg,
		Data: nil,
	}
}
func (e *err) i() {}

func (e *err) WithData(data interface{}) Error {
	e.Data = data
	return e
}

func (e *err) WithID(id string) Error {
	return e
}

// ToString 返回 JSON 格式的错误详情
func (e *err) ToString() string {
	err := &struct {
		Code int         `json:"code"`
		Msg  string      `json:"msg"`
		Data interface{} `json:"data"`
	}{
		Code: e.Code,
		Msg:  e.Msg,
		Data: e.Data,
	}

	raw, _ := json.Marshal(err)
	return string(raw)
}

func (e *err) GetCode() int {
	return e.Code
}

func (e *err) GetData() interface{} {
	return e.Data
}

func (e *err) GetMsg() string {
	return e.Msg
}
