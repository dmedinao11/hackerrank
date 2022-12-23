package main

import (
	"reflect"
	"testing"
)

func Test_climbingLeaderboard(t *testing.T) {
	type args struct {
		ranked []int32
		player []int32
	}
	tests := []struct {
		name string
		args args
		want []int32
	}{
		{args: args{ranked: []int32{997, 988, 981, 966, 957, 937, 933, 930, 929, 928, 927, 926, 922, 920, 916, 915, 903, 896, 887, 874, 872, 866, 863, 863, 860, 859, 858, 857, 857, 847, 847, 842, 830, 819, 815, 809, 803, 797, 796, 794, 794, 789, 785, 783, 778, 772, 765, 765, 764, 757, 755, 751, 744, 740, 737, 733, 730, 730, 724, 716, 710, 709, 691, 690, 684, 677, 676, 653, 652, 650, 625, 620, 619, 602, 587, 587, 585, 583, 571, 568, 568, 556, 552, 546, 541, 540, 538, 531, 530, 529, 527, 506, 504, 501, 498, 493, 493, 492, 489, 482, 475, 468, 457, 452, 445, 442, 441, 438, 435, 435, 433, 430, 429, 427, 422, 422, 414, 408, 404, 400, 396, 394, 387, 384, 380, 379, 374, 371, 369, 369, 369, 368, 366, 365, 363, 354, 351, 341, 337, 336, 328, 325, 318, 316, 314, 307, 306, 302, 287, 282, 281, 277, 276, 271, 246, 238, 236, 230, 229, 229, 228, 227, 220, 212, 199, 194, 179, 173, 171, 168, 150, 144, 136, 125, 125, 124, 122, 118, 98, 98, 95, 92, 88, 85, 70, 68, 61, 60, 59, 44, 43, 35, 32, 30, 28, 23, 20, 13, 12, 12}, player: []int32{997}}, want: []int32{1}},
		{args: args{ranked: []int32{100, 90, 90, 80}, player: []int32{70, 80, 105}}, want: []int32{4, 3, 1}},
		{args: args{ranked: []int32{100, 100, 50, 40, 40, 20, 10}, player: []int32{5, 25, 50, 120}}, want: []int32{6, 4, 2, 1}},
		{args: args{ranked: []int32{100, 90, 90, 80, 75, 60}, player: []int32{50, 65, 77, 90, 102}}, want: []int32{6, 5, 4, 2, 1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := climbingLeaderboard(tt.args.ranked, tt.args.player); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("climbingLeaderboard() = %v, want %v", got, tt.want)
			}
		})
	}
}
