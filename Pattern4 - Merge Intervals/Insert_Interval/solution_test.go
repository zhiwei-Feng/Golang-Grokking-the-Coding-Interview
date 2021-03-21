package Insert_Interval

import (
	"reflect"
	"testing"
)

func Test_insert(t *testing.T) {
	type args struct {
		intervals   []Interval
		newInterval Interval
	}
	tests := []struct {
		name string
		args args
		want []Interval
	}{
		{"case1", args{[]Interval{{1, 3}, {5, 7}, {8, 12}}, Interval{4, 6}}, []Interval{{1, 3}, {4, 7}, {8, 12}}},
		{"case2", args{[]Interval{{1, 3}, {5, 7}, {8, 12}}, Interval{4, 10}}, []Interval{{1, 3}, {4, 12}}},
		{"case3", args{[]Interval{{2, 3}, {5, 7}}, Interval{1, 4}}, []Interval{{1, 4}, {5, 7}}},
		{"case4", args{[]Interval{}, Interval{1, 4}}, []Interval{{1, 4}}},
		{"case5", args{[]Interval{{2, 3}, {5, 7}}, Interval{3, 4}}, []Interval{{2, 4}, {5, 7}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := insert(tt.args.intervals, tt.args.newInterval); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("insert() = %v, want %v", got, tt.want)
			}
		})
	}
}
