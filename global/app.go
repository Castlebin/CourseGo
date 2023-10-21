package global

import (
	"CourseGo/config"
	"github.com/spf13/viper"
)

type Application struct {
	ConfigViper *viper.Viper
	Config      config.Configuration // 代表 config.yaml 的数据结构
}

var App = new(Application) // 全局变量
