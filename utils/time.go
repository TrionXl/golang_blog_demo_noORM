package utils

import (
	"strconv"
	"time"
)

func GetNowTime() string {
	return strconv.FormatInt(time.Now().Unix(), 10)
}