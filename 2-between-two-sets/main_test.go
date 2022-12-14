package main

import "testing"

func Test_getTotalX(t *testing.T) {
	type args struct {
		a []int32
		b []int32
	}
	tests := []struct {
		name string
		args args
		want int32
	}{
		{args: args{a: []int32{2, 6}, b: []int32{24, 36}}, want: 2},
		{args: args{a: []int32{2, 4}, b: []int32{16, 32, 96}}, want: 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getTotalX(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("getTotalX() = %v, want %v", got, tt.want)
			}
		})
	}
}
