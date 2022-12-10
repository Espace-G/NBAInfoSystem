package utils

import (
	"strconv"
	"time"
)

func GetOnlyFileName(filename string) (newName string) {
	s := strconv.FormatInt(time.Now().UnixMicro(), 10)
	newName = s + filename
	return
}
