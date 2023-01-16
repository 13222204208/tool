package tool

import "time"

func Today() string {
	startTimestamp := time.Now().Unix()
	//获得时间戳
	startTimeStr := time.Unix(startTimestamp, 0).Format("2006-01-02")
	return startTimeStr
}

func TodayTime() string {
	startTimestamp := time.Now()
	return startTimestamp.Format("2006-01-02 15:04:05")
}

//不带秒显示的时间
func TodayMinutes() string {
	startTimestamp := time.Now()
	return startTimestamp.Format("2006-01-02 15:04")
}

//日期转时间戳
func DateToTime(timeStr string) int64 {
	var LOC, _ = time.LoadLocation("Asia/Shanghai")
	//要转换成时间日期的格式模板（go诞生时间，模板必须是这个时间）
	temp := "2006-01-02"
	tim, _ := time.ParseInLocation(temp, timeStr, LOC)
	c := tim.Unix()
	return c
}

// GetBeforeTime 获取n天前的秒时间戳、日期时间戳
// _day为负则代表取前几天，为正则代表取后几天，0则为今天
func GetBeforeTime(_day int) (int64, string) {
	// 时区
	//timeZone, _ := time.LoadLocation(ServerInfo["timezone"])
	timeZone := time.FixedZone("CST", 8*3600) // 东八区

	// 前n天
	nowTime := time.Now().In(timeZone)
	beforeTime := nowTime.AddDate(0, 0, _day)

	// 时间转换格式
	beforeTimeS := beforeTime.Unix()                             // 秒时间戳
	beforeDate := time.Unix(beforeTimeS, 0).Format("2006-01-02") // 固定格式的日期时间戳

	return beforeTimeS, beforeDate
}

// GetBetweenDates 根据开始日期和结束日期计算出时间段内所有日期
// 参数为日期格式，如：2020-01-01
func GetBetweenDates(startDate, endDate string) []string {
	var d []string
	timeFormatTpl := "2006-01-02 15:04:05"
	if len(timeFormatTpl) != len(startDate) {
		timeFormatTpl = timeFormatTpl[0:len(startDate)]
	}
	date, err := time.Parse(timeFormatTpl, startDate)
	if err != nil {
		// 时间解析，异常
		return d
	}
	date2, err := time.Parse(timeFormatTpl, endDate)
	if err != nil {
		// 时间解析，异常
		return d
	}
	if date2.Before(date) {
		// 如果结束时间小于开始时间，异常
		return d
	}
	// 输出日期格式固定
	timeFormatTpl = "2006-01-02"
	date2Str := date2.Format(timeFormatTpl)
	d = append(d, date.Format(timeFormatTpl))
	for {
		date = date.AddDate(0, 0, 1)
		dateStr := date.Format(timeFormatTpl)
		d = append(d, dateStr)
		if dateStr == date2Str {
			break
		}
	}
	return d
}