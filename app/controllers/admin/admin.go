package admin

import (
	controller "base-service/app/controllers"
	"base-service/app/dao"
	"base-service/app/services"
	"base-service/app/validator"
	"base-service/utils"
	"github.com/gin-gonic/gin"
)

type AdminController struct {
	controller.BaseController
}

func (cc AdminController) Edit(c *gin.Context) {
	var params validator.AdminEditValidator
	if err := c.ShouldBind(&params); err != nil {
		cc.Out(c, nil, validator.Out(params, err))
		return
	}
	id, editErr := services.AdminService{}.Edit(c, params)
	if editErr != nil {
		cc.Error(c, editErr)
		return
	}
	cc.Ok(c, id)
}

func (cc AdminController) Items(c *gin.Context) {
	var params validator.AdminItemsValidator
	if err := c.ShouldBind(&params); err != nil {
		cc.Out(c, nil, validator.Out(params, err))
		return
	}
	items, err := services.AdminService{}.Items(params)
	if err != nil {
		cc.Fail(c)
		return
	}
	cc.Ok(c, items)
}

func (cc AdminController) Status(c *gin.Context) {
	var params validator.AdminStatusValidator
	if err := c.ShouldBind(&params); err != nil {
		cc.Out(c, nil, validator.Out(params, err))
		return
	}

	if params.Status == dao.StatusFail && len(params.Reason) < 1 {
		cc.Error(c, utils.ErrorReasonParams)
		return
	}

	is, pErr := services.AdminService{}.Status(c, params)

	if pErr != nil {
		cc.Error(c, pErr)
		return
	}
	if is {
		cc.Ok(c, nil)
		return
	}

	cc.Error(c, utils.ErrorNoChange)
	return
}

func (cc AdminController) Password(c *gin.Context) {
	var params validator.AdminPasswordValidator
	if err := c.ShouldBind(&params); err != nil {
		cc.Out(c, nil, validator.Out(params, err))
		return
	}

	is, pErr := services.AdminService{}.Password(c, params.Id, params.Password)

	if pErr != nil {
		cc.Error(c, pErr)
		return
	}
	if is {
		cc.Ok(c, nil)
		return
	}

	cc.Error(c, utils.ErrorNoChange)
	return
}

func (cc AdminController) OwnPassword(c *gin.Context) {
	var params validator.AdminOwnPasswordValidator
	if err := c.ShouldBind(&params); err != nil {
		cc.Out(c, nil, validator.Out(params, err))
		return
	}

	is, pErr := services.AdminService{}.OwnPassword(c, params.Password)

	if pErr != nil {
		cc.Error(c, pErr)
		return
	}
	if is {
		cc.Ok(c, nil)
		return
	}

	cc.Error(c, utils.ErrorNoChange)
	return
}

func (cc AdminController) Info(c *gin.Context) {

	var params validator.AdminInfoValidator

	if err := c.ShouldBind(&params); err != nil {
		cc.Out(c, nil, validator.Out(params, err))
		return
	}
	info, err := services.AdminService{}.Info(params.Id)
	if err != nil {
		cc.Fail(c)
		return
	}
	if info.Id < 1 {
		cc.Ok(c, nil)
		return
	}

	cc.Ok(c, info)
}
