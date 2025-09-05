package services

import (
	"base-service/app/dao"
	"base-service/app/models"
	"base-service/app/validator"
	"base-service/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
	"strings"
)

type AdminService struct {
}

// Edit 创建/编辑
func (s AdminService) Edit(c *gin.Context, validator validator.AdminEditValidator) (int, utils.Error) {
	info, check := s.checkUserNameAndPhone(validator.Username, validator.Phone)
	if check != nil {
		return 0, check
	}
	model := models.AdminModel{}.Init()
	if validator.Id > 0 && info.Id < 1 {
		w := models.Params{Eq: map[string]string{"id": strconv.Itoa(validator.Id)}}
		info = dao.Admin{}
		err := model.GetOne(w, &info)
		if err != nil {
			utils.Logger.Error(err)
			return 0, utils.Fail
		}
		if info.Id < 0 {
			return 0, utils.ErrorNotFund
		}
	}
	//编辑
	if validator.Id > 0 {
		if info.Id < 1 {
			return 0, utils.ErrorNotFund
		}
		if info.Id != validator.Id {
			return 0, utils.ErrorExistUser
		}
	} else {
		if len(validator.Password) < 1 {
			return 0, utils.NewError(utils.ErrorMissingParams.GetCode(), "密码不能为空")

		}
		if info.Id > 0 {
			return 0, utils.ErrorExistUser
		}
	}
	adminId, _ := c.Get(utils.USERID)
	adminName, _ := c.Get(utils.USERNAME)
	var roles []string
	if len(validator.Roles) > 0 {
		for _, role := range validator.Roles {
			roles = append(roles, strconv.Itoa(role))
		}
	}
	admin := dao.Admin{
		Username:   validator.Username,
		Name:       validator.Name,
		Phone:      validator.Phone,
		Roles:      strings.Join(roles, ","),
		Email:      validator.Email,
		Avatar:     validator.Avatar,
		CreateTime: utils.Now(),
		AdminId:    adminId.(int),
		AdminName:  adminName.(string),
	}

	var err error
	id := validator.Id
	action := dao.ActionAdd{
		Type:       dao.ActionTypeAdd,
		ModuleName: dao.ModuleNameAdmin,
	}
	if info.Id > 0 {
		admin.UpdateTime = utils.Now()
		admin.Password = info.Password
		admin.Salt = info.Salt
		admin.Id = id
		admin.Status = info.Status
		w := models.Params{Eq: map[string]string{"id": strconv.Itoa(admin.Id)}}
		_, err = model.Edit(w, admin)
		action.Type = dao.ActionTypeEdit
		action.NewContent = admin
		action.OldContent = info
	} else {
		salt := utils.GetRandstring(5)
		pwd := fmt.Sprintf("%s%s", salt, validator.Password)
		pwd, _ = utils.EncryptPassword(pwd)
		admin.Password = pwd
		admin.CreateTime = utils.Now()
		admin.Salt = salt
		err = model.Init().Create(&admin)
		id = admin.Id
		action.NewContent = admin
	}

	LogService{}.Add(c, action)

	if err != nil {
		utils.Logger.Error(err)
		return 0, utils.Fail
	}

	return id, nil
}

