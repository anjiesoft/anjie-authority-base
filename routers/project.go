package routers

import (
	"base-service/app/controllers/admin"
	"base-service/middlewares"
	"github.com/gin-gonic/gin"
)

func ProjectRouter(routers *gin.RouterGroup) {
	group := routers.Group("/project", middlewares.JwtMiddleWare(),
		middlewares.VerifyMiddleWare())
	{
		obj := admin.ProjectController{}
		//创建/编辑
		group.POST("/edit", obj.Edit)
		//详情
		group.POST("/info", obj.Info)
		//列表
		group.POST("/items", obj.Items)
		//更改状态
		group.POST("/status", obj.Status)
		//名称列表
		group.POST("/name_items", obj.NameItems)
	}
}
