package services

import (
	"base-service/app/dao"
	"base-service/app/models"
	"base-service/app/validator"
	"base-service/utils"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/mssola/user_agent"
	"strconv"
)

type LogService struct {
}

func (s LogService) Fail(c *gin.Context, reason, username string) {
	userAgentHeader := c.Request.UserAgent() // 获取User-Agent字符串
	ua := user_agent.New(userAgentHeader)    // 创建user_agent解析器实例
	name, version := ua.Browser()            // 获取浏览器信息
	data := dao.LoginLog{
		Ip:             c.ClientIP(),
		Username:       username,
		Reason:         reason,
		BrowserInfo:    userAgentHeader,
		BrowserName:    name,
		BrowserVersion: version,
		Status:         dao.LoginLogStatusFail,
		CreateTime:     utils.Now(),
	}

	models.LoginLogModel{}.Init().Create(&data)
}

func (s LogService) Ok(c *gin.Context, info utils.JwtInfo) {
	userAgentHeader := c.Request.UserAgent() // 获取User-Agent字符串
	ua := user_agent.New(userAgentHeader)    // 创建user_agent解析器实例
	name, version := ua.Browser()            // 获取浏览器信息
	data := dao.LoginLog{
		Ip:             c.ClientIP(),
		Name:           info.Name,
		Username:       info.Username,
		AdminId:        info.Id,
		BrowserInfo:    userAgentHeader,
		BrowserName:    name,
		BrowserVersion: version,
		Status:         dao.LoginLogStatusOk,
		CreateTime:     utils.Now(),
	}

	err := models.LoginLogModel{}.Init().Create(&data)
	if err != nil {
		utils.Logger.Error(err)
		return
	}
}

func (s LogService) Login(validator validator.LogLoginItemsValidator) (dao.LoginLogItems, error) {
	model := models.LoginLogModel{}.Init()
	w := models.Params{}
	w.Page = 1
	w.Size = dao.SIZE
	w.Order = "id desc"

	if validator.Page > 0 {
		w.Page = validator.Page
	}
	if validator.Size > 0 {
		w.Size = validator.Size
	}
	w.Like = make(map[string]string)
	w.Other = make(map[string]string)
	eq := make(map[string]string)
	if len(validator.Username) > 0 {
		w.Like["username LIKE ? "] = "%%" + validator.Username + "%%"
	}

	if len(validator.Ip) > 0 {
		w.Like["ip LIKE ? "] = "%%" + validator.Ip + "%%"
	}

	if len(validator.Time) > 0 {
		w.Other["create_time > ?"] = validator.Time[0]
		w.Other["create_time < ?"] = validator.Time[1]
	}

	if validator.Status > 0 {
		eq["status"] = strconv.Itoa(int(validator.Status))
	}

	if len(validator.Name) > 0 {
		w.Like["name LIKE ? "] = "%%" + validator.Name + "%%"
	}

	if validator.AdminId > 0 {
		eq["admin_id"] = strconv.Itoa(validator.AdminId)
	}

	w.Eq = eq

	var data dao.LoginLogItems
	var ids []dao.Count
	count, err := model.Count(w, &ids)
	if err != nil {
		utils.Logger.Error(err)
		return data, err
	}
	model.Page(w, &data.Items)
	data.Count = count
	return data, nil
}

// ApiAdd 创建
func (s LogService) ApiAdd(c *gin.Context, validator validator.LogActionValidator) (int, utils.Error) {
	adminId, _ := c.Get(utils.USERID)
	adminName, _ := c.Get(utils.USERNAME)

	now := utils.Now()
	edit := dao.ActionLog{
		AdminName:      adminName.(string),
		AdminId:        adminId.(int),
		Content:        validator.Content,
		Type:           validator.Type,
		ModuleName:     validator.ModuleName,
		Ip:             validator.Ip,
		BrowserInfo:    validator.BrowserInfo,
		BrowserName:    validator.BrowserName,
		BrowserVersion: validator.BrowserVersion,
		ProjectId:      validator.ProjectId,
		Reason:         validator.Reason,
		CreateTime:     now,
	}

	err := models.ActionLogModel{}.Init().Create(&edit)

	if err != nil {
		utils.Logger.Error(err)
		return 0, utils.Fail
	}

	return edit.Id, nil
}

