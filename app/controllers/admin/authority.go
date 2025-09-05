package admin

import (
	controller "base-service/app/controllers"
	"base-service/app/dao"
	"base-service/app/services"
	"base-service/app/validator"
	"base-service/utils"
	"github.com/gin-gonic/gin"
)

type AuthorityController struct {
	controller.BaseController
}

func (cc AuthorityController) Edit(c *gin.Context) {
	var params validator.AuthorityEditValidator
	if err := c.ShouldBind(&params); err != nil {
		cc.Out(c, nil, validator.Out(params, err))
		return
	}
	id, editErr := services.AuthorityService{}.Edit(c, params)
	if editErr != nil {
		cc.Error(c, editErr)
		return
	}
	cc.Ok(c, id)
}

func (cc AuthorityController) Items(c *gin.Context) {
	var params validator.AuthorityItemsValidator
	if err := c.ShouldBind(&params); err != nil {
		cc.Out(c, nil, validator.Out(params, err))
		return
	}
	items, err := services.AuthorityService{}.Items(params)
	if err != nil {
		cc.Fail(c)
		return
	}
	cc.Ok(c, items)
}

func (cc AuthorityController) Status(c *gin.Context) {
	var params validator.AuthorityStatusValidator
	if err := c.ShouldBind(&params); err != nil {
		cc.Out(c, nil, validator.Out(params, err))
		return
	}

	if params.Status == dao.StatusFail && len(params.Reason) < 1 {
		cc.Error(c, utils.ErrorReasonParams)
		return
	}

	is, pErr := services.AuthorityService{}.Status(c, params)

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

func (cc AuthorityController) NameItems(c *gin.Context) {
	var params validator.AuthorityNameItemsValidator
	if err := c.ShouldBind(&params); err != nil {
		cc.Out(c, nil, validator.Out(params, err))
		return
	}
	items, err := services.AuthorityService{}.NameItems(params.ProjectId, params.Type)
	if err != nil {
		cc.Fail(c)
		return
	}
	cc.Ok(c, items)
}

func (cc AuthorityController) Info(c *gin.Context) {
	var params validator.AuthorityInfoValidator
	if err := c.ShouldBind(&params); err != nil {
		cc.Out(c, nil, validator.Out(params, err))
		return
	}
	items, err := services.AuthorityService{}.Info(params.Id)
	if err != nil {
		cc.Fail(c)
		return
	}
	cc.Ok(c, items)
}
