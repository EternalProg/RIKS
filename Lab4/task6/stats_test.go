package task6

import (
	"math"
	"testing"
)

func almostEqual(got, want float64) bool {
	return math.Abs(got-want) < 1e-9
}

func TestCountDivisibleBy4And3(t *testing.T) {
	tests := []struct {
		name   string
		values []int
		want   int
	}{
		{
			name:   "mixed values",
			values: []int{12, 24, 36, 8, 9, 48, -12, 0, 7},
			want:   6,
		},
		{
			name:   "no matches",
			values: []int{1, 2, 3, 5, 7, 11},
			want:   0,
		},
		{
			name:   "empty slice",
			values: []int{},
			want:   0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := CountDivisibleBy4And3(tt.values)
			if got != tt.want {
				t.Fatalf("CountDivisibleBy4And3() = %d, want %d", got, tt.want)
			}
		})
	}
}

func TestAverageInt(t *testing.T) {
	got := Average([]int{12, 24, 36, 48})

	const want = (12 + 24 + 36 + 48) / 4
	if !almostEqual(got, want) {
		t.Fatalf("Average() = %v, want %v", got, want)
	}
}

func TestAverageFloat(t *testing.T) {
	got := AverageFloat([]float64{0, 0.5, 1})

	const want = (0 + 0.5 + 1) / 3
	if !almostEqual(got, want) {
		t.Fatalf("AverageFloat() = %v, want %v", got, want)
	}
}

func TestAverageEmpty(t *testing.T) {
	got := Average([]int{})
	if got != 0 {
		t.Fatalf("Average() value = %v, want 0", got)
	}
}

func TestAverageFloatEmpty(t *testing.T) {
	got := AverageFloat([]float64{})
	if got != 0 {
		t.Fatalf("AverageFloat() value = %v, want 0", got)
	}
}
