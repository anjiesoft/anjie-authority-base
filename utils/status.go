package utils

type Status struct {
	Code int    `json:"code"`
	Msg  string `json:"message"`
}

const (
	SuperCompanyId = 0
	USERINFO       = "user_info"
	USERID         = "user_id"
	USERNAME       = "user_name"
	SUPERID        = "supper_id"
)

var (
	// 基础类

	OK             = NewError(10000, "OK")
	Fail           = NewError(999, "fail")
	ErrorTypeError = NewError(998, "类型不对")
	ErrorExist     = NewError(997, "exist")

	// 参数类

	ErrorMissingParams         = NewError(10001, "缺少参数")
	ErrorReasonParams          = NewError(10001, "原因不能为空")
	ErrorPwdError              = NewError(10002, "密码错误")
	ErrorNoLogin               = NewError(10003, "没有登录")
	ErrorJwtError              = NewError(10004, "jwt错误")
	ErrorJwtExpired            = NewError(10005, "jwt过期")
	ErrorMissingProjectId      = NewError(10006, "缺少权限项目ID")
	ErrorMissingIdentification = NewError(10007, "缺少权限标识")
	ErrorAuthority             = NewError(10008, "没有权限")
	ErrorExistUser             = NewError(10009, "账号或手机号已存在")

	// 数据类

	ErrorNotFund            = NewError(11001, "not fund")
	ErrorNoChange           = NewError(11002, "没有改变")
	ErrorAuthorityChange    = NewError(11003, "该权限下有其它权限")
	ErrorMenuMissingShow    = NewError(11005, "是否显示")
	ErrorMenuMissingType    = NewError(11006, "是否显示")
	ErrorMenuMissingProject = NewError(11007, "项目错误")
)
