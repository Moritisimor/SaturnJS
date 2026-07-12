package timelib

import "time"

func Now() map[string]any {
	rn := time.Now()

	return map[string]any{
		"second": rn.Second(),
		"minute": rn.Minute(),
		"hour": rn.Hour(),
	}
}
