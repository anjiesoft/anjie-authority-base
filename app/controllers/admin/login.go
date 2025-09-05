package admin

import (
	controller "base-service/app/controllers"
	"base-service/app/services"
	"base-service/app/validator"
	"base-service/utils"
	"github.com/gin-gonic/gin"
)

type LoginController struct {
	controller.BaseController
}

func (cc LoginController) Login(c *gin.Context) {
	var params validator.LoginValidator
	loginLog := services.LogService{}
	if err := c.ShouldBind(&params); err != nil {
		out := validator.Out(params, err)
		cc.Out(c, nil, out)
		loginLog.Fail(c, out.ToString(), params.Username)
		return
	}
	admin, err := services.LoginService{}.Login(params.Username)
	if err != nil {
		loginLog.Fail(c, err.Error(), params.Username)
		cc.Fail(c)
		return
	}
	pwd := utils.GetSaltPassword(admin.Salt, params.Password)
	eq := utils.EqualsPassword(pwd, admin.Password)
	if !eq {
		cc.Error(c, utils.ErrorPwdError)
		services.LoginService{}.Fail(admin.Id, admin.FailNumber)
		loginLog.Fail(c, utils.ErrorPwdError.GetMsg(), params.Username)
		return
	}

	data := utils.JwtInfo{
		Id:       admin.Id,
		Name:     admin.Name,
		Email:    admin.Email,
		Phone:    admin.Phone,
		Avatar:   admin.Avatar,
		Username: admin.Username,
		Roles:    admin.Roles,
		LastTime: utils.Now(),
	}
	jwt := utils.Jwt{}
	accessToken, accessErr := jwt.GetToken(data, true)
	if accessErr != nil {
		cc.Fail(c)
		loginLog.Fail(c, accessErr.Error(), params.Username)
		return
	}
	refreshToken, refreshErr := jwt.GetToken(data, false)
	if refreshErr != nil {
		cc.Fail(c)
		loginLog.Fail(c, refreshErr.Error(), params.Username)
		return
	}
	loginLog.Ok(c, data)
	services.LoginService{}.Ok(admin.Id)
	cc.Ok(c, utils.Token{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	})
}

func (cc LoginController) Refresh(c *gin.Context) {
	token := c.GetHeader("RefreshAuthorization")
	if len(token) < 10 {
		cc.Error(c, utils.ErrorMissingParams)
		return
	}
	// 去掉Bearer
	token = token[7:]
	// 使用之前定义好的解析JWT的函数 ParseToken 来解析它
	jwt := utils.Jwt{}
	userInfo, err := jwt.ParseToken(token, false)
	if err != nil {
		cc.Error(c, utils.ErrorPwdError)
		return
	}

	accessToken, accessErr := jwt.GetToken(*userInfo, true)
	if accessErr != nil {
		cc.Fail(c)
		return
	}
	refreshToken, refreshErr := jwt.GetToken(*userInfo, false)
	if refreshErr != nil {
		cc.Fail(c)
		return
	}

	services.LoginService{}.Ok(userInfo.Id)
	cc.Ok(c, utils.Token{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	})
}

func (cc LoginController) Info(c *gin.Context) {
	id, isSet := c.Get(utils.USERID)
	if !isSet {
		cc.Error(c, utils.ErrorNoLogin)
		return
	}

	info, err := services.LoginService{}.Info(id.(int))
	if err != nil {
		cc.Fail(c)
		return
	}
	cc.Ok(c, info)
}

func (cc LoginController) Routes(c *gin.Context) {
	var params validator.LoginRoutesValidator
	if err := c.ShouldBind(&params); err != nil {
		cc.Out(c, nil, validator.Out(params, err))
		return
	}

	items, err := services.RoutesService{}.Items(c, params)
	if err != nil {
		cc.Error(c, err)
		return
	}

	cc.Ok(c, items)
}
