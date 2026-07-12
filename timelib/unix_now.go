package timelib

import "time"

func UnixNow() func() int64 {
	return func() int64 {
		return time.Now().Unix()
	}
}

func UnixNowMicros() func() int64 {
	return func() int64 {
		return time.Now().UnixMicro()
	}
}

func UnixNowMillis() func() int64 {
	return func() int64 {
		return time.Now().UnixMilli()
	}
}
