package yc

import "log"

type Func[T, U any] func(T) U
type TagFunc[T, U any] func(Func[T, U]) Func[T, U]
type CombinatorFunc[T, U any] func(CombinatorFunc[T, U]) Func[T, U]

func Y[T, U any](f TagFunc[T, U]) Func[T, U] {
	g := func(self CombinatorFunc[T, U]) Func[T, U] {
		return f(func(t T) U {
			return self(self)(t)
		})
	}
	return g(g)
}

func Adapt[T, U any](f TagFunc[T, U], adapters ...TagFunc[T, U]) TagFunc[T, U] {
	return func(self Func[T, U]) Func[T, U] {
		for i := len(adapters) - 1; i >= 0; i-- {
			self = adapters[i](self)
		}
		return f(self)
	}
}

func Memo[T comparable, U any]() TagFunc[T, U] {
	memo := map[T]U{}
	return func(f Func[T, U]) Func[T, U] {
		return func(t T) U {
			result, ok := memo[t]
			if ok {
				return result
			}
			tmp := f(t)
			memo[t] = tmp
			return tmp
		}
	}
}

func Trace[T, U any]() TagFunc[T, U] {
	return func(f Func[T, U]) Func[T, U] {
		return func(t T) U {
			log.Printf("call before f(%v)", t)
			result := f(t)
			log.Printf("called after f(%v) = %v", t, result)
			return result
		}
	}
}
