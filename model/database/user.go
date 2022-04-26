package database

import "time"

// User 用户
type User struct {
	ID       uint64    `gorm:"primary_key; autoIncrement; not null;" json:"id"`
	Name     string    `gorm:"size:32; not null; unique;" json:"name"`
	Password string    `gorm:"size:128; not null;" json:"password"`
	RegTime  time.Time `gorm:"autoCreateTime" json:"regTime"`
}
