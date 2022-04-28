package v1

import (
	"github.com/Super-BUAA-2021/Gin-demo/model/database"
	"github.com/Super-BUAA-2021/Gin-demo/model/response"
	"github.com/Super-BUAA-2021/Gin-demo/service"
	"github.com/Super-BUAA-2021/Gin-demo/utils"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

// Register
// @Summary      注册
// @Description  注册新用户
// @Tags         登录模块
// @Accept       json
// @Produce      json
// @Param        data  body      response.RegisterQ  true  "用户名，密码"
// @Success      200   {object}  response.CommonA  "是否成功，返回信息，Token"
// @Router       /register [post]
func Register(c *gin.Context) {
	// 获取请求数据
	data := utils.BindJsonAndValid(c, &response.RegisterQ{}).(*response.RegisterQ)
	// 用户的用户名已经注册过的情况
	if _, notFound := service.GetUserByUsername(data.Name); !notFound {
		c.JSON(http.StatusOK, response.CommonA{Success: false, Message: "用户已存在"})
		return
	}
	// 将密码进行哈希处理
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(data.Password), 12)
	if err != nil {
		panic("CreateUser: hash password error")
	}
	// 成功创建用户
	if err := service.CreateUser(&database.User{Name: data.Name, Password: string(hashedPassword)}); err != nil {
		panic("CreateUser: create user error")
	}
	// 返回响应
	c.JSON(http.StatusOK, response.CommonA{Success: true, Message: "注册成功"})
}

// Login
// @Summary      用户登录
// @Description  根据用户邮箱和密码等生成token，并将token返回给用户
// @Tags         登录模块
// @Accept       json
// @Produce      json
// @Param        data  body      response.LoginQ  true  "用户名，密码"
// @Success      200   {string}  string           "{"success": true, "message": "登录成功", "data": "model.User的所有信息"}"
// @Router       /login [post]
func Login(c *gin.Context) {

	var data response.LoginQ
	if err := utils.ShouldBindAndValid(c, &data); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"success": false,
			"message": "参数信息不合法" + err.Error(),
		})
		return
	}
	//fmt.Println(c.Request.Body,c.Params,c.Request.Form)
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "登录成功",
		"data":    "username:" + data.Name + ",password:" + data.Password,
	})
}
