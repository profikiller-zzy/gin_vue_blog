package model

type LoginDataModel struct {
	MODEL
	UserID   uint      `json:"user_id"` // 登录用户的ID
	User     UserModel `gorm:"foreignKey:UserID" json:"-"`
	Ip       string    `gorm:"size:32" json:"ip"`        // 用户登录的IP
	NickName string    `gorm:"size:36" json:"nick_name"` // 昵称
	Token    string    `gorm:"size:256" json:"token"`    // 用户登录的token
	Device   string    `gorm:"size:256" json:"device"`   // 用户登录的设备
	Addr     string    `gorm:"size:64" json:"addr"`      // 用户登录的地址
}
