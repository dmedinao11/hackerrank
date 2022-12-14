package main

import "testing"

func Test_birthday(t *testing.T) {
	type args struct {
		s []int32
		d int32
		m int32
	}
	tests := []struct {
		name string
		args args
		want int32
	}{
		{args: args{s: []int32{1, 2, 1, 3, 2}, d: 3, m: 2}, want: 2},
		{args: args{s: []int32{1, 1, 1, 1, 1, 1}, d: 3, m: 2}, want: 0},
		{args: args{s: []int32{4}, d: 4, m: 1}, want: 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := birthday(tt.args.s, tt.args.d, tt.args.m); got != tt.want {
				t.Errorf("birthday() = %v, want %v", got, tt.want)
			}
		})
	}
}
