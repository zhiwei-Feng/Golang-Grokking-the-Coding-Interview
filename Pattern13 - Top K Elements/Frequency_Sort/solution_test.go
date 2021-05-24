package frequencysort

import "testing"

func Test_frequencySort(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{"case1", args{"tree"}, []string{"eert", "eetr"}},
		{"case2", args{"cccaaa"}, []string{"cccaaa", "aaaccc"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := frequencySort(tt.args.s)
			for _, v := range tt.want {
				if v == got {
					return
				}
			}
			t.Errorf("frequencySort() = %v, want %v", got, tt.want)
		})
	}
}
