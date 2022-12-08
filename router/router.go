package router

import (
	"blog/controller"

	"github.com/gin-gonic/gin"
)

func Start() {
	e := gin.Default()
	e.LoadHTMLGlob("templates/*")
	e.Static("/assets", "./assets")

	// 先进入首页
	e.GET("/", controller.Index)
	// 点击注册进入注册页面
	e.GET("/register", controller.GoRegister)
	// 提交注册表单后把用户信息加入数据库并回到首页
	e.POST("/register", controller.Register)
	// 点击登录进入登录页面
	e.GET("/login", controller.GoLogin)
	// 提交登录表单后回到首页
	e.POST("/login", controller.Login)
	// 点击博客进入博客页面
	e.GET("/postindex", controller.GoPostIndex)
	// 跳转到添加博客页面
	e.GET("/addpost", controller.GoAddPost)
	// 提交博客表单以添加博客
	e.POST("/addpost", controller.AddPost)
	// 跳转到博客详细页面
	e.GET("/blogdetail", controller.BlogDetail)

	e.Run(":8888")
}
