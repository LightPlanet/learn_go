package main_test

import(
	"testing"
	"unitest"
)

func TestIsEven(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		value int
		want  string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := main.IsEven(tt.value)
			// TODO: update the condition below to compare got with tt.want.
			if true {
				t.Errorf("IsEven() = %v, want %v", got, tt.want)
			}
		})
	}
}
