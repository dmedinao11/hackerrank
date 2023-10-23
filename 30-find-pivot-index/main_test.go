package main

import "testing"

func Test_pivotIndex(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			args: args{nums: []int{0, 2}},
			want: -1,
		},
		{
			args: args{nums: []int{0, 1}},
			want: 1,
		},
		{
			args: args{nums: []int{1, 7, 3, 6, 5, 6}},
			// 					   0, 1, 2, 3, 4, 5
			// 1, 8, 11, 17, 22, 28
			want: 3,
		},
		{
			args: args{nums: []int{1, 2, 3}},
			want: -1,
		},
		{
			args: args{nums: []int{2, 1, -1}},
			want: 0,
		},
		{
			args: args{nums: []int{1}},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := pivotIndex(tt.args.nums); got != tt.want {
				t.Errorf("pivotIndex() = %v, want %v", got, tt.want)
			}
		})
	}
}
