package routers

import (
	"base-service/app/controllers/admin"
	"base-service/middlewares"
	"github.com/gin-gonic/gin"
)

func LogRouter(routers *gin.RouterGroup) {
	group := routers.Group("/log", middlewares.JwtMiddleWare(),
		middlewares.VerifyMiddleWare())
	{
		log := admin.LogController{}
		//登录
		group.POST("/login", log.Login)
		//操作
		group.POST("/action", log.Action)
		//操作
		group.POST("/add", log.Add)
	}
}
