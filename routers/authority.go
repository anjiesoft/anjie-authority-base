package routers

import (
	"base-service/app/controllers/admin"
	"base-service/middlewares"
	"github.com/gin-gonic/gin"
)

func AuthorityRouter(routers *gin.RouterGroup) {
	group := routers.Group("/authority", middlewares.JwtMiddleWare(),
		middlewares.VerifyMiddleWare())
	{
		obj := admin.AuthorityController{}
		//创建/编辑
		group.POST("/edit", obj.Edit)
		//列表
		group.POST("/items", obj.Items)
		//详情
		group.POST("/info", obj.Info)
		//更改状态
		group.POST("/status", obj.Status)
		//编辑权限
		group.POST("/name_items", obj.NameItems)
	}
}
