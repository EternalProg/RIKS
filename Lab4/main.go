package main

import (
	"fmt"

	"lab4/task10"
)

var defaultStep = 0.25

func buildOutput(report task10.Report) string {
	return fmt.Sprintf(
		"Int sequence: %v\nFloat sequence: %v\nCount divisible by 4 and 3: %d\nAverage of int sequence: %.2f\nAverage of float sequence: %.2f",
		report.IntValues,
		report.FloatValues,
		report.DivisibleBy4And3,
		report.IntAverage,
		report.FloatAverage,
	)
}

func run(step float64) string {
	report := task10.AnalyzeGeneratedSlices(step)
	return buildOutput(report)
}

func main() {
	fmt.Println(run(defaultStep))
}
