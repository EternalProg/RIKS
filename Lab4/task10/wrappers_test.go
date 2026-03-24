package task10

import (
	"math"
	"testing"
)

func almostEqual(got, want float64) bool {
	return math.Abs(got-want) < 1e-9
}

func TestAnalyzeSliceSuccess(t *testing.T) {
	count, average := AnalyzeSlice([]int{12, 24, 5})

	if count != 2 {
		t.Fatalf("AnalyzeSlice() count = %d, want 2", count)
	}

	const wantAverage = 41.0 / 3
	if !almostEqual(average, wantAverage) {
		t.Fatalf("AnalyzeSlice() average = %v, want %v", average, wantAverage)
	}
}

func TestAnalyzeSliceEmpty(t *testing.T) {
	count, average := AnalyzeSlice([]int{})
	if count != 0 {
		t.Fatalf("AnalyzeSlice() count = %d, want 0", count)
	}

	if average != 0 {
		t.Fatalf("AnalyzeSlice() average = %v, want 0", average)
	}
}

func TestAnalyzeGeneratedSlicesSuccess(t *testing.T) {
	report := AnalyzeGeneratedSlices(0.5)

	if len(report.IntValues) != 25 {
		t.Fatalf("len(IntValues) = %d, want 25", len(report.IntValues))
	}

	if report.IntValues[0] != 23 || report.IntValues[len(report.IntValues)-1] != 47 {
		t.Fatalf("IntValues boundaries are incorrect: %v", report.IntValues)
	}

	if report.DivisibleBy4And3 != 2 {
		t.Fatalf("DivisibleBy4And3 = %d, want 2", report.DivisibleBy4And3)
	}

	if !almostEqual(report.IntAverage, 35) {
		t.Fatalf("IntAverage = %v, want 35", report.IntAverage)
	}

	if !almostEqual(report.FloatAverage, 0.5) {
		t.Fatalf("FloatAverage = %v, want 0.5", report.FloatAverage)
	}
}

func TestAnalyzeGeneratedSlicesInvalidStep(t *testing.T) {
	report := AnalyzeGeneratedSlices(0)

	if report.FloatAverage != 0 {
		t.Fatalf("FloatAverage = %v, want 0", report.FloatAverage)
	}
}
