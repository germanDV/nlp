// Package shaker throws a dice cup/shaker
package shaker

import (
	"math"
	"math/rand"
	"time"
)

// dice returns a random number between [1, 6]
func dice() int {
	rand.Seed(time.Now().UnixNano())
	min := 1
	max := 6
	return rand.Intn((max+1)-min) + min
}

// Shake rolls a given number of dice
func Shake(n int) int {
	res := 0
	for i := 0; i < n; i++ {
		res += dice() * int(math.Pow10(i))
	}
	return res
}
