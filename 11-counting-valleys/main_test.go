package main

import "testing"

func Test_countingValleys(t *testing.T) {
	type args struct {
		steps int32
		path  string
	}
	tests := []struct {
		name string
		args args
		want int32
	}{
		{args: args{steps: 8, path: "UDDDUDUU"}, want: 1},
		{args: args{steps: 8, path: "DDUUUUDD"}, want: 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countingValleys(tt.args.steps, tt.args.path); got != tt.want {
				t.Errorf("countingValleys() = %v, want %v", got, tt.want)
			}
		})
	}
}
