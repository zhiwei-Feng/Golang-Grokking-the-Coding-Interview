package Remove_Duplicates

import "testing"

func Test_removeKey(t *testing.T) {
	type args struct {
		arr       []int
		targetKey int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"case1", args{[]int{3, 2, 3, 6, 3, 10, 9, 3}, 3}, 4},
		{"case2", args{[]int{2, 11, 2, 2, 1}, 2}, 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := removeKey(tt.args.arr, tt.args.targetKey); got != tt.want {
				t.Errorf("removeKey() = %v, want %v", got, tt.want)
			}
		})
	}
}
