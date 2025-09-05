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

type RoleService struct {
}

// Edit 创建/编辑
func (s RoleService) Edit(c *gin.Context, validator validator.RoleEditValidator) (int, utils.Error) {
	adminId, _ := c.Get(utils.USERID)
	adminName, _ := c.Get(utils.USERNAME)

	check, err := s.check(validator.Name, validator.ProjectId)
	if err != nil {
		return 0, err
	}

	now := utils.Now()
	edit := dao.Role{
		Name:       validator.Name,
		ProjectId:  validator.ProjectId,
		Rules:      s.getIds(validator.Rules, validator.ProjectId),
		Status:     dao.StatusOk,
		AdminId:    adminId.(int),
		AdminName:  adminName.(string),
		CreateTime: now,
		UpdateTime: now,
	}
	action := dao.ActionAdd{
		Type:       dao.ActionTypeAdd,
		ModuleName: dao.ModuleNameRole,
	}

	model := models.RoleModel{}.Init()
	id := validator.Id
	var editErr error
	if id > 0 {
		if check.Id > 0 && check.Id != id {
			return 0, utils.ErrorExist
		}

		w := models.Params{Eq: map[string]string{"id": strconv.Itoa(validator.Id)}}
		info := dao.Role{}
		InfoErr := model.GetOne(w, &info)
		if InfoErr != nil {
			utils.Logger.Error(InfoErr)
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
		_, editErr = model.Edit(w, edit)
	} else {
		if check.Id > 0 {
			return 0, utils.ErrorExist
		}
		edit.CreateTime = now
		editErr = model.Create(&edit)
		id = edit.Id
	}

	if editErr != nil {
		utils.Logger.Error(editErr)
		return 0, utils.Fail
	}

	action.NewContent = edit
	LogService{}.Add(c, action)
	return id, nil
}

func (s RoleService) getIds(ids []int, projectId int) string {
	var roles []string
	if len(ids) > 0 {
		for _, role := range ids {
			roles = append(roles, strconv.Itoa(role))
		}
	}

	var items []dao.Authority
	w := models.Params{
		Eq: map[string]string{"status": strconv.Itoa(int(dao.StatusOk)), "project_id": strconv.Itoa(projectId)},
		In: map[string][]string{"id in ?": roles},
	}
	w.Order = "level ASC, parent_id ASC, sort ASC"
	err := models.AuthorityModel{}.Init().Items(w, &items)
	if err != nil {
		return ""
	}

	var idsArr []string
	for _, item := range items {
		idsArr = append(idsArr, strconv.Itoa(item.Id))
		tmp := strings.Split(item.ParentIds, ",")
		idsArr = append(idsArr, tmp...)
	}

	var retIdsArr []string
	keys := make(map[string]bool)
	for _, id := range idsArr {
		if id != "0" {
			if _, value := keys[id]; !value {
				keys[id] = true
				retIdsArr = append(retIdsArr, id)
			}
		}
	}

	return strings.Join(retIdsArr, ",")
}

// Status 修改状态
func (s RoleService) Status(c *gin.Context, validator validator.RoleStatusValidator) (bool, utils.Error) {
	info := dao.RoleStatus{}
	w := models.Params{Eq: map[string]string{"id": strconv.Itoa(validator.Id)}}
	model := models.RoleModel{}.Init()
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
		ModuleName: dao.ModuleNameRole,
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

// Rule 编辑权限
func (s RoleService) Rule(c *gin.Context, id int, rules string) (bool, utils.Error) {
	info := dao.RoleRule{}
	w := models.Params{Eq: map[string]string{"id": strconv.Itoa(id)}}
	model := models.RoleModel{}.Init()
	err := model.GetOne(w, &info)
	if err != nil {
		utils.Logger.Error(err)
		return false, utils.Fail
	}
	if info.Id < 1 {
		return false, utils.ErrorNotFund
	}
	action := dao.ActionAdd{
		Type:       dao.ActionTypeEdit,
		ModuleName: dao.ModuleNameRole,
		OldContent: info,
	}
	now := utils.Now()
	info.Rules = rules
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

func (s RoleService) Info(id int) (dao.RoleInfo, error) {
	role := dao.Role{}
	model := models.RoleModel{}.Init()
	w := models.Params{Eq: map[string]string{"id": strconv.Itoa(id)}}
	err := model.GetOne(w, &role)

	var rules []int
	ruleArrs := strings.Split(role.Rules, ",")
	for _, rule := range ruleArrs {
		r, _ := strconv.Atoi(rule)
		rules = append(rules, r)
	}

	info := dao.RoleInfo{
		Id:        role.Id,
		Name:      role.Name,
		Status:    role.Status,
		ProjectId: role.ProjectId,
		Reason:    role.Reason,
		Rules:     rules,
	}

	return info, err
}

func (s RoleService) Items(validator validator.RoleItemsValidator) ([]dao.RoleItems, error) {
	model := models.RoleModel{}.Init()
	w := models.Params{}
	w.Page = 1
	w.Size = dao.SIZE
	w.Order = "id asc"

	if validator.Page > 0 {
		w.Page = validator.Page
	}
	if validator.Size > 0 {
		w.Size = validator.Size
	}
	w.Like = make(map[string]string)
	w.Eq = make(map[string]string)
	w.Eq["project_id"] = strconv.Itoa(validator.ProjectId)
	if validator.Status > 0 {
		w.Eq["status"] = strconv.Itoa(int(validator.Status))
	}
	if len(validator.Name) > 0 {
		w.Like["name LIKE ? "] = "%%" + validator.Name + "%%"
	}

	var data []dao.Role
	err := model.Items(w, &data)
	if err != nil {
		utils.Logger.Error(err)
		return nil, err
	}

	var items []dao.RoleItems
	rulesInfos := s.getRulesInfo(data)
	for i := 0; i < len(data); i++ {
		var rules []string
		for _, rule := range strings.Split(data[i].Rules, ",") {
			if _, ok := rulesInfos[rule]; ok {
				rules = append(rules, rulesInfos[rule])
			}
		}
		items = append(items, dao.RoleItems{
			Role:      data[i],
			RuleItems: rules,
		})
	}

	return items, nil
}

func (s RoleService) getRulesInfo(data []dao.Role) map[string]string {
	var ids []string
	for i := 0; i < len(data); i++ {
		ids = append(ids, strings.Split(data[i].Rules, ",")...)
	}

	w := models.Params{
		In:    map[string][]string{"id in ?": ids},
		Other: map[string]string{"type != ?": strconv.Itoa(dao.TypeDir)},
	}
	var items []dao.AuthorityNames
	retData := make(map[string]string)

	err := models.AuthorityModel{}.Init().Items(w, &items)
	if err != nil {
		return retData
	}

	for i := 0; i < len(items); i++ {
		retData[strconv.Itoa(items[i].Id)] = items[i].Name
	}

	return retData
}

func (s RoleService) NameItems(name string) ([]dao.RoleProjectNameItems, error) {
	projects, pErr := ProjectService{}.NameItems("")
	if pErr != nil {
		utils.Logger.Error(pErr)
		return nil, pErr
	}

	var items []dao.RoleProjectNameItems
	var roles []dao.Role
	model := models.RoleModel{}.Init()
	w := models.Params{}
	w.Eq = map[string]string{"status": strconv.Itoa(int(dao.StatusOk))}
	w.Like = make(map[string]string)
	if len(name) > 0 {
		w.Like["name LIKE ? "] = "%%" + name + "%%"
	}
	err := model.Items(w, &roles)
	tmp := map[int][]dao.RoleNameItems{}
	for _, item := range roles {
		tmp[item.ProjectId] = append(tmp[item.ProjectId], dao.RoleNameItems{
			Name: item.Name,
			Id:   item.Id,
		})
	}
	for _, pro := range projects {
		if _, ok := tmp[pro.Id]; ok {
			items = append(items, dao.RoleProjectNameItems{
				Id:    pro.Id,
				Name:  pro.Name,
				Items: tmp[pro.Id],
			})
		}
	}
	return items, err
}

func (s RoleService) GetByIds(ids []string, projectId int) ([]dao.RoleRulesItems, error) {
	w := models.Params{
		Eq: map[string]string{"status": strconv.Itoa(int(dao.StatusOk)), "project_id": strconv.Itoa(projectId)},
		In: map[string][]string{"id in ?": ids},
	}
	var info []dao.RoleRulesItems
	err := models.RoleModel{}.Init().Items(w, &info)
	return info, err
}

func (s RoleService) check(name string, id int) (dao.Role, utils.Error) {
	w := models.Params{Eq: map[string]string{"name": name, "project_id": strconv.Itoa(id)}}
	info := dao.Role{}
	err := models.RoleModel{}.Init().GetOne(w, &info)
	if err != nil {
		utils.Logger.Error(err)
		return info, utils.Fail
	}
	return info, nil
}
