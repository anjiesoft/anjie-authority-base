package middlewares

import (
	"base-service/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

// JwtMiddleWare 定义一个中间件函数来验证JWT令牌
func JwtMiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 假设Token放在Header的Authorization中，并使用Bearer开头
		token := c.GetHeader("Authorization")
		if token == "" || !strings.HasPrefix(token, "Bearer ") {
			c.JSON(http.StatusOK, utils.ErrorNoLogin)
			c.Abort()
			return
		}
		// 去掉Bearer
		token = token[7:]
		// 使用之前定义好的解析JWT的函数 ParseToken 来解析它
		jwt := utils.Jwt{}
		userInfo, err := jwt.ParseToken(token, true)
		if err != nil {
			if err.Error() == "signature is invalid" {
				c.JSON(http.StatusOK, utils.ErrorJwtError)
			} else if strings.HasPrefix(err.Error(), "token is expired ") {
				c.JSON(http.StatusOK, utils.ErrorJwtExpired)
			} else {
				c.JSON(http.StatusOK, utils.Fail)
			}
			c.Abort()
			return
		}

		super := make(map[string]string)
		tmp := utils.GetConfigString("base.super_id")
		for _, v := range strings.Split(tmp, ",") {
			super[v] = v
		}

		//添加数据到全局中
		c.Set(utils.USERID, userInfo.Id)
		c.Set(utils.USERNAME, userInfo.Name)
		c.Set(utils.USERINFO, userInfo)
		c.Set(utils.SUPERID, super)
		c.Next()
	}

}
