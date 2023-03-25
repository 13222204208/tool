package tool

import (
	"fmt"
	"math/rand"
	"time"
)

// 六位数字验证码
func RandCode() string {
	rand.Seed(time.Now().UnixNano())

	// 生成第一位非零数字
	firstDigit := rand.Intn(9) + 1

	// 生成后面五位数字
	num := firstDigit*100000 + rand.Intn(90000) + 10000

	// 转换为字符串
	str := fmt.Sprintf("%d", num)
	return str
}
