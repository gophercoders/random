package random

import (
	"math/rand"
	"testing"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func TestGetRandomNumberInRange(t *testing.T) {
	var n int
	for i := 0; i < 100000000; i++ {
		n = GetRandomNumberInRange(1, 12)
		if n < 1 {
			t.Fatal("Got a unexpected zero")
		}
		if n > 12 {
			t.Fatal("Got an number larger than 12")
		}
	}
}
