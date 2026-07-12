package timelib

import "time"

func Date() map[string]any {
	rn := time.Now()
	year, month, day := rn.Date()

	return map[string]any{
		"year":  year,
		"month": int(month),
		"day":   day,
	}
}
