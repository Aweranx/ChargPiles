package types

type UserServiceReq struct {
	NickName string `form:"nick_name" json:"nick_name"`
	//UserName    string `form:"user_name" json:"user_name"`
	Password    string `form:"password" json:"password"`
	Email       string `json:"email"`
	PhoneNumber string `form:"phone_number" json:"phone_number"`
	// Key      string `form:"key" json:"key"` // 前端进行判断
}

type UserRegisterReq struct {
	NickName string `form:"nick_name" json:"nick_name"`
	//UserName    string `form:"user_name" json:"user_name"`
	Password         string `form:"password" json:"password"`
	Email            string `json:"email"`
	PhoneNumber      string `form:"phone_number" json:"phone_number"`
	VerificationCode string `form:"verification_code" json:"verification_code"`
	// Key      string `form:"key" json:"key"` // 前端进行判断
}

type UserVerificationCodeReq struct {
	PhoneNumber string `form:"phone_number" json:"phone_number"`
}

type UserRegisterResp struct {
	VerificationCode string
}

type UserTokenData struct {
	User         interface{} `json:"user"`
	AccessToken  string      `json:"access_token"`
	RefreshToken string      `json:"refresh_token"`
}

type UserLoginReq struct {
	PhoneNumber string `form:"phone_number" json:"phone_number"`
	Password    string `form:"password" json:"password"`
}

type UserInfoUpdateReq struct {
	NickName string `form:"nick_name" json:"nick_name"`
}

type UserInfoShowReq struct {
}

type UserFollowingReq struct {
	Id uint `json:"id" form:"id"`
}

type UserUnFollowingReq struct {
	Id uint `json:"id" form:"id"`
}

type SendEmailServiceReq struct {
	Email    string `form:"email" json:"email"`
	Password string `form:"password" json:"password"`
	// OpertionType 1:绑定邮箱 2：解绑邮箱 3：改密码
	OperationType uint `form:"operation_type" json:"operation_type"`
}

type ValidEmailServiceReq struct {
	Token string `json:"token" form:"token"`
}

type UserInfoResp struct {
	ID          uint   `json:"id"`
	PhoneNumber string `json:"phone_number"`
	NickName    string `json:"nickname"`
	Type        int    `json:"type"`
	Email       string `json:"email"`
	Status      string `json:"status"`
	Avatar      string `json:"avatar"`
	CreateAt    int64  `json:"create_at"`
}
