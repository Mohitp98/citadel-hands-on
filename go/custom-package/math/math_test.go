package math

import "testing"

type testpair struct {
	values  []float64
	average float64
}

// struct to represent the inputs and outputs
var tests = []testpair{
	{[]float64{1, 2}, 1.5},
	{[]float64{1, 1, 1, 1, 1, 1}, 1},
	{[]float64{-1, 1}, 0},
	{[]float64{}, 0},
}

// Tests are identified by starting a function with the word Test
// and taking one argument of type *testing.T.
func TestAverage(t *testing.T) {
	// sample test case
	// {
	// 	var v float64 = Average([]float64{1, 2})
	// 	if v != 1.5 {
	// 		t.Error("Expected 1.5, got ", v)
	// 	}
	// }

	{
		for _, pair := range tests {
			v := Average(pair.values)
			if v != pair.average {
				t.Error(
					"for", pair.average,
					"expected", pair.average,
					"got", v,
				)
			}
		}
	}
}
