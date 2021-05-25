package kthlargestnumberinastream

import (
	"testing"
)

func TestSolution(t *testing.T) {
	model := Constructor(4, []int{3, 1, 5, 12, 2, 11})
	a1 := model.Add(6)
	if a1!=5{
		t.Errorf("Constructor() = %v, want %v", a1, 5)
	}
	a2 := model.Add(13)
	if a2!=6{
		t.Errorf("Constructor() = %v, want %v", a2, 6)
	}
	a3 := model.Add(4)
	if a3!=6{
		t.Errorf("Constructor() = %v, want %v", a3, 6)
	}
}
