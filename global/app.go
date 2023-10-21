package global

import (
	"CourseGo/config"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

type Application struct {
	ConfigViper *viper.Viper
	Config      config.Configuration // 代表 config.yaml 的数据结构
	Log         *zap.Logger          // zap 日志
}

var App = new(Application) // 全局变量
