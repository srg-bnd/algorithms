// Algorithm: Moving average
package algo

// Method: Naive
func MovingAvrg(timeseries []int, k int) []float64 {
	var result []float64

	for i := (k - 1); i < len(timeseries); i++ {
		temp := 0.0

		for j := (i + 1 - k); j <= i; j++ {
			temp += float64(timeseries[j])
		}

		result = append(result, (temp / float64(k)))
	}

	return result
}

// Method: 2-pointer
func MovingAvrgWith2Pointers(timeseries []int, k int) []float64 {
	var result []float64

	sum := 0.0
	for i := 0; i < k; i++ {
		sum += float64(timeseries[i])
	}
	result = append(result, (sum / float64(k)))

	for i := k; i < len(timeseries); i++ {
		sum -= float64(timeseries[i-k])
		sum += float64(timeseries[i])
		result = append(result, (sum / float64(k)))
	}

	return result
}
