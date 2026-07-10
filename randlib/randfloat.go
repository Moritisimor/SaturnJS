package randlib

import "math/rand/v2"

func RandFloat(low, high float64) float64 {
	return (rand.Float64() * (high - low) + low)
}
