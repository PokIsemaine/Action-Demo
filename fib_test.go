package Action_Demo

import (
	"testing"
)

func BenchmarkFib10(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var _ = Fib(40)
	}
}

func BenchmarkFib20(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var _ = Fib(20)
	}
}
