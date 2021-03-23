package Challenge1_Minimum_Meeting_Room

import "testing"

func Test_findMinimumMeetingRooms(t *testing.T) {
	type args struct {
		meetings []Meeting
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"case1", args{[]Meeting{{1, 4}, {2, 5}, {7, 9}}}, 2},
		{"case2", args{[]Meeting{{6, 7}, {2, 4}, {8, 12}}}, 1},
		{"case3", args{[]Meeting{{1, 4}, {2, 3}, {3, 6}}}, 2},
		{"case4", args{[]Meeting{{4, 5}, {2, 3}, {2, 4}, {3, 5}}}, 2},
		{"case5", args{[]Meeting{}}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findMinimumMeetingRooms(tt.args.meetings); got != tt.want {
				t.Errorf("findMinimumMeetingRooms() = %v, want %v", got, tt.want)
			}
		})
	}
}
