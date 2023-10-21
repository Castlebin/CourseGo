package config

type Configuration struct { // 代表 config.yaml 的数据结构
	App App `mapstructure:"app" json:"app" yaml:"app"` // 代表 config.yaml 的数据结构中的 app 属性
}
