package isodd

import "testing"

func TestIsOdd(t *testing.T) {
	tests := []struct {
		name string
		args int
		want bool
	}{
		{"0", 0, false},
		{"1", 1, true},
		{"2", 2, false},
		{"3", 3, true},
		{"4", 4, false},
		{"5", 5, true},
		{"6", 6, false},
		{"7", 7, true},
		{"8", 8, false},
		{"9", 9, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsOdd(tt.args); got != tt.want {
				t.Errorf("IsOdd() = %v, want %v", got, tt.want)
			}
		})
	}
}
