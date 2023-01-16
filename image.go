package tool

import (
	"encoding/base64"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

//本地图片转换为Base64 带前缀
func PrefixImgBase64(path string) (baseImg string, err error) {
	str, _ := os.Getwd()
	path = str + "/" + path
	fmt.Println("当前路径", path)
	//获取本地文件
	file, err := os.Open(path)
	if err != nil {
		err = errors.New("获取本地图片失败")
		return
	}
	defer file.Close()
	imgByte, _ := ioutil.ReadAll(file)

	// 判断文件类型，生成一个前缀，拼接base64后可以直接粘贴到浏览器打开，不需要可以不用下面代码
	//取图片类型
	mimeType := http.DetectContentType(imgByte)
	switch mimeType {
	case "image/jpeg":
		baseImg = "data:image/jpeg;base64," + base64.StdEncoding.EncodeToString(imgByte)
	case "image/png":
		baseImg = "data:image/png;base64," + base64.StdEncoding.EncodeToString(imgByte)
	}
	return
}

//本地图片转换为Base64 不带前缀
func ImgBase64(path string) (baseImg string, err error) {
	str, _ := os.Getwd()
	path = str + "/" + path
	fmt.Println("当前路径", path)
	//获取本地文件
	file, err := os.Open(path)
	if err != nil {
		err = errors.New("获取本地图片失败")
		return
	}
	defer file.Close()
	imgByte, _ := ioutil.ReadAll(file)

	baseImg = base64.StdEncoding.EncodeToString(imgByte)
	return
}