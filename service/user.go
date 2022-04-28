package service

import (
	"errors"
	"github.com/Super-BUAA-2021/Gin-demo/global"
	"github.com/Super-BUAA-2021/Gin-demo/model/database"
	"gorm.io/gorm"
)

// Helper

// 数据库操作

// CreateUser 创建用户
func CreateUser(user *database.User) (err error) {
	if err = global.DB.Create(user).Error; err != nil {
		return err
	}
	return nil
}

// GetUserByID 根据用户 ID 查询某个用户
func GetUserByID(ID uint64) (user database.User, notFound bool) {
	err := global.DB.First(&user, ID).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		return user, true
	} else if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return user, true
	} else {
		return user, false
	}
}
