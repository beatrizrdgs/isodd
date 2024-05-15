package isodd

import "testing"

func TestIsOdd(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"1", args{1}, true},
		{"2", args{2}, false},
		{"3", args{3}, true},
		{"4", args{4}, false},
		{"5", args{5}, true},
		{"6", args{6}, false},
		{"7", args{7}, true},
		{"8", args{8}, false},
		{"9", args{9}, true},
		{"0", args{10}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsOdd(tt.args.n); got != tt.want {
				t.Errorf("IsOdd() = %v, want %v", got, tt.want)
			}
		})
	}
}
