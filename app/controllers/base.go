package controller

import (
	"base-service/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

var PostJsonData map[string]interface{}

type BaseController struct{}

// Ok 成功的输出结果
func (b *BaseController) Ok(c *gin.Context, data any) {
	b.Out(c, data, utils.OK)
}

// Fail 基础错误的输出结果
func (b *BaseController) Fail(c *gin.Context) {
	b.Out(c, nil, utils.Fail)
}

// Error 错误的输出结果
func (b *BaseController) Error(c *gin.Context, cm utils.Error) {
	b.Out(c, nil, cm)
}

// Out 输出结果
func (b *BaseController) Out(c *gin.Context, data any, cm utils.Error) {
	cm.WithData(data)

	c.JSON(http.StatusOK, cm)
}

func (b *BaseController) GetPost(key string) string {
	if values, ok := PostJsonData[key]; ok {
		switch values.(type) {
		case string:
			return values.(string)
		case float64:
			ft := values.(float64)
			return strconv.FormatFloat(ft, 'f', -1, 64)
		}
	}
	return ""
}

func (b *BaseController) GetPostDef(key, def string) string {
	if values, ok := PostJsonData[key]; ok {
		switch values.(type) {
		case string:
			return values.(string)
		case float64:
			ft := values.(float64)
			return strconv.FormatFloat(ft, 'f', -1, 64)
		}
	}
	return def
}

func (b *BaseController) GetPostMap(params map[string]string) map[string]interface{} {
	retData := make(map[string]interface{})
	for k, v := range params {
		if value, ok := PostJsonData[k]; ok {
			retData[k] = value
		} else {
			retData[k] = v
		}
	}

	return retData
}
