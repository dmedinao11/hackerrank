package main

import "testing"

func Test_divisibleSumPairs(t *testing.T) {
	type args struct {
		n  int32
		k  int32
		ar []int32
	}
	tests := []struct {
		name string
		args args
		want int32
	}{
		{args: args{n: 6, k: 5, ar: []int32{1, 2, 3, 4, 5, 6}}, want: 3},
		{args: args{n: 6, k: 3, ar: []int32{1, 3, 2, 6, 1, 2}}, want: 5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := divisibleSumPairs(tt.args.n, tt.args.k, tt.args.ar); got != tt.want {
				t.Errorf("divisibleSumPairs() = %v, want %v", got, tt.want)
			}
		})
	}
}
