package tool

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"unsafe"
)

func PostJson(url string, info map[string]string) (error, string) {

	bytesData, err := json.Marshal(info)
	if err != nil {
		fmt.Println(err.Error())
		return err, ""
	}

	reader := bytes.NewReader(bytesData)

	request, err := http.NewRequest("POST", url, reader)
	defer request.Body.Close() //程序在使用完回复后必须关闭回复的主体
	if err != nil {
		fmt.Println(err.Error())
		return err, ""
	}
	request.Header.Set("Content-Type", "application/json;charset=UTF-8")
	//必须设定该参数,POST参数才能正常提交，意思是以json串提交数据

	client := http.Client{}
	resp, err := client.Do(request) //Do 方法发送请求，返回 HTTP 回复
	if err != nil {
		fmt.Println("http回复的错误", err.Error())
		return err, ""
	}

	respBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("body体数据错误", err.Error())
		return err, ""
	}

	//byte数组直接转成string，优化内存
	str := (*string)(unsafe.Pointer(&respBytes))
	fmt.Println("返回的字符串数据", *str)
	return err, *str
	//fmt.Println(string(respBytes))
}

