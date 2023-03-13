package eztools

// ConvArray 切片类型转换
func ConvArray[SRC, DST any](srcArr []SRC, fn func(src SRC) DST) []DST {
	var dstArr = make([]DST, 0)
	for i := 0; i < len(srcArr); i++ {
		dstArr = append(dstArr, fn(srcArr[i]))
	}
	return dstArr
}

// Array 切片类型
func Array[T comparable](arr []T) *arrayT[T] {
	t := new(arrayT[T])
	t.object = make([]T, len(arr))
	copy(t.object, arr)
	return t
}

type arrayT[T comparable] struct {
	object []T
}

// Result sort、filter 等操作后的值
func (receiver *arrayT[T]) Result() []T {
	return receiver.object
}

// Sort 排序
func (receiver *arrayT[T]) Sort(fn func(a, b T) bool) *arrayT[T] {
	for i := 0; i < len(receiver.object); i++ {
		for j := i + 1; j < len(receiver.object); j++ {
			if fn(receiver.object[i], receiver.object[j]) {
				receiver.object[i], receiver.object[j] = receiver.object[j], receiver.object[i]
			}
		}
	}
	return receiver
}

// Filter 过滤
func (receiver *arrayT[T]) Filter(fn func(T) bool) *arrayT[T] {
	var newArr = make([]T, 0)
	for i := 0; i < len(receiver.object); i++ {
		if fn(receiver.object[i]) {
			newArr = append(newArr, receiver.object[i])
		}
	}
	receiver.object = newArr
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

// Contains 是否包含
func (receiver *arrayT[T]) Contains(t T) bool {
	for i := 0; i < len(receiver.object); i++ {
		if receiver.object[i] == t {
			return true
		}
	}
	return false
}

// IndexOf 获取索引
func (receiver *arrayT[T]) IndexOf(t T) int {
	for i := 0; i < len(receiver.object); i++ {
		if receiver.object[i] == t {
			return i
		}
	}
	return -1
}

// LastIndexOf 获取最后一个索引
func (receiver *arrayT[T]) LastIndexOf(t T) int {
	for i := len(receiver.object) - 1; i >= 0; i-- {
		if receiver.object[i] == t {
			return i
		}
	}
	return -1
}

// Remove 删除
func (receiver *arrayT[T]) Remove(t T) *arrayT[T] {
	for i := 0; i < len(receiver.object); i++ {
		if receiver.object[i] == t {
			receiver.object = append(receiver.object[:i], receiver.object[i+1:]...)
			i--
		}
	}
	return receiver
}

// RemoveAt 删除
func (receiver *arrayT[T]) RemoveAt(index int) *arrayT[T] {
	if index >= 0 && index < len(receiver.object) {
		receiver.object = append(receiver.object[:index], receiver.object[index+1:]...)
	}
	return receiver
}

// RemoveRange 删除
func (receiver *arrayT[T]) RemoveRange(start, end int) *arrayT[T] {
	if start >= 0 && start < len(receiver.object) && end >= 0 && end < len(receiver.object) {
		receiver.object = append(receiver.object[:start], receiver.object[end+1:]...)
	}
	return receiver
}

// Reverse 反转
func (receiver *arrayT[T]) Reverse() *arrayT[T] {
	for i := 0; i < len(receiver.object)/2; i++ {
		receiver.object[i], receiver.object[len(receiver.object)-1-i] = receiver.object[len(receiver.object)-1-i], receiver.object[i]
	}
	return receiver
}

// Insert 插入
func (receiver *arrayT[T]) Insert(index int, t T) *arrayT[T] {
	if index >= 0 && index < len(receiver.object) {
		receiver.object = append(receiver.object[:index], append([]T{t}, receiver.object[index:]...)...)
	}
	return receiver
}

// InsertRange 插入
func (receiver *arrayT[T]) InsertRange(index int, arr []T) *arrayT[T] {
	if index >= 0 && index < len(receiver.object) {
		receiver.object = append(receiver.object[:index], append(arr, receiver.object[index:]...)...)
	}
	return receiver
}

// ReplaceRange 替换
func (receiver *arrayT[T]) ReplaceRange(start, end int, arr []T) *arrayT[T] {
	if start >= 0 && start < len(receiver.object) && end >= 0 && end < len(receiver.object) {
		receiver.object = append(receiver.object[:start], append(arr, receiver.object[end+1:]...)...)
	}
	return receiver
}

// Merge 合并
func (receiver *arrayT[T]) Merge(arr []T) *arrayT[T] {
	receiver.object = append(receiver.object, arr...)
	return receiver
}

// Distinct 去重
func (receiver *arrayT[T]) Distinct() *arrayT[T] {
	var newArr = make([]T, 0)
	for i := 0; i < len(receiver.object); i++ {
		if !receiver.Contains(receiver.object[i]) {
			newArr = append(newArr, receiver.object[i])
		}
	}
	receiver.object = newArr
	return receiver
}

// Len 长度
func (receiver *arrayT[T]) Len() int {
	return len(receiver.object)
}

// IsEmpty 是否为空
func (receiver *arrayT[T]) IsEmpty() bool {
	return len(receiver.object) == 0
}
