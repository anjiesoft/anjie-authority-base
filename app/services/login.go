package services

import (
	"base-service/app/dao"
	"base-service/app/models"
	"base-service/utils"
	"strconv"
)

type LoginService struct {
}

// Login 登录
func (s LoginService) Login(userName string) (admin dao.Admin, err error) {
	w := models.Params{Or: []map[string]string{{"username": userName}, {"phone": userName}}}
	admin = dao.Admin{}
	err = models.AdminModel{}.Init().GetOne(w, &admin)
	if err != nil {
		utils.Logger.Error(err)
		return
	}
	return
}

func (s LoginService) Info(id int) (info dao.AdminInfo, err error) {
	w := models.Params{Or: []map[string]string{{"id": strconv.Itoa(id)}}}
	info = dao.AdminInfo{}
	err = models.AdminModel{}.Init().GetOne(w, &info)
	if err != nil {
		utils.Logger.Error(err)
		return
	}
	return
}

func (s LoginService) Fail(id, number int) {
	fail := dao.AdminFail{Id: id, FailNumber: number + 1}
	w := models.Params{Eq: map[string]string{"id": strconv.Itoa(id)}}
	_, err := models.AdminModel{}.Init().Edit(w, &fail)
	utils.Logger.Error(err)
	if err != nil {
		utils.Logger.Error(err)
		return
	}
	utils.Logger.Error(fail)
	return
}

func (s LoginService) Ok(id int) {
	fail := dao.AdminFail{Id: id, FailNumber: 0}
	w := models.Params{Eq: map[string]string{"id": strconv.Itoa(id)}}
	_, err := models.AdminModel{}.Init().Edit(w, &fail)
	admin := dao.AdminLastTime{
		Id:       id,
		LastTime: utils.Now(),
	}
	_, err = models.AdminModel{}.Init().Edit(w, &admin)
	if err != nil {
		return
	}
	return
}