// Status 修改状态
func (s AdminService) Status(c *gin.Context, params validator.AdminStatusValidator) (bool, utils.Error) {
	info := dao.AdminStatus{}
	w := models.Params{Eq: map[string]string{"id": strconv.Itoa(params.Id)}}
	model := models.AdminModel{}.Init()
	err := model.GetOne(w, &info)
	if err != nil {
		utils.Logger.Error(err)
		return false, utils.Fail
	}
	action := dao.ActionAdd{
		Type:       dao.ActionTypeStatus,
		ModuleName: dao.ModuleNameAdmin,
		OldContent: info,
		Reason:     params.Reason,
	}
	if info.Id < 1 {
		return false, utils.ErrorNotFund
	}

	info.Status = params.Status
	info.Reason = params.Reason
	info.UpdateTime = utils.Now()
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

// Password 修改列表中的密码
func (s AdminService) Password(c *gin.Context, id int, password string) (bool, utils.Error) {
	info := dao.AdminPassword{}
	w := models.Params{Eq: map[string]string{"id": strconv.Itoa(id)}}
	model := models.AdminModel{}.Init()
	err := model.GetOne(w, &info)
	if err != nil {
		utils.Logger.Error(err)
		return false, utils.Fail
	}
	if info.Id < 1 {
		return false, utils.ErrorNotFund
	}

	pwd := fmt.Sprintf("%s%s", info.Salt, password)
	pwd, _ = utils.EncryptPassword(pwd)
	info.Password = pwd
	info.UpdateTime = utils.Now()
	adminId, _ := c.Get(utils.USERID)
	adminName, _ := c.Get(utils.USERNAME)
	info.AdminId = adminId.(int)
	info.AdminName = adminName.(string)
	num, editErr := model.Edit(w, info)
	if editErr != nil {
		utils.Logger.Error(err)
		return false, utils.Fail
	}
	action := dao.ActionAdd{
		Type:       dao.ActionTypePassword,
		ModuleName: dao.ModuleNameAdmin,
		OldContent: info,
	}
	LogService{}.Add(c, action)
	return num > 0, nil
}

// OwnPassword 修改自己的密码
func (s AdminService) OwnPassword(c *gin.Context, password string) (bool, utils.Error) {
	info := dao.AdminPassword{}
	adminId, _ := c.Get(utils.USERID)
	w := models.Params{Eq: map[string]string{"id": strconv.Itoa(adminId.(int))}}
	model := models.AdminModel{}.Init()
	err := model.GetOne(w, &info)
	if err != nil {
		utils.Logger.Error(err)
		return false, utils.Fail
	}
	if info.Id < 1 {
		return false, utils.ErrorNotFund
	}

	pwd := fmt.Sprintf("%s%s", info.Salt, password)
	pwd, _ = utils.EncryptPassword(pwd)
	info.Password = pwd
	info.UpdateTime = utils.Now()
	adminName, _ := c.Get(utils.USERNAME)
	info.AdminId = adminId.(int)
	info.AdminName = adminName.(string)
	num, editErr := model.Edit(w, info)
	if editErr != nil {
		utils.Logger.Error(err)
		return false, utils.Fail
	}
	return num > 0, nil
}

func (s AdminService) Info(id int) (dao.AdminInfoRet, error) {
	info := dao.AdminInfo{}
	model := models.AdminModel{}.Init()
	w := models.Params{Eq: map[string]string{"id": strconv.Itoa(id)}}
	err := model.GetOne(w, &info)
	var roles []int
	for _, role := range strings.Split(info.Roles, ",") {
		r, _ := strconv.Atoi(role)
		roles = append(roles, r)
	}
	retData := dao.AdminInfoRet{
		AdminInfo: info,
		Roles:     roles,
	}
	return retData, err
}

func (s AdminService) Items(validator validator.AdminItemsValidator) (dao.AdminItems, error) {

	model := models.AdminModel{}.Init()
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
	if len(validator.Name) > 0 {
		w.Like["name LIKE ? "] = "%%" + validator.Name + "%%"
	}

	if len(validator.Phone) > 0 {
		w.Like["phone LIKE ? "] = "%%" + validator.Phone + "%%"
	}
	if len(validator.Email) > 0 {
		w.Like["email LIKE ? "] = "%%" + validator.Email + "%%"
	}
	if validator.Status > 0 {
		w.Eq = map[string]string{"status": strconv.Itoa(int(validator.Status))}
	}

	var data []dao.AdminInfo
	var items dao.AdminItems
	var ids []dao.Count
	count, err := model.Count(w, &ids)
	if err != nil {
		utils.Logger.Error(err)
		return items, err
	}

	model.Page(w, &data)
	roleInfos := s.getRoleInfos(data)
	for i := 0; i < len(data); i++ {
		var roles []dao.RoleNameItems
		for _, role := range strings.Split(data[i].Roles, ",") {
			if _, ok := roleInfos[role]; ok {
				roles = append(roles, roleInfos[role])
			}
		}
		items.Items = append(items.Items, dao.AdminItemsRet{
			AdminInfo: data[i],
			RoleInfo:  roles,
		})
	}

	items.Count = count
	return items, nil
}

func (s AdminService) checkUserNameAndPhone(user, phone string) (dao.Admin, utils.Error) {
	w := models.Params{Or: []map[string]string{
		{"username": phone}, {"phone": phone},
		{"username": user}, {"phone": user}}}
	info := dao.Admin{}
	err := models.AdminModel{}.Init().GetOne(w, &info)
	if err != nil {
		utils.Logger.Error(err)
		return info, utils.Fail
	}

	return info, nil
}

func (s AdminService) getRoleInfos(data []dao.AdminInfo) map[string]dao.RoleNameItems {
	var ids []string
	for i := 0; i < len(data); i++ {
		ids = append(ids, strings.Split(data[i].Roles, ",")...)
	}

	w := models.Params{In: map[string][]string{"id in ?": ids}}
	var items []dao.RoleNameItems
	retData := make(map[string]dao.RoleNameItems)

	err := models.RoleModel{}.Init().Items(w, &items)
	if err != nil {
		return retData
	}

	w = models.Params{}
	var projectItems []dao.ProjectNameItems
	pErr := models.ProjectModel{}.Init().Items(w, &projectItems)
	projectIds := make(map[int]string)
	if pErr == nil {
		for i := 0; i < len(projectItems); i++ {
			projectIds[projectItems[i].Id] = projectItems[i].Name
		}
	}
	for i := 0; i < len(items); i++ {
		projectName := ""
		if _, ok := projectIds[items[i].ProjectId]; ok {
			projectName = projectIds[items[i].ProjectId] + ":"
		}
		retData[strconv.Itoa(items[i].Id)] = dao.RoleNameItems{
			Id:   items[i].Id,
			Name: projectName + items[i].Name,
		}
	}

	return retData
}
