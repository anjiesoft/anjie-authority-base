package routers

import (
	"base-service/app/controllers/admin"
	"base-service/middlewares"
	"github.com/gin-gonic/gin"
)

func RoleRouter(routers *gin.RouterGroup) {
	group := routers.Group("/role", middlewares.JwtMiddleWare(),
		middlewares.VerifyMiddleWare())
	{
		obj := admin.RoleController{}
		//创建/编辑
		group.POST("/edit", obj.Edit)
		//详情
		group.POST("/info", obj.Info)
		//列表
		group.POST("/items", obj.Items)
		//更改状态
		group.POST("/status", obj.Status)
		////编辑权限
		//group.POST("/rule", obj.Rule)
		//名称列表
		group.POST("/name_items", obj.NameItems)
	}
}
