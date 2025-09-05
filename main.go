package main

import (
	"base-service/routers"
	"base-service/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	env := utils.GetConfigString("base.environment")
	if len(env) > 0 {
		// 获取默认的 gin Engine，Engine 中包含了所有路由处理的接口
		gin.SetMode(env)
	}
	r := gin.Default()
	r.Use(Cors())
	//开启 http2
	r.UseH2C = true
	routers.BaseRouter(r)

	// 监听端口默认为8080
	r.Run(":8085")

}

func Cors() gin.HandlerFunc {
	return func(context *gin.Context) {
		method := context.Request.Method

		context.Header("Access-Control-Allow-Origin", "*")
		context.Header("Access-Control-Allow-Credentials", "true")
		context.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		context.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")
		context.Header("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token, Authorization, Refresh_Authorization, Project, Auth")
		if method == "OPTIONS" {
			context.AbortWithStatus(http.StatusNoContent)
		}
		context.Next()
	}
}
