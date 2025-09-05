package models

import (
	"base-service/app/dao"
	"base-service/utils"
	"encoding/json"
	"errors"
	"gorm.io/gorm"
)

var (
	db = *utils.DB
)

type Params struct {
	Eq    map[string]string   //等于
	Not   map[string]string   //不等于
	Or    []map[string]string //或者
	In    map[string][]string //包含
	Other map[string]string   //包含
	Like  map[string]string   //like
	Join  []string            //join
	Order string
	Page  int
	Size  int
}

type base struct {
	table string
}

func (m base) baseTable() string {
	return utils.GetConfigString("mysql.prefix")
}

// Edit 编辑
//
//	w := models.Params{
//			Eq:		map[string]string{"id": "371"},
//			Or:		[]map[string]string{{"id": "3", "nickname": "昵称"}, {"id": "4", "nickname": "昵称"}},
//			Not:	map[string]string{"id": "1", "nickname": "2"},
//			In:		map[string][]string{"id in ?": {"1", "3"}, "nickname in ?": {"33", "22"}},
//			Like:	map[string]string{"gender LIKE ? ": "%%2%%"},
//			Join:	[]string{"LEFT JOIN user_binding_host as b on b.user_openid = user.openid"},
//			Order:	"user.id desc",
//		}
//		user := dao.UserA{}
//		err := models.User{}.Init().Edit(w, user)
func (m base) Edit(w Params, data interface{}) (int64, error) {
	o := m.setWhere(w)
	update := make(map[string]interface{})

	jsonBytes, _ := json.Marshal(data)
	json.Unmarshal(jsonBytes, &update)

	ret := o.Updates(update)
	return ret.RowsAffected, m.error(ret.Error)
}

// Del 删除
//
//	w := models.Params{
//			Eq:    map[string]string{"id": "371"},
//			Or:    []map[string]string{{"id": "3", "nickname": "昵称"}, {"id": "4", "nickname": "昵称"}},
//			Not:   map[string]string{"id": "1", "nickname": "2"},
//			In:    map[string][]string{"id in ?": {"1", "3"}, "nickname in ?": {"33", "22"}},
//			Like:  map[string]string{"gender LIKE ? ": "%%2%%"},
//			Join:  []string{"LEFT JOIN user_binding_host as b on b.user_openid = user.openid"},
//			Order: "user.id desc",
//		}
//		user := dao.UserA{}
//		err := models.User{}.Init().Del(w, user)
func (m base) Del(w Params, data interface{}) (int64, error) {
	o := m.setWhere(w)
	ret := o.Delete(data)
	return ret.RowsAffected, m.error(ret.Error)
}

// Sql 原生语句获取数据
//
//	sql 	:= "SELECT * FROM lone_user WHERE id = @Id"
//	params	:= Params{Id: "1"}
//	user	:= dao.User{}
//	err		:= models.User{}.Sql(sql, params, &user)
func (m base) Sql(sql string, params interface{}, retData interface{}) error {
	err := db.Table(m.table).Raw(sql, params).Scan(&retData).Error

	return m.error(err)
}

// GetById 根据ID获取一条数据
//
//	user := dao.User{Id: 1}
//	err := models.User{}.Init().GetById(&user)
func (m base) GetById(retData interface{}) error {
	err := db.Table(m.table).First(&retData).Error

	return m.error(err)
}

// GetOne 获取一条数据
//
//	w := models.Params{
//		Eq:    map[string]string{"id": "1", "nickname": "昵称"},
//		Or:    []map[string]string{{"id": "3", "nickname": "昵称"}, {"id": "4", "nickname": "昵称"}},
//		Not:   map[string]string{"id": "1", "nickname": "2"},
//		In:    map[string][]string{"id in ?": {"1", "3"}, "nickname in ?": {"33", "22"}},
//		Like:  map[string]string{"gender LIKE ? ": "%%2%%"},
//		Join:  []string{"LEFT JOIN user_binding_host as b on b.user_openid = user.openid"},
//		Order: "user.id desc",
//	}
//	user := models.User{}
//	err := models.User{}.Init().Items(w, &user)
func (m base) GetOne(w Params, retData interface{}) error {
	o := m.setWhere(w)
	err := o.First(&retData).Error

	return m.error(err)
}

