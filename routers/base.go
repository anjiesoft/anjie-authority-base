package routers

import (
	"github.com/gin-gonic/gin"
)

func BaseRouter(c *gin.Engine) {
	routers := c.Group("/admin")
	{
		AdminRouter(routers)
		RoleRouter(routers)
		ProjectRouter(routers)
		LogRouter(routers)
		AuthorityRouter(routers)
	}
}
