package globalUtility

import (
	"time"
)

func GetCurrentTimestamp() string {
	currentTime := time.Now().Format("2006-01-02 15:04:05.000")
	return currentTime
}