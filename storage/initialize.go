package storage

import (
	"github.com/Super-BUAA-2021/Gin-demo/config"
	"github.com/Super-BUAA-2021/Gin-demo/global"
)

// Setup 配置storage组件
func Setup() {

	// 读取yml配置文件
	global.VP = config.InitViper()

	// 初始化数据库
	global.DB = config.InitMySQL()

}