// Add 创建
func (s LogService) Add(c *gin.Context, info dao.ActionAdd) (int, utils.Error) {
	adminId, _ := c.Get(utils.USERID)
	adminName, _ := c.Get(utils.USERNAME)
	userAgentHeader := c.Request.UserAgent() // 获取User-Agent字符串
	ua := user_agent.New(userAgentHeader)    // 创建user_agent解析器实例
	name, version := ua.Browser()            // 获取浏览器信息

	content := map[string]interface{}{
		"old_content": info.OldContent,
		"new_content": info.NewContent,
	}
	jsonData, _ := json.Marshal(content)

	if info.ProjectId < 1 {
		w := models.Params{Eq: map[string]string{"status": strconv.Itoa(int(dao.StatusOk))}}
		project := dao.ProjectNameItems{}
		models.ProjectModel{}.Init().GetOne(w, &project)
		info.ProjectId = project.Id
	}

	now := utils.Now()
	edit := dao.ActionLog{
		AdminName:      adminName.(string),
		AdminId:        adminId.(int),
		Content:        string(jsonData),
		Type:           info.Type,
		ModuleName:     info.ModuleName,
		Ip:             c.ClientIP(),
		BrowserInfo:    userAgentHeader,
		BrowserName:    name,
		BrowserVersion: version,
		ProjectId:      info.ProjectId,
		Reason:         info.Reason,
		CreateTime:     now,
	}

	err := models.ActionLogModel{}.Init().Create(&edit)

	if err != nil {
		utils.Logger.Error(err)
		return 0, utils.Fail
	}

	return edit.Id, nil
}

func (s LogService) Action(validator validator.LogActionItemsValidator) (dao.ActionItems, error) {
	model := models.ActionLogModel{}.Init()
	w := models.Params{}
	w.Page = 1
	w.Size = dao.SIZE
	w.Order = "id desc"

	if validator.Page > 0 {
		w.Page = validator.Page
	}
	if validator.Size > 0 {
		w.Size = validator.Size
	}
	w.Like = make(map[string]string)
	eq := make(map[string]string)
	w.Other = make(map[string]string)
	if len(validator.AdminName) > 0 {
		w.Like["admin_name LIKE ? "] = "%%" + validator.AdminName + "%%"
	}

	if len(validator.Ip) > 0 {
		w.Like["ip LIKE ? "] = "%%" + validator.Ip + "%%"
	}

	if len(validator.ModuleName) > 0 {
		w.Like["module_name LIKE ? "] = "%%" + validator.ModuleName + "%%"
	}

	if validator.Type > 0 {
		eq["type"] = strconv.Itoa(int(validator.Type))
	}

	if len(validator.Time) > 0 {
		w.Other["create_time > ?"] = validator.Time[0]
		w.Other["create_time < ?"] = validator.Time[1]
	}

	if validator.AdminId > 0 {
		eq["admin_id"] = strconv.Itoa(validator.AdminId)
	}

	if validator.ProjectId > 0 {
		eq["project_id"] = strconv.Itoa(validator.ProjectId)
	}

	w.Eq = eq

	var data dao.ActionItems
	var ids []dao.Count
	count, err := model.Count(w, &ids)
	if err != nil {
		utils.Logger.Error(err)
		return data, err
	}
	model.Page(w, &data.Items)
	var Type []dao.ActionType
	Type = append(Type, dao.ActionType{Id: dao.ActionTypeAdd, Name: "增加"})
	Type = append(Type, dao.ActionType{Id: dao.ActionTypeEdit, Name: "编辑"})
	Type = append(Type, dao.ActionType{Id: dao.ActionTypeStatus, Name: "修改状态"})
	Type = append(Type, dao.ActionType{Id: dao.ActionTypePassword, Name: "修改密码"})
	Type = append(Type, dao.ActionType{Id: dao.ActionTypeOut, Name: "强退"})
	Type = append(Type, dao.ActionType{Id: dao.ActionTypeDown, Name: "下载"})
	Type = append(Type, dao.ActionType{Id: dao.ActionTypeDel, Name: "删除"})
	data.Type = Type
	data.Count = count
	return data, nil
}
