package admin

import (
	controller "base-service/app/controllers"
	"base-service/app/services"
	"base-service/app/validator"
	"github.com/gin-gonic/gin"
)

type LogController struct {
	controller.BaseController
}

func (cc LogController) Login(c *gin.Context) {
	var params validator.LogLoginItemsValidator
	if err := c.ShouldBind(&params); err != nil {
		cc.Out(c, nil, validator.Out(params, err))
		return
	}
	items, err := services.LogService{}.Login(params)

	if err != nil {
		cc.Fail(c)
		return
	}
	cc.Ok(c, items)
}

func (cc LogController) Action(c *gin.Context) {
	var params validator.LogActionItemsValidator
	if err := c.ShouldBind(&params); err != nil {
		cc.Out(c, nil, validator.Out(params, err))
		return
	}
	items, err := services.LogService{}.Action(params)
	if err != nil {
		cc.Fail(c)
		return
	}
	cc.Ok(c, items)
}

func (cc LogController) Add(c *gin.Context) {
	var params validator.LogActionValidator
	if err := c.ShouldBind(&params); err != nil {
		cc.Out(c, nil, validator.Out(params, err))
		return
	}
	id, editErr := services.LogService{}.ApiAdd(c, params)
	if editErr != nil {
		cc.Error(c, editErr)
		return
	}
	cc.Ok(c, id)
}
