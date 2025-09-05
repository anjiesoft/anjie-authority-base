package dao

const (
	StatusOk   uint8  = 1   //全局基础正常状态
	StatusFail uint8  = 2   //全局基础禁用状态
	MaxSort    uint16 = 999 //全局基础最大排序
	MinSort    uint8  = 1   //全局基础最小排序
	SIZE              = 20  //全局基础每页数

	LoginLogStatusOk   uint8 = 1
	LoginLogStatusFail uint8 = 2
)
