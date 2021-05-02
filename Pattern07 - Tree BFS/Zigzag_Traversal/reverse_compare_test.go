package Zigzag_Traversal

import "testing"

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
