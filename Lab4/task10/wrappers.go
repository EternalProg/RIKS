package task10

import (
	"lab4/task6"
	"lab4/task7"
)

type Report struct {
	IntValues        []int
	FloatValues      []float64
	DivisibleBy4And3 int
	IntAverage       float64
	FloatAverage     float64
}

func AnalyzeSlice(values []int) (int, float64) {
	divisibleCount := task6.CountDivisibleBy4And3(values)
	average := task6.Average(values)

	return divisibleCount, average
}

func AnalyzeGeneratedSlices(step float64) Report {
	intValues := task7.GenerateIntSequence()
	divisibleCount, intAverage := AnalyzeSlice(intValues)

	floatValues := task7.GenerateFloatSequence(step)
	floatAverage := task6.AverageFloat(floatValues)

	return Report{
		IntValues:        intValues,
		FloatValues:      floatValues,
		DivisibleBy4And3: divisibleCount,
		IntAverage:       intAverage,
		FloatAverage:     floatAverage,
	}
}
