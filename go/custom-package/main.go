package main

import (
	"fmt"

	"github.com/Mohitp98/citadel-hands-on/go/custom-package/math"
)

func main() {
	xs := []float64{1, 2, 3, 4}
	// xs := []float64{}
	avg := math.Average(xs)

	fmt.Println(avg)
}
