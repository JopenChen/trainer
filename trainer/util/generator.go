package util

import (
	"time"
)

func GenerateID() int64 {
	return time.Now().Unix()
}
