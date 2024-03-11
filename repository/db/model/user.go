package model

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Password    string
	IsDeleted   int `gorm:"type:int;default:0"`
	Email       string
	PhoneNumber string `gorm:"unique"`
	Status      string
	Money       string
	Avatar      string `gorm:"size:1000"`
	NickName    string
}

func (u *User) CheckPassword(password string) bool {
	if u.Password == password {
		return true
	} else {
		return false
	}
}

// AvatarURL 头像地址
func (u *User) AvatarURL() string {
	return "avatar.jpg"
}
