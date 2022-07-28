package util

import (
	"fmt"
	"time"
)

func GetCalculateTime(currentTimer time.Time, d string) (time.Time, error) {
	duration, err := time.ParseDuration(d)
	if err != nil {
		return time.Time{}, err
	}
	return currentTimer.Add(duration), nil
}

const (
	CSTLayout = "2006-01-02 15:04:05"
	IntLayout = "20060102150405"
	YMDLayout = "20060102"
)

func GetNowTime() time.Time {
	return time.Now()
}

// MicrosecondsStr 将 time.Duration 类型（nano seconds 为单位）
// 输出为小数点后 3 位的 ms （microsecond 毫秒，千分之一秒）
func MicrosecondsStr(elapsed time.Duration) string {
	return fmt.Sprintf("%.3fms", float64(elapsed.Nanoseconds())/1e6)
}
