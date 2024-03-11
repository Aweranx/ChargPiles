package sms

import (
	"ChargPiles/config"
	"ChargPiles/consts"
	"ChargPiles/repository/cache"
	"encoding/json"
	"fmt"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/dysmsapi"
	"log"
	"math/rand"
	"time"
)

var client *dysmsapi.Client

func Init() {
	sms := config.Config.Sms
	client, _ = dysmsapi.NewClientWithAccessKey(sms.EndPoint, sms.AccessKeyId, sms.AccessKeySercet)
}

func SendVerificationCode(phoneNumber string) {
	_, err := cache.RedisClient.Get(phoneNumber).Result()
	if err == nil {
		return
	}
	request := dysmsapi.CreateSendSmsRequest() //创建请求
	request.Scheme = consts.Scheme             //请求协议
	request.PhoneNumbers = phoneNumber         //接收短信的手机号码
	request.SignName = consts.SignName         //短信签名名称
	request.TemplateCode = consts.TemplateCode //短信模板ID
	code := getVerificationCode(phoneNumber)
	par, _ := json.Marshal(map[string]interface{}{ //定义短信模板参数（具体需要几个参数根据自己短信模板格式）
		"code": code,
	})
	request.TemplateParam = string(par) //将短信模板参数传入短信模板

	_, err = client.SendSms(request) //调用阿里云API发送信息
	if err != nil {                  //处理错误
		log.Println(err.Error())
	}

}

func getVerificationCode(phoneNumber string) (code string) {
	code = randomCode()
	err := cache.RedisClient.Set(phoneNumber, code, time.Minute*5).Err()
	if err != nil {
		panic(err)
	}
	return code
}

func CheckCode(phoneNumber string, verificationCode string) (correct bool, err string) {
	if verificationCode == "" {
		return false, "验证码不能为空"
	}
	code := cache.RedisClient.Get(phoneNumber)
	if code == nil {
		return false, "验证码已过期"
	}
	if code.Val() == verificationCode {
		return true, ""
	} else {
		return false, "验证码错误"
	}
}

func randomCode() string {
	rand := rand.New(rand.NewSource(time.Now().UnixNano()))
	return fmt.Sprintf("%06d", rand.Intn(1000000))
}
