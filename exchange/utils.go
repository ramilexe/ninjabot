package exchange

import (
	"math"
)

const float64EqualityThreshold = 1e-9

func almostEqual(a, b float64) bool {
	return math.Abs(a-b) <= float64EqualityThreshold
}

func IsZero(a float64) bool {
	return almostEqual(a, 0)
}
