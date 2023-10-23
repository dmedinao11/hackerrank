package main

import "testing"

func Test_longestOnes(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			args: args{
				nums: []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
				k:    10,
			},
			want: 10,
		},
		{
			args: args{
				nums: []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
				k:    11,
			},
			want: 11,
		},
		{
			args: args{
				nums: []int{0},
				k:    1,
			},
			want: 1,
		},
		{
			args: args{
				nums: []int{0},
				k:    0,
			},
			want: 0,
		},
		{
			args: args{
				nums: []int{1},
				k:    0,
			},
			want: 1,
		},
		{
			args: args{
				nums: []int{1, 1, 1, 0, 0, 0, 1, 1, 1, 1, 0},
				k:    4,
			},
			want: 11,
		},
		{
			args: args{
				nums: []int{1, 1, 1, 0, 0, 0, 1, 1, 1, 1, 0},
				//i = 4
				k: 1,
			},
			want: 5,
		},
		{
			args: args{
				nums: []int{1, 1, 1, 0, 0, 0, 1, 1, 1, 1, 0},
				k:    0,
			},
			want: 4,
		},
		{
			args: args{
				nums: []int{0, 0, 1, 1, 0, 0, 1, 1, 1, 0, 1, 1, 0, 0, 0, 1, 1, 1, 1},
				//i=9
				//j=4
				k: 3,
			},
			want: 10,
		},
		{
			args: args{
				nums: []int{1, 1, 1, 0, 0, 0, 1, 1, 1, 1, 0},
				k:    2,
			},
			want: 6,
		},
		{
			args: args{
				nums: []int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
				k:    0,
			},
			want: 11,
		},
		{
			args: args{
				nums: []int{1, 1, 1, 0, 0, 0, 1, 1, 1, 1, 0},
				k:    4,
			},
			want: 11,
		},
		{
			args: args{
				nums: []int{1, 1, 1, 0, 0, 0, 1, 1, 1, 1, 0, 1},
				k:    4,
			},
			want: 12,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestOnes(tt.args.nums, tt.args.k); got != tt.want {
				t.Errorf("longestOnes() = %v, want %v", got, tt.want)
			}
		})
	}
}
