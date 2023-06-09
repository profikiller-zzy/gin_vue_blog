package config

// Config 包含所有配置信息
type Config struct {
	Mysql      Mysql      `yaml:"mysql"`
	Logger     Logger     `yaml:"logger"`
	System     System     `yaml:"system"`
	SaveUpload SaveUpload `yaml:"save_upload"`
	SiteInfo   SiteInfo   `yaml:"site_info"`
	QQ         QQ         `yaml:"qq"`
	Email      Email      `yaml:"email"`
	QiNiu      QiNiu      `yaml:"qi_niu"`
	Jwt        Jwt        `yaml:"jwt"`
	Redis      Redis      `yaml:"redis"`
	ES         ES         `yaml:"es"`
}
