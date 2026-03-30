package task7

// Напишіть невиконаний (службовий) модуль, що містить функції, які виконують такі операції: створення зрізу, що складається з послідовності цілих (від 23 до 47) і значень з плаваючою крапкою від 0 до 1, з кроком, що задається.

import "math"

func GenerateIntSequence() []int {
	sequence := make([]int, 0, 25)
	for value := 23; value <= 47; value++ {
		sequence = append(sequence, value)
	}

	return sequence
}

func GenerateFloatSequence(step float64) []float64 {
	if step <= 0 || step > 1 {
		return []float64{0}
	}

	sequence := []float64{0}
	for value := step; value < 1; value += step {
		sequence = append(sequence, round(value))
	}

	if math.Abs(sequence[len(sequence)-1]-1) > 1e-9 {
		sequence = append(sequence, 1)
	}

	return sequence
}

func round(value float64) float64 {
	const precision = 1e9
	return math.Round(value*precision) / precision
}
