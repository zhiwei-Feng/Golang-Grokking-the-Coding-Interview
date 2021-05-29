package challenge3

import (
	"testing"
)

func TestConstructor(t *testing.T) {

	fs := Constructor()
	fs.Push(5)
	fs.Push(7)
	fs.Push(5)
	fs.Push(7)
	fs.Push(4)
	fs.Push(5)
	x := fs.Pop()
	if x != 5 {
		t.Errorf("pop %v, want %v", x, 5)
	}
	x = fs.Pop()
	if x != 7 {
		t.Errorf("pop %v, want %v", x, 7)
	}
	x = fs.Pop()
	if x != 5 {
		t.Errorf("pop %v, want %v", x, 5)
	}
	x = fs.Pop()
	if x != 4 {
		t.Errorf("pop %v, want %v", x, 4)
	}
}
