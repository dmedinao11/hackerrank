package main

import (
	"reflect"
	"testing"
)

func Test_prefixSumArray(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			args: args{arr: []int{10, 20, 10, 5, 15}},
			want: []int{10, 30, 40, 45, 60},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := prefixSumArray(tt.args.arr); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("prefixSumArray() = %v, want %v", got, tt.want)
			}
		})
	}
}
