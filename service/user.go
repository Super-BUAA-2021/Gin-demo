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

// GetUserByUsername 根据用户名查询某个用户
func GetUserByUsername(username string) (user database.User, notFound bool) {
	err := global.DB.Where("name = ?", username).First(&user).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		return user, true
	} else if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return user, true
	} else {
		return user, false
	}
}

// SaveUser 根据信息保存用户
func SaveUser(user *database.User) (err error) {
	err = global.DB.Save(user).Error
	return err
}

// DeleteUserByID 根据uid 删除用户
func DeleteUserByID(ID uint64) (err error) {
	if err = global.DB.Where("id = ?", ID).Delete(database.User{}).Error; err != nil {
		return err
	}
	return nil
}

//// Query 查询用户对某问题的所有评测结果
//func QueryUserResult(uid uint64, pid uint64) (results []database.Result) {
//	global.DB.Where("user_id = ? AND problem_id = ?", uid, pid).Find(&results)
//	return results
//}
