package dao

import (
	"ChargPiles/repository/db/model"
	"context"
	"gorm.io/gorm"
)

type UserDao struct {
	*gorm.DB
}

func NewUserDao(ctx context.Context) *UserDao {
	return &UserDao{NewDBClient(ctx)}
}

func (dao *UserDao) GetUserById(uId uint) (user *model.User, err error) {
	err = dao.DB.Model(&model.User{}).Where("id=?", uId).
		First(&user).Error
	return
}

func (dao *UserDao) UpdateUserById(uId uint, user *model.User) (err error) {
	return dao.DB.Model(&model.User{}).Where("id=?", uId).
		Updates(&user).Error
}

// ExistOrNotByUserName 根据username判断是否存在该名字
func (dao *UserDao) ExistOrNotByPhone(phoneNumber string) (user *model.User, exist bool, err error) {
	var count int64
	err = dao.DB.Model(&model.User{}).Where("phone_number = ?", phoneNumber).Count(&count).Error
	if count == 0 {
		return user, false, err
	}
	err = dao.DB.Model(&model.User{}).Where("phone_number = ?", phoneNumber).First(&user).Error
	if err != nil {
		return user, false, err
	}
	return user, true, nil
}

// CreateUser 创建用户
func (dao *UserDao) CreateUser(user *model.User) error {
	return dao.DB.Model(&model.User{}).Create(&user).Error
}
