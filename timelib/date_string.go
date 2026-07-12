package timelib

import (
	"fmt"
	"time"
)

func DateString(sep string) string {
	if sep == "" {
		sep = "-"
	}

	rn := time.Now()
	year, m, day := rn.Date()
	month := m.String()

	return fmt.Sprintf(
		"%d%s%s%s%d",
		day, sep,
		month, sep,
		year,
	)
}

func DateStringAmerican(sep string) string {
	if sep == "" {
		sep = "-"
	}
	
	rn := time.Now()
	year, m, day := rn.Date()
	month := m.String()

	return fmt.Sprintf(
		"%s%s%d%s%d",
		month, sep,
		day, sep,
		year,
	)
}
