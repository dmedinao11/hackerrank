package main

import "testing"

func Test_dayOfProgrammer(t *testing.T) {
	type args struct {
		year int32
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{args: args{year: 2016}, want: "12.09.2016"},
		{args: args{year: 2017}, want: "13.09.2017"},
		{args: args{year: 1918}, want: "26.09.1918"},
		{args: args{year: 1800}, want: "12.09.1800"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := dayOfProgrammer(tt.args.year); got != tt.want {
				t.Errorf("dayOfProgrammer() = %v, want %v", got, tt.want)
			}
		})
	}
}