// Count 获取数量
//
//	w := models.Params{
//		Eq:    map[string]string{"id": "1"},
//		Or:    []map[string]string{{"id": "3", "nickname": "昵称"}, {"id": "4", "nickname": "昵称"}},
//		Not:   map[string]string{"id": "1", "nickname": "2"},
//		In:    map[string][]string{"id in ?": {"1", "3"}, "nickname in ?": {"33", "22"}},
//		Like:  map[string]string{"gender LIKE ? ": "%%2%%"},
//		Join:  []string{"LEFT JOIN lone_user_binding_host as b on b.user_openid = user.openid"},
//	}
//	user := models.User{}
//	total, err := models.User{}.Init().Count(w, &user)
func (m base) Count(w Params, retData interface{}) (int64, error) {
	o := m.setWhere(w)
	var count int64
	err := o.Find(retData).Count(&count).Error

	return count, m.error(err)
}

// Page 获取分页数据
//
//	w := models.Params{
//		Eq:    map[string]string{"id": "1"},
//		Or:    []map[string]string{{"id": "3", "nickname": "昵称"}, {"id": "4", "nickname": "昵称"}},
//		Not:   [string]string{"id": "1", "nickname": "2"},
//		In:    map[string][]string{"id in ?": {"1", "3"}, "nickname in ?": {"33", "22"}},
//		Like:  map[string]string{"gender LIKE ? ": "%%2%%"},
//		Join:  []string{"LEFT JOIN lone_user_binding_host as b on b.user_openid = user.openid"},
//		Order: "id desc",
//		Page:  1,
//		Size:  10,
//	}
//	user := []models.User{}
//	err := models.User{}.Init().Page(w, &user)
func (m base) Page(w Params, retData interface{}) error {
	o := m.setWhere(w)
	if w.Page < 1 {
		w.Page = 1
	}
	if w.Size < 1 {
		w.Size = 10
	}
	err := o.Offset((w.Page - 1) * w.Size).Limit(w.Size).Order(w.Order).Find(retData).Error
	return m.error(err)
}

// Items 获取列表数据
//
// user := []dao.User{}
//
//	w := models.Params{
//		Eq:    map[string]string{"id": "1", "nickname": "昵称"},
//		Or:    []map[string]string{{"id": "3", "nickname": "昵称"}, {"id": "4", "nickname": "昵称"}},
//		Not:   map[string]string{"id": "1", "nickname": "2"},
//		In:    map[string][]string{"id in ?": {"1", "3"}, "nickname in ?": {"33", "22"}},
//		Like:  map[string]string{"gender LIKE ? ": "%%2%%"},
//		Join:  []string{"LEFT JOIN user_binding_host as b on b.user_openid = user.openid"},
//		Order: "user.id desc",
//	}
//
// models.User{}.Init().Items(w, &user)
func (m base) Items(w Params, retData interface{}) error {
	o := m.setWhere(w)
	err := o.Order(w.Order).Find(retData).Error

	return m.error(err)
}

// 统一设置条件 Unified Settings
func (m base) setWhere(w Params) *gorm.DB {
	o := db.Debug().Table(m.table).Where(w.Eq)

	if w.Or != nil {
		for i := 0; i < len(w.Or); i++ {
			o.Or(w.Or[i])
		}
	}
	if w.In != nil {
		for k, v := range w.In {
			o.Where(k, v)
		}
	}
	if w.Like != nil {
		for k, v := range w.Like {
			o.Where(k, v)
		}
	}

	if w.Join != nil {
		for i := 0; i < len(w.Join); i++ {
			o.Joins(w.Join[i])
		}
	}

	if w.Other != nil {
		for k, v := range w.Other {
			o.Where(k, v)
		}
	}

	o.Not(w.Not)

	return o
}

func (m base) error(err error) error {
	//过滤空数据
	if errors.Is(err, gorm.ErrRecordNotFound) {
		err = nil
	}

	return err
}

// DemoCommit 事务 事例
//func (m Admin) DemoCommit() {
//	tx := db.Begin()
//	user := dao.User{Openid: "99999"}
//	uErr := tx.Table(m.TableName()).Create(&user).Error
//	utils.Logger.Info(uErr)
//	t2 := UserBindingHost{}.TableName()
//	userHost := dao.UserBindingHost{HostNickname: "aooana"}
//	uhErr := tx.Table(t2).Create(&userHost).Error
//	utils.Logger.Info(uhErr)
//	if uhErr != nil || uErr != nil {
//		tx.Rollback()
//	}
//
//	tx.Commit()
//}

// CreateMulti Create 添加数据
func (m base) CreateMulti(data *[]dao.Admin) error {
	ret := db.Debug().Table(m.table).Create(&data)
	return ret.Error
}
