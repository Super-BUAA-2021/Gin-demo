package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Login doc
// @description 用户登录
// @Tags user
// @Param username formData string true "用户名"
// @Param password formData string true "密码"
// @Success 200 {string} string "{"success": true, "message": "登录成功", "data": "model.User的所有信息"}"
// @Router /user/login [post]
func Login(c *gin.Context) {
	// var user model.User
	// var notFound bool

	// username := c.Request.FormValue("username")
	// password := c.Request.FormValue("password")

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "登录成功",
		"data":    "success",
	})

	return
}
