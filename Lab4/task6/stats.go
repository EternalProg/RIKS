package task6

func CountDivisibleBy4And3(values []int) int {
	count := 0

	for _, value := range values {
		if value%4 == 0 && value%3 == 0 {
			count++
		}
	}

	return count
}

func Average(values []int) float64 {
	if len(values) == 0 {
		return 0
	}

	sum := 0
	for _, value := range values {
		sum += value
	}

	return float64(sum) / float64(len(values))
}

func AverageFloat(values []float64) float64 {
	if len(values) == 0 {
		return 0
	}

	var sum float64
	for _, value := range values {
		sum += value
	}

	return sum / float64(len(values))
}
