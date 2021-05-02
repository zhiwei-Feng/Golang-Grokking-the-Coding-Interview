package Reverse_Level_Order_Traversal

import "testing"

func Test_method2(t *testing.T) {
	tests := []struct {
		name string
	}{
		{"case1"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			method1()
		})
	}
}

func Test_methond1(t *testing.T) {
	tests := []struct {
		name string
	}{
		{"case1"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			method2()
		})
	}
}

func BenchmarkMethod1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		method1()
	}
}

func BenchmarkMethod2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		method2()
	}
}
