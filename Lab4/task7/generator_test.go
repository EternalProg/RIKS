package task7

import (
	"reflect"
	"testing"
)

func TestGenerateIntSequence(t *testing.T) {
	sequence := GenerateIntSequence()

	if len(sequence) != 25 {
		t.Fatalf("len(GenerateIntSequence()) = %d, want 25", len(sequence))
	}

	if sequence[0] != 23 {
		t.Fatalf("first value = %d, want 23", sequence[0])
	}

	if sequence[len(sequence)-1] != 47 {
		t.Fatalf("last value = %d, want 47", sequence[len(sequence)-1])
	}

	for i := 1; i < len(sequence); i++ {
		if sequence[i]-sequence[i-1] != 1 {
			t.Fatalf("sequence is not contiguous at index %d", i)
		}
	}
}

func TestGenerateFloatSequenceQuarterStep(t *testing.T) {
	got := GenerateFloatSequence(0.25)

	want := []float64{0, 0.25, 0.5, 0.75, 1}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("GenerateFloatSequence() = %v, want %v", got, want)
	}
}

func TestGenerateFloatSequenceNonDivisibleStep(t *testing.T) {
	got := GenerateFloatSequence(0.3)

	want := []float64{0, 0.3, 0.6, 0.9, 1}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("GenerateFloatSequence() = %v, want %v", got, want)
	}
}

func TestGenerateFloatSequenceStepOne(t *testing.T) {
	got := GenerateFloatSequence(1)

	want := []float64{0, 1}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("GenerateFloatSequence() = %v, want %v", got, want)
	}
}

func TestGenerateFloatSequenceInvalidStep(t *testing.T) {
	invalidSteps := []float64{0, -0.1, 1.1}

	for _, step := range invalidSteps {
		got := GenerateFloatSequence(step)
		want := []float64{0}
		if !reflect.DeepEqual(got, want) {
			t.Fatalf("GenerateFloatSequence(%v) = %v, want %v", step, got, want)
		}
	}
}
