# YC

YC is the Y combinator in Go with generics

This package is written based on the content of "[The Y combinator in Go with generics](https://eli.thegreenplace.net/2022/the-y-combinator-in-go-with-generics/)".

Implemented the Y combinator and some adaptions (e.g. memoize, tracing) for it.

## Synopsis

```go
package main

var factorialTag = func(recurse yc.Func[int, int]) yc.Func[int, int] {
	return func(n int) int {
		if n == 0 {
			return 1
		}
		return n * recurse(n-1)
	}
}

func main() {
    fac := yc.Y(yc.Adapt(factorialTag, yc.Memo[int, int](), yc.Trace[int, int]()))
    got := fac(10)
    fmt.Println(got) // 3628800
}
```

## Run tests

```
$ go test -timeout 30s ./... github.com/Code-Hex/yc
```

## Run benchmark

```
$ go test -benchmem -bench . github.com/Code-Hex/yc
```
