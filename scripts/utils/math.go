package utils

import "math/rand"

// Between min-max int
func RandInt(min, max int) int {
	return min + rand.Intn(max+1-min)
}

// Between min-max float
func RandFloat(min, max float64) float64 {
	return min + rand.Float64()*(max-min)
}

func MinInt(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func MaxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}
