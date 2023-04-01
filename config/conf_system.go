package config

type System struct {
	Host string `yaml:"host"`
	Port int    `yaml:"port"`
	Env  string `yaml:"env"` // 开发模式debug，打印所有信息
}
