package random

import "testing"

var randomTestResults = []struct {
	min int
	max int
}{
	{0, 12},
	{-12, 0},
	{1, 12},
	{-12, -1},
	{-12, 1},
}

func TestGetRandomNumberInRange(t *testing.T) {
	var n int
	for _, td := range randomTestResults {
		// loop 100 million times - to make sure the random-ness averages out
		// this will take a long time to run.
		for i := 0; i < 100000000; i++ {
			n = GetRandomNumberInRange(td.min, td.max)
			if n < td.min {
				t.Fatalf("Random Number to small. Got %d but min is %d\n", n, td.min)
			}
			if n > td.max {
				t.Fatalf("Random Number to big. Got %d but max is %d\n", n, td.max)
			}
		}
	}
}
