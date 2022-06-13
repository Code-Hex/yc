package yc_test

import (
	"testing"

	"github.com/Code-Hex/yc"
)

var factorialTag = func(recurse yc.Func[int, int]) yc.Func[int, int] {
	return func(n int) int {
		if n == 0 {
			return 1
		}
		return n * recurse(n-1)
	}
}

var fibTag = func(recurse yc.Func[int, int]) yc.Func[int, int] {
	return func(n int) int {
		if n <= 1 {
			return n
		}
		return recurse(n-1) + recurse(n-2)
	}
}

func BenchmarkFac(b *testing.B) {
	fac := yc.Y(factorialTag)
	for i := 0; i < b.N; i++ {
		_ = fac(10)
	}
}

func BenchmarkFacMemo(b *testing.B) {
	fac := yc.Y(yc.Adapt(factorialTag, yc.Memo[int, int]()))
	for i := 0; i < b.N; i++ {
		_ = fac(10)
	}
}

func BenchmarkFib(b *testing.B) {
	fib := yc.Y(fibTag)
	for i := 0; i < b.N; i++ {
		_ = fib(10)
	}
}

func BenchmarkFibMemo(b *testing.B) {
	fib := yc.Y(yc.Adapt(fibTag, yc.Memo[int, int]()))
	for i := 0; i < b.N; i++ {
		_ = fib(10)
	}
}

func TestAdapt(t *testing.T) {
	fac := yc.Y(yc.Adapt(factorialTag, yc.Memo[int, int](), yc.Trace[int, int]()))
	got := fac(10)
	want := 3628800
	if want != got {
		t.Fatalf("want %d, but got %d", want, got)
	}
}
