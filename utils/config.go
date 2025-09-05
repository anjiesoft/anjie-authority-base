package utils

import (
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"os"
)

// GetConfig Get 获取配置数据
func GetConfig(key string) interface{} {
	con := viper.Get(key)
	return con
}

func GetConfigString(key string) string {
	con := viper.GetString(key)
	return con
}

func GetConfigInt(key string) int {
	con := viper.GetInt(key)
	return con
}

func GetConfigInt64(key string) int {
	con := viper.GetInt(key)
	return con
}

func GetConfigBool(key string) bool {
	con := viper.GetBool(key)
	return con
}

// 监听配置变化 - 热更新
func watch() {
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {})
}

// 加载配置viper
func init() {
	//设置文件名和文件后缀
	viper.SetConfigName("env")
	viper.SetConfigType("toml")
	//配置文件所在的文件夹
	work, _ := os.Getwd()
	viper.AddConfigPath(work + "/config")
	err := viper.ReadInConfig()
	if err != nil {
		Logger.Error("loading config error", err)
	}

	watch()
}
