package Happy_Number

import "testing"

func Test_squareSum(t *testing.T) {
	type args struct {
		num int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"case1", args{5}, 25},
		{"case2", args{23}, 13},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := squareSum(tt.args.num); got != tt.want {
				t.Errorf("squareSum() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_findHappyNumber(t *testing.T) {
	type args struct {
		num int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"case1", args{23}, true},
		{"case2", args{12}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findHappyNumber(tt.args.num); got != tt.want {
				t.Errorf("findHappyNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}
