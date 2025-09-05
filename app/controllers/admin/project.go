package admin

import (
	controller "base-service/app/controllers"
	"base-service/app/dao"
	"base-service/app/services"
	"base-service/app/validator"
	"base-service/utils"
	"github.com/gin-gonic/gin"
)

type ProjectController struct {
	controller.BaseController
}

func (cc ProjectController) Edit(c *gin.Context) {
	var params validator.ProjectEditValidator
	if err := c.ShouldBind(&params); err != nil {
		cc.Out(c, nil, validator.Out(params, err))
		return
	}
	id, editErr := services.ProjectService{}.Edit(c, params)
	if editErr != nil {
		cc.Error(c, editErr)
		return
	}
	cc.Ok(c, id)
}

func (cc ProjectController) Items(c *gin.Context) {
	var params validator.ProjectItemsValidator
	if err := c.ShouldBind(&params); err != nil {
		cc.Out(c, nil, validator.Out(params, err))
		return
	}
	items, err := services.ProjectService{}.Items(params)
	if err != nil {
		cc.Fail(c)
		return
	}
	cc.Ok(c, items)
}

func (cc ProjectController) NameItems(c *gin.Context) {
	var params validator.ProjectNameItemsValidator
	if err := c.ShouldBind(&params); err != nil {
		cc.Out(c, nil, validator.Out(params, err))
		return
	}
	items, err := services.ProjectService{}.NameItems(params.Name)
	if err != nil {
		cc.Fail(c)
		return
	}
	cc.Ok(c, items)
}

func (cc ProjectController) Status(c *gin.Context) {
	var params validator.ProjectStatusValidator
	if err := c.ShouldBind(&params); err != nil {
		cc.Out(c, nil, validator.Out(params, err))
		return
	}
	if params.Status == dao.StatusFail && len(params.Reason) < 1 {
		cc.Error(c, utils.ErrorReasonParams)
		return
	}

	is, pErr := services.ProjectService{}.Status(c, params)

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

func (cc ProjectController) Info(c *gin.Context) {

	var params validator.ProjectInfoValidator

	if err := c.ShouldBind(&params); err != nil {
		cc.Out(c, nil, validator.Out(params, err))
		return
	}
	info, err := services.ProjectService{}.Info(params.Id)
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
