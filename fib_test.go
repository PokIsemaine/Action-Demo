package Action_Demo

import (
	"testing"
)

func BenchmarkFib10(b *testing.B) {
	s := make([]int, 10000)
	_ = s
	for i := 0; i < b.N; i++ {
		var _ = Fib(5)
	}
}

func BenchmarkFib20(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var _ = Fib(20)
	}
}

func BenchmarkFib30(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var _ = Fib(30)
	}
}
