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

type RoutesService struct {
}

func (s RoutesService) Items(c *gin.Context, validator validator.LoginRoutesValidator) (dao.Routes, utils.Error) {
	var items = dao.Routes{}
	var rules []string
	var rErr utils.Error
	super, _ := c.Get(utils.SUPERID)
	adminId, _ := c.Get(utils.USERID)
	all := true
	if _, ok := super.(map[string]string)[strconv.Itoa(adminId.(int))]; !ok {
		rules, rErr = s.getIds(adminId.(int), validator.ProjectId)
		if rErr != nil {
			return items, rErr
		}
		if len(rules) == 0 {
			return items, nil
		}
		all = false
	}

	items, rErr = s.userAuthority(rules, validator.ProjectId, all)
	if rErr != nil {
		return items, rErr
	}

	return items, nil
}

func (s RoutesService) getIds(adminId, projectId int) ([]string, utils.Error) {
	ids, uErr := s.userRoles(adminId)
	if uErr != nil {
		return nil, uErr
	}

	if len(ids) < 1 {
		return nil, nil
	}

	rules, rErr := s.userRules(ids, projectId)
	if rErr != nil {
		return nil, rErr
	}

	if len(rules) < 1 {
		return nil, nil
	}

	return rules, nil
}

func (s RoutesService) userRoles(adminId int) (string, utils.Error) {
	info, err := LoginService{}.Info(adminId)
	if err != nil {
		return "", utils.ErrorNotFund
	}

	return info.Roles, nil
}

func (s RoutesService) userRules(idsString string, projectId int) ([]string, utils.Error) {
	idsArray := strings.Split(idsString, ",")
	var rules []string
	roles, rErr := RoleService{}.GetByIds(idsArray, projectId)
	if rErr != nil {
		return nil, utils.Fail
	}
	if len(roles) < 1 {
		return rules, nil
	}

	for _, role := range roles {
		if len(role.Rules) > 0 {
			tmp := strings.Split(role.Rules, ",")
			rules = append(rules, tmp...)
		}
	}
	var ruleArray []string
	keys := make(map[string]bool)
	for _, rule := range rules {
		if _, value := keys[rule]; !value {
			keys[rule] = true
			ruleArray = append(ruleArray, rule)
		}
	}

	return ruleArray, nil
}

func (s RoutesService) userAuthority(idsArray []string, projectId int, all bool) (dao.Routes, utils.Error) {
	var items = dao.Routes{}
	w := models.Params{
		Eq: map[string]string{"status": strconv.Itoa(int(dao.StatusOk)), "project_id": strconv.Itoa(projectId)},
	}
	if !all {
		w.In = map[string][]string{"id in ?": idsArray}
	}
	w.Order = "level ASC, parent_id ASC, sort ASC"
	var info []dao.Authority
	err := models.AuthorityModel{}.Init().Items(w, &info)
	if err != nil {
		return items, utils.Fail
	}

	items.Menu = s.getList(info, 0)
	for _, item := range info {
		items.Auth = append(items.Auth, dao.RoutesAuth{
			Id:             item.Id,
			Type:           item.Type,
			Identification: item.Identification,
		})
	}

	return items, nil

}

// 获取层级的下接数据 列表使用
func (s RoutesService) getList(data []dao.Authority, pid int) []dao.RoutesMenu {
	var dataArr []dao.RoutesMenu
	for _, v := range data {
		if v.ParentId == pid {
			// 这里可以理解为每次都从最原始的数据里面找出相对就的ID进行匹配，直到找不到就返回
			children := s.getList(data, v.Id)
			node := dao.RoutesMenu{
				Id:             v.Id,
				Name:           v.Name,
				Sort:           v.Sort,
				ParentId:       v.ParentId,
				Remarks:        v.Remarks,
				Level:          v.Level,
				Children:       children,
				Api:            v.Api,
				Path:           v.Path,
				ViewPath:       v.ViewPath,
				Identification: v.Identification,
				Type:           v.Type,
				IsShow:         v.IsShow,
				Icon:           v.Icon,
			}
			dataArr = append(dataArr, node)
		}
	}
	return dataArr
}

func (s RoutesService) GetAuthority(adminId, projectId int) map[string]int {
	rules, rErr := s.getIds(adminId, projectId)
	if rErr != nil {
		return nil
	}
	if len(rules) == 0 {
		return nil
	}

	w := models.Params{
		Eq: map[string]string{"status": strconv.Itoa(int(dao.StatusOk)), "project_id": strconv.Itoa(projectId)},
		In: map[string][]string{"id in ?": rules},
	}

	w.Order = "level ASC, parent_id ASC, sort ASC"
	var info []dao.Authority
	err := models.AuthorityModel{}.Init().Items(w, &info)
	if err != nil {
		return nil
	}

	var items = map[string]int{}
	for _, item := range info {
		items[item.Identification] = item.Id
	}

	return items
}
