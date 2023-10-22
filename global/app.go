package global

import (
	"CourseGo/config"
	"github.com/go-redis/redis/v8"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type Application struct {
	ConfigViper *viper.Viper
	Config      config.Configuration // 代表 config.yaml 的数据结构
	Log         *zap.Logger          // zap 日志
	DB          *gorm.DB             // gorm 数据库实例
	Redis       *redis.Client        // redis 实例
}

var App = new(Application) // 全局变量
