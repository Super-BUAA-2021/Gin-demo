package initialize

import (
	"github.com/Super-BUAA-2021/Gin-demo/global"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	v1 "github.com/Super-BUAA-2021/Gin-demo/api/v1"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

func InitRouter(r *gin.Engine) {
	// 跨域配置
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AllowHeaders = append(config.AllowHeaders, "x-token")
	r.Use(cors.New(config))
	// 是否开启api文档页面
	if global.VP.GetBool("server.docs") {
		r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}
	// 禁用代理访问
	if err := r.SetTrustedProxies(nil); err != nil {
		global.LOG.Panic("初始化失败：禁止使用代理访问失败")
	}

	// 登录模块
	rawRouter := r.Group("/api/v1")
	{
		rawRouter.POST("/login", v1.Login)
		// rawRouter.POST("/users", v1.CreateUser)
	}

}
