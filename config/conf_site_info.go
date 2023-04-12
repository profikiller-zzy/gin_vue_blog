package config

type SiteInfo struct {
	CreatedAt   string `yaml:"created_at" json:"created_at"`
	BeiAn       string `yaml:"bei_an" json:"bei_an"`
	Title       string `yaml:"title" json:"title"`
	QQImage     string `yaml:"qq_image" json:"qq_image"`
	Version     string `yaml:"version" json:"version"`
	Email       string `yaml:"email" json:"email"`
	WechatImage string `yaml:"wechat_image" json:"wechat_image"`
	Name        string `yaml:"name" json:"name"`
	Job         string `yaml:"job" json:"job"`
	Slogan      string `yaml:"slogan" json:"slogan"`
	SloganEN    string `yaml:"slogan_EN" json:"slogan_EN"`
	Web         string `yaml:"web" json:"web"`
	BilibiliURL string `yaml:"bilibili_Url" json:"bilibili_Url"`
	GiteeURL    string `yaml:"gitee_Url" json:"gitee_Url"`
	GithubURL   string `yaml:"github_url" json:"github_url"`
}
