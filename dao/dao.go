package dao

import (
	"blog/model"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Manager interface {
	// 用户注册和登录
	Register(user *model.User)
	Login(username string) model.User

	// 博客操作
	AddPost(post *model.Post)
	GetAllPost() []model.Post
	GetPost(pid int) model.Post
}

type manager struct {
	db *gorm.DB
}

// 数据库管理者
var Mgr Manager

func init() {
	dsn := "root:jhc@tcp(127.0.0.1:3306)/golang_db?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	Mgr = &manager{db: db}
	// 创建表
	db.AutoMigrate(&model.User{})
	db.AutoMigrate(&model.Post{})
}

// 添加用户
func (mgr *manager) Register(user *model.User) {
	mgr.db.Model(&model.User{}).Create(user)
}

// 登录时查询用户
func (mgr *manager) Login(username string) model.User {
	var user model.User
	mgr.db.Model(&model.User{}).Where("username = ?", username).First(&user)
	return user
}

// 增加博客
func (mgr *manager) AddPost(post *model.Post) {
	mgr.db.Model(&model.Post{}).Create(post)
}

// 查询所有博客
func (mgr *manager) GetAllPost() []model.Post {
	var allpost []model.Post
	mgr.db.Model(&model.Post{}).Find(&allpost)
	return allpost
}

// 查询某个博客
func (mgr *manager) GetPost(pid int) model.Post {
	var post model.Post
	mgr.db.Model(&model.Post{}).Where("ID = ?", pid).First(&post)
	return post
}
