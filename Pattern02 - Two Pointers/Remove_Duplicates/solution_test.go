package Remove_Duplicates

import "testing"

func Test_remove(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"case1", args{[]int{2, 3, 3, 3, 6, 9, 9}}, 4},
		{"case2", args{[]int{2, 2, 2, 11}}, 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := remove(tt.args.arr); got != tt.want {
				t.Errorf("remove() = %v, want %v", got, tt.want)
			}
		})
	}
}
