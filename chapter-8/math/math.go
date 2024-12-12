package math

// Average finds the average of a series of nnumbers
func Average(xs []float64) float64 {
	var sum float64
	for _, x := range xs {
		sum += x
	}
	return sum / float64(len(xs))
}
