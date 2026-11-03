package main_test

import (
	main "homework_testing"
	"testing"
)

func TestDivide(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		a       float64
		b       float64
		want    float64
		wantErr bool
	}{
		{"x/x x != 0", 1.0, 1.0, 1.0, false},
		{"x/y x,y != 0", 6.0, 2.0, 3.0, false},
		{"x/0 x != 0", 1.0, 0.0, 0, true},
		{"0/0", 0.0, 0.0, 0, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, gotErr := main.Divide(tt.a, tt.b)
			if gotErr != nil {
				if !tt.wantErr {
					t.Errorf("Divide() failed: %v", gotErr)
				}
				return
			}
			if tt.wantErr {
				t.Fatalf("Divide() succeeded unexpectedly: %v", tt)
			}
			if got != tt.want {
				t.Errorf("Divide() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSumAll(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		vals []int
		want int
	}{
		{"empty", []int{}, 0},
		{"single", []int{42}, 42},
		{"multiple", []int{1, 2, 3, 4, 5}, (1 + 5) / 2 * 5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := main.SumAll(tt.vals)
			if got != tt.want {
				t.Errorf("SumAll() = %v, want %v", got, tt.want)
			}
		})
	}
}
