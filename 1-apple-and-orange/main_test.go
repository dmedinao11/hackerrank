package main

import "testing"

func Test_countApplesAndOranges(t *testing.T) {
	type args struct {
		s       int32
		t       int32
		a       int32
		b       int32
		apples  []int32
		oranges []int32
	}
	tests := []struct {
		name  string
		args  args
		want  int32
		want1 int32
	}{
		{args: args{s: 7, t: 11, a: 5, b: 15, apples: []int32{-2, 2, 1}, oranges: []int32{5, -6}}, want: 1, want1: 1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := countApplesAndOranges(tt.args.s, tt.args.t, tt.args.a, tt.args.b, tt.args.apples, tt.args.oranges)
			if got != tt.want {
				t.Errorf("countApplesAndOranges() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("countApplesAndOranges() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
