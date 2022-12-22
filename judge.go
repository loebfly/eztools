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

func (receiver judge[T]) IfThen(condition bool, trueVal T) T {
	if condition {
		return trueVal
	}
	return any(nil).(T)
}

func (receiver judge[T]) IfThenElse(condition bool, trueVal T, falseVal T) T {
	if condition {
		return trueVal
	}
	return falseVal
}
