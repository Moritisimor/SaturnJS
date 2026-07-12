package timelib

import "github.com/dop251/goja"

func RegisterFuncs(js *goja.Runtime) error {
	return js.Set("time", map[string]any{
		"now":           Now,
		"unixNow":       UnixNow(),
		"unixNowMillis": UnixNowMillis(),
		"unixNowMicros": UnixNowMicros(),

		"date":               Date,
		"dateString":         DateString,
		"dateStringAmerican": DateStringAmerican,

		"sleepSecs":   SleepSecs,
		"sleepMillis": SleepMillis,
	})
}
