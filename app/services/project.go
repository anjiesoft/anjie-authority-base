package services

import (
	"base-service/app/dao"
	"base-service/app/models"
	"base-service/app/validator"
	"base-service/utils"
	"github.com/gin-gonic/gin"
	"strconv"
)

type ProjectService struct {
}

// Edit 创建
func (s ProjectService) Edit(c *gin.Context, validator validator.ProjectEditValidator) (int, utils.Error) {
	adminId, _ := c.Get(utils.USERID)
	adminName, _ := c.Get(utils.USERNAME)

	now := utils.Now()
	edit := dao.Project{
		Name:       validator.Name,
		Logo:       validator.Logo,
		Remarks:    validator.Remarks,
		Status:     dao.StatusOk,
		AdminId:    adminId.(int),
		AdminName:  adminName.(string),
		CreateTime: now,
		UpdateTime: now,
	}
	action := dao.ActionAdd{
		Type:       dao.ActionTypeAdd,
		ModuleName: dao.ModuleNameProject,
	}

	model := models.ProjectModel{}.Init()
	var err error
	id := validator.Id
	if id > 0 {
		w := models.Params{Eq: map[string]string{"id": strconv.Itoa(validator.Id)}}
		info := dao.Project{}
		err = model.GetOne(w, &info)
		if err != nil {
			utils.Logger.Error(err)
			return 0, utils.Fail
		}
		if info.Id < 1 {
			return 0, utils.ErrorNotFund
		}
		action.OldContent = info
		action.Type = dao.ActionTypeEdit
		edit.Id = info.Id
		edit.UpdateTime = now
		edit.CreateTime = info.CreateTime
		edit.Status = info.Status
		w = models.Params{Eq: map[string]string{"id": strconv.Itoa(info.Id)}}
		_, err = model.Edit(w, edit)
	} else {
		edit.CreateTime = now
		err = model.Create(&edit)
		id = edit.Id
	}

	if err != nil {
		utils.Logger.Error(err)
		return 0, utils.Fail
	}

	action.NewContent = edit
	LogService{}.Add(c, action)
	return id, nil
}

// Status 修改状态
func (s ProjectService) Status(c *gin.Context, validator validator.ProjectStatusValidator) (bool, utils.Error) {
	info := dao.ProjectStatus{}
	w := models.Params{Eq: map[string]string{"id": strconv.Itoa(validator.Id)}}
	model := models.ProjectModel{}.Init()
	err := model.GetOne(w, &info)
	if err != nil {
		utils.Logger.Error(err)
		return false, utils.Fail
	}
	if info.Id < 1 {
		return false, utils.ErrorNotFund
	}
	action := dao.ActionAdd{
		Type:       dao.ActionTypeStatus,
		ModuleName: dao.ModuleNameProject,
		OldContent: info,
		Reason:     validator.Reason,
	}

	now := utils.Now()
	info.Status = validator.Status
	info.Reason = validator.Reason
	info.UpdateTime = now
	adminId, _ := c.Get(utils.USERID)
	adminName, _ := c.Get(utils.USERNAME)
	info.AdminId = adminId.(int)
	info.AdminName = adminName.(string)
	num, editErr := model.Edit(w, info)
	if editErr != nil {
		utils.Logger.Error(editErr)
		return false, utils.Fail
	}

	action.NewContent = info
	LogService{}.Add(c, action)
	return num > 0, nil
}

func (s ProjectService) Info(id int) (dao.Project, error) {
	info := dao.Project{}
	model := models.ProjectModel{}.Init()
	w := models.Params{Eq: map[string]string{"id": strconv.Itoa(id)}}
	err := model.GetOne(w, &info)
	return info, err
}

func (s ProjectService) Items(validator validator.ProjectItemsValidator) ([]dao.Project, error) {
	model := models.ProjectModel{}.Init()
	w := models.Params{}
	w.Page = 1
	w.Size = dao.SIZE
	w.Order = "id desc"
	w.Like = make(map[string]string)
	if validator.Status > 0 {
		w.Eq = map[string]string{"status": strconv.Itoa(int(validator.Status))}
	}
	if len(validator.Name) > 0 {
		w.Like["name LIKE ? "] = "%%" + validator.Name + "%%"
	}
	var data []dao.Project
	err := model.Items(w, &data)
	if err != nil {
		utils.Logger.Error(err)
		return data, err
	}
	return data, nil
}

func (s ProjectService) NameItems(name string) ([]dao.ProjectNameItems, error) {
	var items []dao.ProjectNameItems
	model := models.ProjectModel{}.Init()
	w := models.Params{}
	w.Eq = map[string]string{"status": strconv.Itoa(int(dao.StatusOk))}
	w.Like = make(map[string]string)
	if len(name) > 0 {
		w.Like["name LIKE ? "] = "%%" + name + "%%"
	}
	err := model.Items(w, &items)
	return items, err
}
