package tool

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

func UserNum() string {

	hour := time.Now().Hour()
	minute := time.Now().Minute()
	second := time.Now().Second()
	startTimestamp := time.Now().Unix()
	//获得时间戳
	startTimeStr := time.Unix(startTimestamp, 0).Format("20060102") //把时间戳转换成时间,并格式化为年月日
	nowTime := Sup(hour, 2) + Sup(minute, 2) + Sup(second, 2)
	randNum := RandInt(100, 999)

	code := startTimeStr + nowTime + strconv.Itoa(randNum)
	code = code[2:]
	return code
}

func RandInt(min, max int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max-min) + min
}

//对长度不足n的数字前面补0
func Sup(i int, n int) string {
	m := fmt.Sprintf("%d", i)
	for len(m) < n {
		m = fmt.Sprintf("0%s", m)
	}
	return m
}

//生成6位数字
func SixNumber() string {
	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
	code := fmt.Sprintf("%06v", rnd.Int31n(1000000))
	return code
}
