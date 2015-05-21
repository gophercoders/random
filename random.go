package random

import (
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func GetRandomNumberInRange(min, max int) int {
	var n int
	var r int
	if max < min {
		panic("Max < min")
	}
	r = (max - min) + 1
	n = rand.Intn(r)
	return n + min
}
