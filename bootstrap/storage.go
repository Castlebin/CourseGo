package bootstrap

import (
	"CourseGo/global"
	"github.com/jassue/go-storage/local"
)

func InitializeStorage() {
	_, _ = local.Init(global.App.Config.Storage.Disks.Local)
	// 没有这些东东，先注释掉，不然初始化报错
	// _, _ = kodo.Init(global.App.Config.Storage.Disks.QiNiu)
	// _, _ = oss.Init(global.App.Config.Storage.Disks.AliOss)
}
