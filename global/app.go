package global

import (
	"CourseGo/config"
	"github.com/go-redis/redis/v8"
	"github.com/jassue/go-storage/storage"
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

func (app *Application) Disk(disk ...string) storage.Storage {
	// 若未传参，默认使用配置文件驱动
	diskName := app.Config.Storage.Default
	if len(disk) > 0 {
		diskName = storage.DiskName(disk[0])
	}
	s, err := storage.Disk(diskName)
	if err != nil {
		panic(err)
	}
	return s
}
