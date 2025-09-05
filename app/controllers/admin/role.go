package admin

import (
	controller "base-service/app/controllers"
	"base-service/app/dao"
	"base-service/app/services"
	"base-service/app/validator"
	"base-service/utils"
	"github.com/gin-gonic/gin"
)

type RoleController struct {
	controller.BaseController
}

func (cc RoleController) Edit(c *gin.Context) {
	var params validator.RoleEditValidator
	if err := c.ShouldBind(&params); err != nil {
		cc.Out(c, nil, validator.Out(params, err))
		return
	}
	id, editErr := services.RoleService{}.Edit(c, params)
	if editErr != nil {
		cc.Error(c, editErr)
		return
	}
	cc.Ok(c, id)
}

func (cc RoleController) Items(c *gin.Context) {
	var params validator.RoleItemsValidator
	if err := c.ShouldBind(&params); err != nil {
		cc.Out(c, nil, validator.Out(params, err))
		return
	}
	items, err := services.RoleService{}.Items(params)
	if err != nil {
		cc.Fail(c)
		return
	}
	cc.Ok(c, items)
}

func (cc RoleController) Status(c *gin.Context) {
	var params validator.RoleStatusValidator
	if err := c.ShouldBind(&params); err != nil {
		cc.Out(c, nil, validator.Out(params, err))
		return
	}

	if params.Status == dao.StatusFail && len(params.Reason) < 1 {
		cc.Error(c, utils.ErrorReasonParams)
		return
	}

	is, pErr := services.RoleService{}.Status(c, params)

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

func (cc RoleController) Rule(c *gin.Context) {
	var params validator.RoleRuleValidator
	if err := c.ShouldBind(&params); err != nil {
		cc.Out(c, nil, validator.Out(params, err))
		return
	}

	is, pErr := services.RoleService{}.Rule(c, params.Id, params.Rule)

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

func (cc RoleController) Info(c *gin.Context) {

	var params validator.RoleInfoValidator

	if err := c.ShouldBind(&params); err != nil {
		cc.Out(c, nil, validator.Out(params, err))
		return
	}
	info, err := services.RoleService{}.Info(params.Id)
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

func (cc RoleController) NameItems(c *gin.Context) {
	var params validator.RoleNameItemsValidator
	if err := c.ShouldBind(&params); err != nil {
		cc.Out(c, nil, validator.Out(params, err))
		return
	}
	items, err := services.RoleService{}.NameItems(params.Name)
	if err != nil {
		cc.Fail(c)
		return
	}
	cc.Ok(c, items)
}
