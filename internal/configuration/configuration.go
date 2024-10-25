package configuration

import (
	"fmt"
	"time"
)

func LoadConfig() bool {
	// Set the time zone to Kolkata (Asia/Kolkata)
	timeZone, timeLoadErr := time.LoadLocation("Asia/Kolkata")
	if timeLoadErr != nil {
		//error in loading timezone
		panic(fmt.Sprintf("Error in loading timezone : %v", timeLoadErr))
	} else {
		time.Local = timeZone
	}
	return true
}