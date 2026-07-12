package timelib

import "time"

func SleepSecs(secs int) {
	time.Sleep(time.Duration(secs) * time.Second)
}

func SleepMillis(secs int) {
	time.Sleep(time.Duration(secs) * time.Millisecond)
}
