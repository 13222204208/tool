package tool

import (
	"bytes"
	"encoding/base64"
	"errors"
	"fmt"
	"image"
	"image/jpeg"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
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

//base64图片转成本地图片
func Base64ToImag(b64Img, path string) (imgUrl string, err error) {
	// 去掉前缀
	base64Data := strings.Split(b64Img, ",")[1]

	// 解码 Base64 字符串
	data, err := base64.StdEncoding.DecodeString(base64Data)
	if err != nil {
		log.Fatal(err)
	}

	// 创建 bytes.Buffer 对象
	buffer := bytes.Buffer{}
	buffer.Write(data)

	// 解码图片
	img, _, err := image.Decode(&buffer)
	if err != nil {
		log.Fatal(err)
	}

	imgUrl = path + UserNum() + ".jpg"
	// 将图片写入文件
	file, err := os.Create(imgUrl)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// 将图片以 JPEG 格式写入文件
	err = jpeg.Encode(file, img, nil)
	if err != nil {
		log.Fatal(err)
	}
	return
}
