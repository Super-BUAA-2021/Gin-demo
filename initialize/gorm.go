package initialize

import (
	"fmt"
	"github.com/Super-BUAA-2021/Gin-demo/global"
	"github.com/Super-BUAA-2021/Gin-demo/model/database"
	"os"
	"path/filepath"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

func InitMySQL() *gorm.DB {
	// 打开MySQL日志文件
	path, err := os.Executable()
	if err != nil {
		panic("初始化失败：可执行程序路径获取失败")
	}
	path = filepath.Dir(path)
	path = filepath.Join(path, "phoenix-mysql.log")
	// 读取配置数据
	addr := global.VP.GetString("database.ip")
	port := global.VP.GetString("database.port")
	user := global.VP.GetString("database.user")
	password := global.VP.GetString("database.password")
	dbstr := global.VP.GetString("database.database")
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", user, password, addr, port, dbstr)
	// 连接MySQL数据库
	var db *gorm.DB
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	if err != nil {
		panic("初始化失败：连接MySQL数据库失败")
	}
	// 更新MySQL数据库内容
	db.Set("gorm:table_options", "ENGINE=InnoDB")
	err = db.AutoMigrate(
		&database.User{},
	)
	if err != nil {
		panic("初始化失败：更新MySQL数据库内容失败")
	}
	return db
}
