package config

type Configuration struct { // 代表 config.yaml 的数据结构
	App App `mapstructure:"app" json:"app" yaml:"app"` // 代表 config.yaml 的数据结构中的 app 属性

	Log Log `mapstructure:"log" json:"log" yaml:"log"` // 代表 config.yaml 的数据结构中的 log 属性，日志配置

	Database Database `mapstructure:"database" json:"database" yaml:"database"` // 代表 config.yaml 的数据结构中的 database 属性，数据库配置
}
