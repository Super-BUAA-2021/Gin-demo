package main

import (
	"github.com/Super-BUAA-2021/Gin-demo/global"
	"github.com/Super-BUAA-2021/Gin-demo/initialize"
	"github.com/Super-BUAA-2021/Gin-demo/router"
	"github.com/gin-gonic/gin"
)

// @title        Demo API
// @version      1.0
// @description  Demo API description

// @contact.name   BFlameSwift
// @contact.url    https://github.com/BFlameSwift
// @contact.email  bflaemswift@163.com

// @BasePath      /api/v1
// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html
func main() {
	// 初始化全局资源
	global.VP = initialize.InitViper()
	global.DB = initialize.InitMySQL()
	// 创建Router
	r := gin.Default()
	router.InitRouter(r)
	r.Run(":8081")
}
