package helper


import (
	"time"
)

func GetTodayDate() string {
	return time.Now().Format("2006-01-02")
}

func GetUnixTime() int64 {
	return time.Now().Unix();
}

func YmdToTime(timeString string) (string) {
	formatTimeStr,err := time.Parse("2006-01-02",timeString)
	if err !=  nil {
		return ""
	}
	return formatTimeStr.Format("2006/01/02")
}

func YmdToUnix(timeString string) (int64) {
	formatTimeStr,err := time.Parse("2006-01-02",timeString)
	if err !=  nil {
		return 0
	}
	return formatTimeStr.Unix()
}

func TimeToYmd(timestamp int64,formatString string) (string) {
	date := time.Unix(timestamp,0).Format(formatString)
	return date
}