package randlib

import (
	"math/rand/v2"
)

func RandInt(low, high int) int {
	return rand.IntN(high-low+1) + low
}
