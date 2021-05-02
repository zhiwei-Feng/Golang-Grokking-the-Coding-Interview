package Challenge3_Employee_Free_Time

import (
	"reflect"
	"testing"
)

func Test_findEmployeeFreeTime(t *testing.T) {
	type args struct {
		schedule [][]Interval
	}
	tests := []struct {
		name string
		args args
		want []Interval
	}{
		{"case1", args{[][]Interval{
			{{1, 3}, {5, 6}},
			{{2, 3}, {6, 8}},
		}}, []Interval{{3, 5}}},
		{"case2", args{[][]Interval{
			{{1, 3}, {9, 12}},
			{{2, 4}},
			{{6, 8}},
		}}, []Interval{{4, 6}, {8, 9}}},
		{"case3", args{[][]Interval{
			{{1, 3}},
			{{2, 4}},
			{{3, 5}, {7, 9}},
		}}, []Interval{{5, 7}}},
		{"case4", args{[][]Interval{}}, []Interval{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findEmployeeFreeTime(tt.args.schedule); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findEmployeeFreeTime() = %v, want %v", got, tt.want)
			}
		})
	}
}
