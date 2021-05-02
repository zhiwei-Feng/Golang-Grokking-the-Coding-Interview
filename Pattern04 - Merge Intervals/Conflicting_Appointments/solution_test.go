package Conflicting_Appointments

import "testing"

func Test_canAttendAllAppointments(t *testing.T) {
	type args struct {
		intervals []Interval
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"case1", args{intervals: []Interval{{1, 4}, {2, 5}, {7, 9}}}, false},
		{"case2", args{intervals: []Interval{{6, 7}, {2, 4}, {8, 12}}}, true},
		{"case3", args{intervals: []Interval{{4, 5}, {2, 3}, {3, 6}}}, false},
		{"case4", args{intervals: []Interval{{4, 5}, {2, 3}, {5, 6}}}, true},
		{"case5", args{intervals: []Interval{}}, true},
		{"case6", args{intervals: []Interval{{1, 4}}}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := canAttendAllAppointments(tt.args.intervals); got != tt.want {
				t.Errorf("canAttendAllAppointments() = %v, want %v", got, tt.want)
			}
		})
	}
}
