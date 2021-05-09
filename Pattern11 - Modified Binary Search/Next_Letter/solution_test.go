package nextletter

import "testing"

func Test_nextGreatestLetter(t *testing.T) {
	type args struct {
		letters []byte
		target  byte
	}
	tests := []struct {
		name string
		args args
		want byte
	}{
		{"case1", args{[]byte{'a', 'c', 'f', 'h'}, 'f'}, 'h'},
		{"case2", args{[]byte{'a', 'c', 'f', 'h'}, 'b'}, 'c'},
		{"case3", args{[]byte{'a', 'c', 'f', 'h'}, 'm'}, 'a'},
		{"case3", args{[]byte{'a', 'c', 'f', 'h'}, 'h'}, 'a'},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := nextGreatestLetter(tt.args.letters, tt.args.target); got != tt.want {
				t.Errorf("nextGreatestLetter() = %v, want %v", got, tt.want)
			}
		})
	}
}
