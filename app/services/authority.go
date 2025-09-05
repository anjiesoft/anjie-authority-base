package services

import (
	"base-service/app/dao"
	"base-service/app/models"
	"base-service/app/validator"
	"base-service/utils"
	"github.com/gin-gonic/gin"
	"strconv"
	"strings"
)

type AuthorityService struct {
}

// Edit 创建
func (s AuthorityService) Edit(c *gin.Context, validator validator.AuthorityEditValidator) (int, utils.Error) {
	edit, dErr := s.getData(c, validator)
	if dErr != nil {
		return 0, dErr
	}
	model := models.AuthorityModel{}.Init()
	var err error
	id := validator.Id
	action := dao.ActionAdd{
		Type:       dao.ActionTypeAdd,
		ModuleName: dao.ModuleNameAuthority,
	}
	if id > 0 {
		w := models.Params{Eq: map[string]string{"id": strconv.Itoa(validator.Id)}}
		info := dao.Authority{}
		err = model.GetOne(w, &info)
		if err != nil {
			utils.Logger.Error(err)
			return 0, utils.Fail
		}
		if info.Id < 1 {
			return 0, utils.ErrorNotFund
		}
		action.OldContent = info

		edit.Id = info.Id
		edit.Type = info.Type
		edit.Sort = info.Sort
		edit.ProjectId = info.ProjectId
		edit.CreateTime = info.CreateTime
		edit.Status = info.Status
		action.Type = dao.ActionTypeEdit
		w = models.Params{Eq: map[string]string{"id": strconv.Itoa(info.Id)}}
		_, err = model.Edit(w, edit)
	} else {
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
func (s AuthorityService) Status(c *gin.Context, validator validator.AuthorityStatusValidator) (bool, utils.Error) {
	info := dao.AuthorityStatus{}
	w := models.Params{Eq: map[string]string{"id": strconv.Itoa(validator.Id)}}
	model := models.AuthorityModel{}.Init()
	err := model.GetOne(w, &info)
	if err != nil {
		utils.Logger.Error(err)
		return false, utils.Fail
	}
	if info.Id < 1 {
		return false, utils.ErrorNotFund
	}

	countW := models.Params{Eq: map[string]string{
		"project_id": strconv.Itoa(info.ProjectId),
		"status":     strconv.Itoa(int(dao.StatusOk)),
		"parent_id":  strconv.Itoa(validator.Id),
	}}
	infos := dao.AuthorityStatus{}
	count, err := model.Count(countW, &infos)
	if err != nil {
		utils.Logger.Error(err)
		return false, utils.Fail
	}

	if count > 0 {
		return false, utils.ErrorAuthorityChange
	}

	action := dao.ActionAdd{
		Type:       dao.ActionTypeStatus,
		ModuleName: dao.ModuleNameAuthority,
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

func (s AuthorityService) Items(validator validator.AuthorityItemsValidator) ([]dao.AuthorityItems, error) {
	model := models.AuthorityModel{}.Init()
	w := models.Params{}
	w.Order = "level ASC, parent_id ASC, sort ASC"

	w.Eq = map[string]string{"project_id": strconv.Itoa(validator.ProjectId)}
	var data []dao.Authority
	err := model.Items(w, &data)
	if err != nil {
		utils.Logger.Error(err)
		return nil, err
	}

	items := s.getList(data, 0)

	return items, nil
}

func (s AuthorityService) NameItems(projectId int, Type uint8) ([]dao.AuthorityNamesItems, error) {
	var data []dao.AuthorityNames
	model := models.AuthorityModel{}.Init()
	w := models.Params{}
	w.Order = "level ASC, parent_id ASC, sort ASC"
	w.Eq = map[string]string{
		"project_id": strconv.Itoa(projectId),
		"status":     strconv.Itoa(int(dao.StatusOk)),
	}
	if Type == dao.QuerySortLIstWeb {
		w.In = map[string][]string{"type in ?": {strconv.Itoa(dao.TypeDir), strconv.Itoa(dao.TypeWeb)}}
	} else if Type == dao.QuerySortListNotButton {
		w.In = map[string][]string{"type in ?": {strconv.Itoa(dao.TypeDir), strconv.Itoa(dao.TypeWeb)}}
	}

	err := model.Items(w, &data)

	items := s.getNameList(data, 0)
	return items, err
}

func (s AuthorityService) getData(c *gin.Context, validator validator.AuthorityEditValidator) (dao.Authority, utils.Error) {
	adminId, _ := c.Get(utils.USERID)
	adminName, _ := c.Get(utils.USERNAME)

	now := utils.Now()
	edit := dao.Authority{
		Name:           validator.Name,
		Type:           validator.Type,
		Status:         dao.StatusOk,
		Remarks:        validator.Remarks,
		Icon:           validator.Icon,
		Identification: validator.Identification,
		ProjectId:      validator.ProjectId,
		ParentId:       validator.ParentId,
		Api:            validator.Api,
		ViewPath:       validator.ViewPath,
		Path:           validator.Path,
		IsShow:         validator.IsShow,
		AdminId:        adminId.(int),
		AdminName:      adminName.(string),
		CreateTime:     now,
		UpdateTime:     now,
	}

	edit.Path = "/" + strings.TrimLeft(edit.Path, "/")

	utils.Logger.Error(edit.Path)
	if edit.ParentId < 1 {
		edit.ParentId = 0
	}

	check := s.check(validator.Id, map[string]string{
		"project_id":     strconv.Itoa(edit.ProjectId),
		"identification": edit.Identification,
	})
	if !check {
		return edit, utils.NewError(utils.ErrorMissingParams.GetCode(), "权限标识已存在")
	}

	if edit.Type == dao.TypeWeb {
		if len(edit.Api) < 1 {
			return edit, utils.NewError(utils.ErrorMissingParams.GetCode(), "API不能为空")
		}
		if len(edit.ViewPath) < 1 {
			return edit, utils.NewError(utils.ErrorMissingParams.GetCode(), "组件路径不能为空")
		}
		check := s.check(validator.Id, map[string]string{"path": edit.Path})
		if !check {
			return edit, utils.NewError(utils.ErrorMissingParams.GetCode(), "PATH已存在")
		}
	} else if edit.Type == dao.TypeButton {
		edit.Icon = ""
		if len(edit.Api) < 1 {
			return edit, utils.NewError(utils.ErrorMissingParams.GetCode(), "API不能为空")
		}
		edit.Path = dao.NoData
		edit.ViewPath = dao.NoData
	} else if edit.Type == dao.TypeDir {
		edit.Api = dao.NoData
		edit.ViewPath = dao.NoData
		check := s.check(validator.Id, map[string]string{"path": edit.Path})
		if !check {
			return edit, utils.NewError(utils.ErrorMissingParams.GetCode(), "PATH已存在")
		}
	}

	if edit.ParentId == 0 {
		edit.ParentIds = "0"
	} else {
		w := models.Params{Eq: map[string]string{"id": strconv.Itoa(edit.ParentId)}}
		info := dao.Authority{}
		err := models.AuthorityModel{}.Init().GetOne(w, &info)
		if err != nil {
			return edit, nil
		}
		if info.Id < 1 {
			return edit, utils.ErrorNotFund
		}
		edit.ParentIds = info.ParentIds + "," + strconv.Itoa(edit.ParentId)
	}

	idsArr := strings.SplitAfter(edit.ParentIds, ",")
	edit.Level = uint8(len(idsArr))

	return edit, nil
}

func (s AuthorityService) check(id int, check map[string]string) bool {
	w := models.Params{Eq: check}
	info := dao.AuthorityCheck{}
	_ = models.AuthorityModel{}.Init().GetOne(w, &info)
	if id > 0 && info.Id == id {
		return true
	}
	if info.Id > 0 {
		return false
	}

	return true
}

// 获取层级的下接数据 列表使用
func (s AuthorityService) getList(data []dao.Authority, pid int) []dao.AuthorityItems {
	var dataArr []dao.AuthorityItems
	for _, v := range data {
		if v.ParentId == pid {
			// 这里可以理解为每次都从最原始的数据里面找出相对就的ID进行匹配，直到找不到就返回
			children := s.getList(data, v.Id)
			node := dao.AuthorityItems{
				Id:             v.Id,
				Name:           v.Name,
				Sort:           v.Sort,
				ParentId:       v.ParentId,
				ParentIds:      v.ParentIds,
				CreateTime:     v.CreateTime,
				UpdateTime:     v.UpdateTime,
				Remarks:        v.Remarks,
				Status:         v.Status,
				Level:          v.Level,
				Children:       children,
				Api:            v.Api,
				Path:           v.Path,
				ViewPath:       v.ViewPath,
				Identification: v.Identification,
				Type:           v.Type,
				IsShow:         v.IsShow,
				Icon:           v.Icon,
				Reason:         v.Reason,
				AdminName:      v.AdminName,
				AdminId:        v.AdminId,
			}
			dataArr = append(dataArr, node)
		}
	}
	return dataArr
}

// 获取层级的下接数据 列表使用
func (s AuthorityService) getNameList(data []dao.AuthorityNames, pid int) []dao.AuthorityNamesItems {
	var dataArr []dao.AuthorityNamesItems
	for _, v := range data {
		if v.ParentId == pid {
			// 这里可以理解为每次都从最原始的数据里面找出相对就的ID进行匹配，直到找不到就返回
			children := s.getNameList(data, v.Id)
			node := dao.AuthorityNamesItems{
				Id:       v.Id,
				ParentId: v.ParentId,
				Name:     v.Name,
				Level:    v.Level,
				Type:     v.Type,
				Children: children,
			}
			dataArr = append(dataArr, node)
		}
	}
	return dataArr
}

func (s AuthorityService) Info(id int) (dao.AuthorityInfo, error) {
	info := dao.AuthorityInfo{}
	model := models.AuthorityModel{}.Init()
	w := models.Params{Eq: map[string]string{"id": strconv.Itoa(id)}}
	err := model.GetOne(w, &info)
	return info, err
}
