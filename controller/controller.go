package controller

import (
	"blog/dao"
	"blog/model"
	"fmt"
	"html/template"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/russross/blackfriday/v2"
)

// 打开注册页面
func GoRegister(c *gin.Context) {
	c.HTML(200, "register.html", nil)
}

// 发送注册表单
func Register(c *gin.Context) {
	// 获取用户的id、密码
	username := c.PostForm("username")
	password := c.PostForm("password")

	user := model.User{
		Username: username,
		Password: password,
	}
	dao.Mgr.Register(&user)
	c.Redirect(301, "/")
}

// 打开登录页面
func GoLogin(c *gin.Context) {
	c.HTML(200, "login.html", nil)
}

// 发送登录表单
func Login(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	// 得到数据库中用户名为username的用户
	user := dao.Mgr.Login(username)
	fmt.Printf("user: %v\n", user)
	if user.Username == "" {
		c.HTML(200, "login.html", "用户名不存在！")
		fmt.Println("用户名不存在！")
	} else {
		if user.Password != password {
			c.HTML(200, "login.html", "密码错误！")
		} else {
			fmt.Println("登录成功！")
			c.Redirect(301, "/")
		}
	}
}

func Index(c *gin.Context) {
	c.HTML(200, "index.html", nil)
}

func ListUser(c *gin.Context) {
	c.HTML(200, "userlist.html", nil)
}

// 操作博客
func GoAddPost(c *gin.Context) {
	c.HTML(200, "addpost.html", nil)
}

func AddPost(c *gin.Context) {
	title := c.PostForm("title")
	content := c.PostForm("content")
	tag := c.PostForm("tag")

	post := model.Post{
		Title:   title,
		Content: content,
		Tag:     tag,
	}

	dao.Mgr.AddPost(&post)
	c.Redirect(302, "/postindex")
}

func GoPostIndex(c *gin.Context) {
	allpost := dao.Mgr.GetAllPost()
	c.HTML(200, "postindex.html", allpost)
}

func BlogDetail(c *gin.Context) {
	s := c.Query("pid")
	pid, _ := strconv.Atoi(s)
	post := dao.Mgr.GetPost(pid)
	content := blackfriday.Run([]byte(post.Content))
	c.HTML(200, "blogdetail.html", gin.H{
		"Title":   post.Title,
		"Content": template.HTML(content),
	})
}
