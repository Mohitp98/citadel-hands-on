package math

// Average: returns average of array
func Average(xs []float64) float64 {
	total := float64(0)
	if len(xs) == 0 {
		return 0
	}

	for _, v := range xs {
		total += v
	}

	return total / float64(len(xs))
}
