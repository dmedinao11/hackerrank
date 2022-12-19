package main

import "testing"

func Test_formingMagicSquare(t *testing.T) {
	type args struct {
		s [][]int32
	}
	tests := []struct {
		name string
		args args
		want int32
	}{
		{args: args{s: [][]int32{{5, 3, 4}, {1, 5, 8}, {6, 4, 2}}}, want: 7},
		{args: args{s: [][]int32{{4, 9, 2}, {3, 5, 7}, {8, 1, 5}}}, want: 1},
		{args: args{s: [][]int32{{4, 8, 2}, {4, 5, 7}, {6, 1, 6}}}, want: 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := formingMagicSquare(tt.args.s); got != tt.want {
				t.Errorf("formingMagicSquare() = %v, want %v", got, tt.want)
			}
		})
	}
}
