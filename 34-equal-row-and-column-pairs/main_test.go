package main

import "testing"

func Test_equalPairs(t *testing.T) {
	type args struct {
		grid [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			args: args{grid: [][]int{{11, 1}, {1, 11}}},
			want: 2,
		},
		{
			args: args{grid: [][]int{{3, 1, 2, 2}, {1, 4, 4, 5}, {2, 4, 2, 2}, {2, 4, 2, 2}}},
			want: 3,
		},
		{
			args: args{grid: [][]int{
				{3, 1, 2, 2},
				{1, 4, 4, 4},
				{2, 4, 2, 2},
				{2, 5, 2, 2},
			}},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := equalPairs(tt.args.grid); got != tt.want {
				t.Errorf("equalPairs() = %v, want %v", got, tt.want)
			}
		})
	}
}
