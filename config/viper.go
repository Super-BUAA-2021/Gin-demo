package config

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/viper"
)

func InitViper() *viper.Viper {
	// 获取可执行文件所在目录
	rootPath, err := os.Executable()
	if err != nil {
		panic("初始化失败：可执行程序路径获取失败")
	}
	rootPath = filepath.Dir(rootPath)
	// 获取配置文件、题目、教程、用户头像保存路径
	path := filepath.Join(rootPath, "demo.yml")
	fmt.Println(path)
	// 初始化viper，读取配置文件
	v := viper.New()
	v.SetConfigFile(path)
	v.SetConfigType("yaml")
	err = v.ReadInConfig()
	if err != nil {
		panic("初始化失败：读取配置文件失败")
	}
	// 设置常用路径
	v.Set("root_path", rootPath)
	return v
}
