package rand

import (
	"math/rand"
	"time"
)

// Rand [min,max)
func Rand(min, max int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max-min) + min
}

// RandFromZero [0,max)
func RandFromZero(max int) int {
	return Rand(0, max)
}
