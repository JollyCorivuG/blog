package model

import "gorm.io/gorm"

//用户信息
type User struct {
	gorm.Model
	Username string `json:"username"`
	Password string `json:"password"`
}

//博客信息
type Post struct {
	gorm.Model
	Title   string
	Content string `gorm:"type:text"`
	Tag     string
}
