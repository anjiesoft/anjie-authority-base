package middlewares

import (
	"base-service/app/dao"
	"base-service/app/models"
	"base-service/app/services"
	"base-service/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"strings"
)

type verify struct {
	ProjectId int    `json:"project_id"`
	Auth      string `json:"auth"`
	Path      string `json:"path"`
}

// VerifyMiddleWare 定义一个中间件函数来验证权限
func VerifyMiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		err := check(c)
		if err != nil {
			c.JSON(http.StatusOK, err)
			c.Abort()
			return
		}
		c.Next()
	}

}

func getVerify(c *gin.Context) (*verify, utils.Error) {
	project := c.GetHeader("Project")
	projectId, err := strconv.Atoi(project)
	if err != nil || projectId < 1 {
		return nil, utils.ErrorMissingProjectId
	}

	auth := c.GetHeader("Auth")
	if auth == "" {
		return nil, utils.ErrorMissingIdentification
	}

	path := c.FullPath()
	if path == "" {
		return nil, utils.ErrorAuthority
	}
	params := verify{
		ProjectId: projectId,
		Auth:      auth,
		Path:      strings.Trim(path, "/"),
	}

	return &params, nil
}

func check(c *gin.Context) utils.Error {
	adminId, _ := c.Get(utils.USERID)
	super, _ := c.Get(utils.SUPERID)
	//如果是超级管理员
	if _, ok := super.(map[string]string)[strconv.Itoa(adminId.(int))]; ok {
		return nil
	}

	params, err := getVerify(c)
	if err != nil {
		return err
	}

	w := models.Params{}
	w.Eq = map[string]string{
		"project_id":     strconv.Itoa(params.ProjectId),
		"identification": params.Auth,
		"status":         strconv.Itoa(int(dao.StatusOk)),
	}
	info := dao.AuthorityCheck{}
	models.AuthorityModel{}.Init().GetOne(w, &info)

	if info.Id < 0 {
		return utils.ErrorAuthority
	}
	if params.Path != strings.Trim(info.Api, "/") {
		return utils.ErrorAuthority
	}

	auths := services.RoutesService{}.GetAuthority(adminId.(int), params.ProjectId)

	if len(auths) == 0 {
		return utils.ErrorAuthority
	}
	if _, ok := auths[params.Auth]; !ok {
		return utils.ErrorAuthority
	}
	return nil
}
