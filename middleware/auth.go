package middleware

import (
	"github.com/Super-BUAA-2021/Gin-demo/service"
	"github.com/Super-BUAA-2021/Gin-demo/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func AuthRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("x-token")
		id, err := utils.ParseToken(token)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{"success": false, "message": "用户校验失败", "Code": 401})
			c.Abort()
			return
		}
		if user, notFound := service.GetUserByID(id); notFound {
			c.JSON(http.StatusOK, gin.H{"success": false, "message": "用户不存在", "Code": 404})
			c.Abort()
		} else {
			c.Set("user", user)
		}
	}
}
