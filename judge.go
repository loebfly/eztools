package eztools

func Judge[T any]() judge[T] {
	return judge[T]{}
}

type judge[T any] struct{}

func (receiver judge[T]) If(condition bool, trueVal T, falseVal T) T {
	if condition {
		return trueVal
	}
	return falseVal
}

func If[T any](condition bool, trueVal T, falseVal T) T {
	if condition {
		return trueVal
	}
	return falseVal
}

func IfCall(condition bool, trueFn func(), falseFn func()) {
	if condition {
		trueFn()
	} else {
		falseFn()
	}
}
