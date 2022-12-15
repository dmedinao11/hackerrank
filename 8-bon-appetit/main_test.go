package main

import "testing"

func Test_bonAppetit(t *testing.T) {
	type args struct {
		bill []int32
		k    int32
		b    int32
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{args: args{bill: []int32{3, 10, 2, 9}, k: 1, b: 12}, want: "5"},
		{args: args{bill: []int32{3, 10, 2, 9}, k: 1, b: 7}, want: "Bon Appetit"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := bonAppetit(tt.args.bill, tt.args.k, tt.args.b); got != tt.want {
				t.Errorf("bonAppetit() = %v, want %v", got, tt.want)
			}
		})
	}
}
