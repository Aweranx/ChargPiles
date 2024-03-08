package upload

import (
	"ChargPiles/config"
	util "ChargPiles/pkg/utils/log"
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"os"
	"strconv"
)

func ProductUploadToLocalStatic(file multipart.File, bossId uint, productName string) (filePath string, err error) {
	bId := strconv.Itoa(int(bossId))
	basePath := "." + config.Config.PhotoPath.ProductPath + "boss" + bId + "/"
	if !DirExistOrNot(basePath) {
		CreateDir(basePath)
	}
	productPath := fmt.Sprintf("%s%s.jpg", basePath, productName)
	content, err := io.ReadAll(file)
	if err != nil {
		util.LogrusObj.Error(err)
		return "", err
	}
	err = os.WriteFile(productPath, content, 0666)
	if err != nil {
		util.LogrusObj.Error(err)
		return "", err
	}
	return fmt.Sprintf("boss%s/%s.jpg", bId, productName), err
}

// AvatarUploadToLocalStatic 上传头像
func AvatarUploadToLocalStatic(file multipart.File, userId uint, userName string) (filePath string, err error) {
	bId := strconv.Itoa(int(userId))
	basePath := "." + config.Config.PhotoPath.AvatarPath + "user" + bId + "/"
	if !DirExistOrNot(basePath) {
		CreateDir(basePath)
	}
	avatarPath := fmt.Sprintf("%s%s.jpg", basePath, userName)
	content, err := io.ReadAll(file)
	if err != nil {
		util.LogrusObj.Error(err)
		return "", err
	}
	err = os.WriteFile(avatarPath, content, 0666)
	if err != nil {
		util.LogrusObj.Error(err)
		return "", err
	}
	return fmt.Sprintf("user%s/%s.jpg", bId, userName), err
}

// DirExistOrNot 判断文件是否存在
func DirExistOrNot(fileAddr string) bool {
	s, err := os.Stat(fileAddr)
	if err != nil {
		log.Println(err)
		return false
	}
	return s.IsDir()
}

// CreateDir 创建文件夹
func CreateDir(dirName string) bool {
	err := os.MkdirAll(dirName, 0755)
	if err != nil {
		log.Println(err)
		return false
	}
	return true
}
