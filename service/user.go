package service

import (
	"ChargPiles/config"
	"ChargPiles/consts"
	"ChargPiles/pkg/utils/ctl"
	"ChargPiles/pkg/utils/jwt"
	"ChargPiles/pkg/utils/log"
	"ChargPiles/pkg/utils/sms"
	"ChargPiles/pkg/utils/upload"
	"ChargPiles/repository/db/dao"
	"ChargPiles/repository/db/model"
	"ChargPiles/types"
	"context"
	"errors"
	"mime/multipart"
	"sync"
)

var UserSrvIns *UserSrv
var UserSrvOnce sync.Once

type UserSrv struct {
}

func GetUserSrv() *UserSrv {
	UserSrvOnce.Do(func() {
		UserSrvIns = &UserSrv{}
	})
	return UserSrvIns
}

func (s *UserSrv) UserRegister(ctx context.Context, req *types.UserRegisterReq) (resp interface{}, err error) {
	userDao := dao.NewUserDao(ctx)
	_, exist, err := userDao.ExistOrNotByPhone(req.PhoneNumber)
	if err != nil {
		//log.LogrusObj.Error(err)
		return
	}
	if exist {
		err = errors.New("用户已经存在了")
		return
	}
	correct, str := sms.CheckCode(req.PhoneNumber, req.VerificationCode)
	if correct == false {
		err = errors.New(str)
		return
	}
	user := &model.User{
		//UserName:    req.UserName,
		Password:    req.Password,
		Email:       req.Email,
		PhoneNumber: req.PhoneNumber,
		NickName:    req.NickName,
	}

	userDao.CreateUser(user)
	return
}

func (s *UserSrv) UserVerificationCode(ctx context.Context, req *types.UserVerificationCodeReq) {
	sms.SendVerificationCode(req.PhoneNumber)
}

func (s *UserSrv) UserLogin(ctx context.Context, req *types.UserServiceReq) (resp interface{}, err error) {

	userDao := dao.NewUserDao(ctx)
	user, exist, err := userDao.ExistOrNotByPhone(req.PhoneNumber)
	if !exist {
		return nil, errors.New("用户不存在")
	}

	if !user.CheckPassword(req.Password) {
		return nil, errors.New("账号/密码不正确")
	}

	accessToken, refreshToken, err := jwt.GenerateToken(user.ID, req.PhoneNumber)
	if err != nil {
		return nil, err
	}

	userResp := &types.UserInfoResp{
		ID:          user.ID,
		PhoneNumber: user.PhoneNumber,
		Email:       user.Email,
		Status:      user.Status,
		Avatar:      user.AvatarURL(),
		CreateAt:    user.CreatedAt.Unix(),
	}

	resp = &types.UserTokenData{
		User:         userResp,
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}
	return
}

func (s *UserSrv) UserInfoUpdate(ctx context.Context, req *types.UserInfoUpdateReq) (resp interface{}, err error) {
	u, _ := ctl.GetUserInfo(ctx)
	userDao := dao.NewUserDao(ctx)
	var user *model.User
	user, err = userDao.GetUserById(u.Id)
	if err != nil {
		log.LogrusObj.Error(err)
		return nil, err
	}
	if req.NickName != "" {
		user.NickName = req.NickName
	}

	err = userDao.UpdateUserById(u.Id, user)
	if err != nil {
		log.LogrusObj.Error(err)
		return nil, err
	}

	return
}

func (s *UserSrv) UserInfoShow(ctx context.Context, req *types.UserInfoShowReq) (resp interface{}, err error) {
	u, err := ctl.GetUserInfo(ctx)
	if err != nil {
		log.LogrusObj.Error(err)
		return
	}
	user, err := dao.NewUserDao(ctx).GetUserById(u.Id)
	if err != nil {
		log.LogrusObj.Error(err)
		return
	}
	resp = &types.UserInfoResp{
		ID:          user.ID,
		PhoneNumber: user.PhoneNumber,
		NickName:    user.NickName,
		Email:       user.Email,
		Status:      user.Status,
		Avatar:      user.AvatarURL(),
		CreateAt:    user.CreatedAt.Unix(),
	}

	return
}

func (s *UserSrv) UserAvatarUpload(ctx context.Context, file multipart.File, fileSize int64, req *types.UserServiceReq) (resp interface{}, err error) {
	u, err := ctl.GetUserInfo(ctx)
	if err != nil {
		log.LogrusObj.Error(err)
		return nil, err
	}
	uId := u.Id
	userDao := dao.NewUserDao(ctx)
	user, err := userDao.GetUserById(uId)
	if err != nil {
		log.LogrusObj.Error(err)
		return nil, err
	}

	var path string
	if config.Config.System.UploadModel == consts.UploadModelLocal { // 兼容两种存储方式
		path, err = upload.AvatarUploadToLocalStatic(file, uId, user.PhoneNumber)
	} else {
		path, err = upload.UploadToQiNiu(file, fileSize)
	}
	if err != nil {
		log.LogrusObj.Error(err)
		return nil, err
	}

	user.Avatar = path
	err = userDao.UpdateUserById(uId, user)
	if err != nil {
		log.LogrusObj.Error(err)
		return nil, err
	}

	return
}
