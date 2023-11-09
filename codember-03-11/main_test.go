package main

import "testing"

func Test_decodeMessage(t *testing.T) {
	type args struct {
		message string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			args: args{message: "cat dog dog car Cat doG sun"},
			want: "cat2dog3car1sun1",
		},
		{
			args: args{message: "keys house HOUSE house keys"},
			want: "keys2house3",
		},
		{
			args: args{message: "cup te a cup"},
			want: "cup2te1a1",
		},
		{
			args: args{message: "houses house housess"},
			want: "houses1house1housess1",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := decodeMessage(tt.args.message); got != tt.want {
				t.Errorf("decodeMessage() = %v, want %v", got, tt.want)
			}
		})
	}
}
