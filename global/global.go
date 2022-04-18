package global

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB       // MySQL实例
	VP  *viper.Viper   // Viper实例
	LOG *logrus.Logger // Logrus实例
)
