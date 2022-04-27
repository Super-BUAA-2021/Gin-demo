package global

import (
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

var (
	DB *gorm.DB     // MySQL实例
	VP *viper.Viper // Viper实例
)
