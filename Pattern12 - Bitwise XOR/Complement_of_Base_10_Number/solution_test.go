package complementofbase10number

import "testing"

func Test_bitwiseComplement(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"case1", args{5}, 2},
		{"case2", args{7}, 0},
		{"case3", args{10}, 5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := bitwiseComplement(tt.args.n); got != tt.want {
				t.Errorf("bitwiseComplement() = %v, want %v", got, tt.want)
			}
		})
	}
}
