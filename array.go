package eztools

// ConvArray 切片类型转换
func ConvArray[SRC, DST any](srcArr []SRC, dstArr []DST, fn func(src SRC) DST) {
	for i := 0; i < len(srcArr); i++ {
		dstArr = append(dstArr, fn(srcArr[i]))
	}
}

// Array 切片类型
func Array[T any](arr []T) *arrayT[T] {
	t := new(arrayT[T])
	t.object = arr
	t.result = make([]T, len(arr))
	copy(t.result, arr)
	return t
}

type arrayT[T any] struct {
	object []T
	result []T
}

// OriVal 获取原始值
func (receiver *arrayT[T]) OriVal() []T {
	return receiver.object
}

// Result sort和filter后的值
func (receiver *arrayT[T]) Result() []T {
	return receiver.result
}

// Sort 排序
func (receiver *arrayT[T]) Sort(fn func(a, b T) bool) *arrayT[T] {
	for i := 0; i < len(receiver.result); i++ {
		for j := i + 1; j < len(receiver.result); j++ {
			if fn(receiver.result[i], receiver.result[j]) {
				receiver.result[i], receiver.result[j] = receiver.result[j], receiver.result[i]
			}
		}
	}
	return receiver
}

// Filter 过滤
func (receiver *arrayT[T]) Filter(fn func(T) bool) *arrayT[T] {
	var newArr = make([]T, 0)
	for i := 0; i < len(receiver.result); i++ {
		if fn(receiver.result[i]) {
			newArr = append(newArr, receiver.result[i])
		}
	}
	receiver.result = newArr
	return receiver
}

// Find 找到第一个符合条件的索引和值
func (receiver *arrayT[T]) Find(fn func(T) bool) (int, T) {
	for i := 0; i < len(receiver.object); i++ {
		if fn(receiver.object[i]) {
			return i, receiver.object[i]
		}
	}
	var t T
	return -1, t
}
