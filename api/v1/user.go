package v1

import (
	"fmt"
	"github.com/Super-BUAA-2021/Gin-demo/model"
	"github.com/Super-BUAA-2021/Gin-demo/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

// Login
// @Summary      用户登录
// @Description  根据用户邮箱和密码等生成token，并将token返回给用户
// @Tags         登录模块
// @Accept       json
// @Produce      json
// @Param        data  body      model.LoginQ  true  "用户名，密码"
// @Success      200       {string}  string  "{"success": true, "message": "登录成功", "data": "model.User的所有信息"}"
// @Router       /api/v1/login [post]
func Login(c *gin.Context) {

	var data model.LoginQ
	if err := utils.ShouldBindAndValid(c, &data); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"success": false,
			"message": "参数信息不合法" + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "登录成功",
		"data":    "username:" + data.Username + ",password:" + data.Password,
	})
}

// ShowAccount
// @Summary      Show an account
// @Description  get string by ID
// @Tags         accounts
// @Accept       json
// @Produce      json
// @Param        id   query  int     true  "Account ID"
// @Success      200  {string}  string  "{"success": true, "message": "登录成功", "data": "model.User的所有信息"}"
// @Router       /api/v1/account [post]
func ShowAccount(c *gin.Context) {
	id := c.Request.URL.Query().Get("id")
	fmt.Println(id, "1")
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "登录成功",
		"data":    id,
	})
}
