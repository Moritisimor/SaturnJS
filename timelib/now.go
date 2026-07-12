package timelib

import "time"

func Now() map[string]any {
	rn := time.Now()

	return map[string]any{
		"year": rn.Year(),
		"month": int(rn.Month()),
		"day": rn.Day(),
	}
}
