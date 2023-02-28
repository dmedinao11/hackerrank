package main

import "testing"

func Test_beautifulDays(t *testing.T) {
	type args struct {
		i int32
		j int32
		k int32
	}
	tests := []struct {
		name string
		args args
		want int32
	}{
		{args: args{i: 20, j: 23, k: 6}, want: 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := beautifulDays(tt.args.i, tt.args.j, tt.args.k); got != tt.want {
				t.Errorf("beautifulDays() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_reverseNumber(t *testing.T) {
	type args struct {
		n int32
	}
	tests := []struct {
		name string
		args args
		want int32
	}{
		{args: args{51}, want: 15},
		{args: args{123456}, want: 654321},
		{args: args{50}, want: 5},
		{args: args{0}, want: 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverseNumber(tt.args.n); got != tt.want {
				t.Errorf("reverseNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}
