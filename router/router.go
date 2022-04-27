package router

import (
	v1 "github.com/Super-BUAA-2021/Gin-demo/api/v1"
	_ "github.com/Super-BUAA-2021/Gin-demo/docs"
	"github.com/Super-BUAA-2021/Gin-demo/middleware"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	//"github.com/swaggo/gin-swagger/swaggerFiles"
	"net/http"
)

func InitRouter(r *gin.Engine) {
	//跨域配置
	r.Use(middleware.Cors())
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	r.GET("/hello", HelloGin)
	// 禁用代理访问
	if err := r.SetTrustedProxies(nil); err != nil {
		//global.LOG.Panic("初始化失败：禁止使用代理访问失败")
	}

	// 登录模块
	rawRouter := r.Group("/api/v1")
	{
		rawRouter.POST("/login", v1.Login)
	}
	// 静态资源模块
	resourceRouter := rawRouter.Group("/resource")
	{
		resourceRouter.POST("/upload", v1.UploadFile)
	}
}

func HelloGin(c *gin.Context) {

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "HelloGin!",
	})
}
