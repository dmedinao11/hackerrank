package main

import "testing"

func Test_utopianTree(t *testing.T) {
	type args struct {
		n int32
	}
	tests := []struct {
		name string
		args args
		want int32
	}{
		{args: args{n: 0}, want: 1},
		{args: args{n: 1}, want: 2},
		{args: args{n: 4}, want: 7},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := utopianTree(tt.args.n); got != tt.want {
				t.Errorf("utopianTree() = %v, want %v", got, tt.want)
			}
		})
	}
}
