package Action_Demo

import (
	"testing"
)

func BenchmarkFib10(b *testing.B) {
	var u uint = 20
	fmt.Println(u)
	for i := 0; i < b.N; i++ {
		var _ = Fib(u)
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
