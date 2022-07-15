package utils

import (
	"math/rand"
	"time"
)

//RandInt generate a number between a and b
func RandInt(a int, b int) int {
	rand.Seed(time.Now().UnixNano())

	min := a
	max := b

	if a >= b {
		min = b
		max = a
	}

	return rand.Intn(max-min) + min
}
