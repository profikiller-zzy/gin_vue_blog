package model

type CategoryModel struct {
	MODEL
	Content  string `gorm:"size:16" json:"content"`
	Describe string `json:"describe"`
}
