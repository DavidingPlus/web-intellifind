package file

import (
	"backend/model/core"
	"fmt"
	"github.com/disintegration/imaging"
	"github.com/gin-gonic/gin"
	"mime/multipart"
	"os"
	"strconv"
	"time"
)

func SaveUploadAvatar(c *gin.Context, file *multipart.FileHeader) (string, error) {

	if flag := IsFileValid(file.Filename, []string{"jpg", "png"}); !flag {
		return "", fmt.Errorf("请上传jpg,png格式文件！")
	}

	var avatar string
	// 确保目录存在，不存在创建
	publicPath := "public"
	dirName := fmt.Sprintf("/uploads/avatars/%d/", c.GetUint("current_user_id"))
	os.MkdirAll(publicPath+dirName, 0755)

	// 保存文件
	fileName := "temp" + "_" + "avatar.jpg"
	// public/uploads/avatars/2021/12/22/1/nFDacgaWKpWWOmOt.png
	avatarPath := publicPath + dirName + fileName

	if err := c.SaveUploadedFile(file, avatarPath); err != nil {
		return avatar, err
	}

	// 裁切图片
	img, err := imaging.Open(avatarPath, imaging.AutoOrientation(true))
	if err != nil {
		return avatar, err
	}
	resizeAvatar := imaging.Thumbnail(img, 256, 256, imaging.Lanczos)
	resizeAvatarName := "avatar.jpg"
	resizeAvatarPath := publicPath + dirName + resizeAvatarName

	err = imaging.Save(resizeAvatar, resizeAvatarPath)
	if err != nil {
		return avatar, err
	}

	// 删除老文件
	err = os.Remove(avatarPath)
	if err != nil {
		return avatar, err
	}

	return dirName + resizeAvatarName, nil
}

func SaveUploadJsonFile(c *gin.Context, file *multipart.FileHeader) (savePath string, filename string, err error) {

	if flag := IsFileValid(file.Filename, []string{"json"}); !flag {
		return "", "", fmt.Errorf("请上传json格式文件！")
	}
	// 确保目录存在，不存在创建
	publicPath := "public"
	dirName := fmt.Sprintf("/uploads/jsonfile/%d/", c.GetUint("current_user_id"))
	os.MkdirAll(publicPath+dirName, 0755)

	// 保存文件
	uid := c.GetUint("current_user_id")
	fmt.Println(file.Filename)
	filename = GenerateFileName(uid, file.Filename)

	savePath = publicPath + dirName + filename

	if err := c.SaveUploadedFile(file, savePath); err != nil {
		return "", "", err
	}

	temp := core.UploadJsonFileRecord{
		UID:        uid,
		UploadTime: time.Now(),
		FileName:   filename,
		SavePath:   savePath,
	}
	temp.Create()

	return savePath, filename, nil
}

func GenerateFileName(uid uint, file string) string {
	uid_string := strconv.Itoa(int(uid))

	filename := uid_string + "_" + strconv.Itoa(int(time.Now().Unix())) + "_" + file
	fmt.Println(filename)
	return filename
}

// 文件格式是否合适 仅通过后缀判断
func IsFileValid(filename string, valid []string) bool {
	length := len(filename)

	for i := length - 1; i >= 0; i-- {
		if filename[i] == '.' {
			break
		}
		length--
	}

	ext := filename[length:]
	for _, v := range valid {
		if ext == v {
			return true
		}
	}

	return false
}
