package utils

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/plugin/dbresolver"
)

var DB *gorm.DB

func init() {

	dbDSN := getConfig("")
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                       dbDSN,
		DefaultStringSize:         256,   // string 类型字段的默认长度
		DisableDatetimePrecision:  true,  // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		DontSupportRenameIndex:    true,  // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameColumn:   true,  // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: false, // 根据当前 MySQL 版本自动配置
	}), &gorm.Config{})

	readDSN := getConfig("_read")
	if readDSN != "" {
		//读写库不为空，就注册读写分离的配置
		err2 := db.Use(dbresolver.Register(dbresolver.Config{
			Sources:  []gorm.Dialector{mysql.Open(dbDSN)},   //写
			Replicas: []gorm.Dialector{mysql.Open(readDSN)}, //读
			Policy:   dbresolver.RandomPolicy{},
		}))
		if err2 != nil {
			Logger.Fatal("读写配置错误：%s", err2)
		}
	}

	if err != nil {
		Logger.Fatal("mysql clint error", err)
	} else {
		DB = db
	}
}

func getConfig(name string) string {
	user := GetConfigString("mysql" + name + ".user")
	if len(user) < 1 {
		return ""
	}
	password := GetConfigString("mysql" + name + ".password")
	host := GetConfigString("mysql" + name + ".host")
	port := GetConfigInt("mysql" + name + ".port")
	database := GetConfigString("mysql" + name + ".database")
	charset := GetConfigString("mysql" + name + ".charset")

	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=True&loc=Local", user, password, host, port, database, charset)
}
