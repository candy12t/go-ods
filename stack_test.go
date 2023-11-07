package ods

import "testing"

func BenchmarkArrayStack(b *testing.B) {
	a := NewArrayStack()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		a.Push(10000)
	}
}

func BenchmarkFastArrayStack(b *testing.B) {
	a := NewFastArrayStack()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		a.Push(10000)
	}
}
