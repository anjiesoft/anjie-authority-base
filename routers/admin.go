package routers

import (
	"base-service/app/controllers/admin"
	"base-service/middlewares"
	"github.com/gin-gonic/gin"
)

func AdminRouter(routers *gin.RouterGroup) {
	obj := admin.LoginController{}
	//登录
	routers.POST("/login", obj.Login)
	//置换token
	routers.POST("/refresh", obj.Refresh)

	//登录
	login := routers.Group("/login", middlewares.JwtMiddleWare())
	{
		//登录信息
		login.POST("/info", obj.Info)
		login.POST("/routes", obj.Routes)
	}

	//管理员组
	group := routers.Group("/admin", middlewares.JwtMiddleWare(),
		middlewares.VerifyMiddleWare())
	{
		obj := admin.AdminController{}
		//创建/编辑
		group.POST("/edit", obj.Edit)
		//修改密码
		group.POST("/password", obj.Password)
		//详情
		group.POST("/info", obj.Info)
		//列表
		group.POST("/items", obj.Items)
		//更改状态
		group.POST("/status", obj.Status)
		//修改自己的密码
		group.POST("/ownpwd", obj.OwnPassword)
	}

}
