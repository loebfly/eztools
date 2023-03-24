package eztools

// ConvMap 字典类型转换
func ConvMap[SrcKey comparable, SrcVal any, DstKey comparable, DstVal any](srcMap map[SrcKey]SrcVal, dstMap map[DstKey]DstVal, fn func(SrcKey, SrcVal) (DstKey, DstVal)) {
	for k, v := range srcMap {
		dstKey, dstVal := fn(k, v)
		dstMap[dstKey] = dstVal
	}
}

// Map 字典类型
func Map[K comparable, V any](m map[K]V) *mapT[K, V] {
	var t = new(mapT[K, V])
	t.object = make(map[K]V)
	for k, v := range m {
		t.object[k] = v
	}
	return t
}

type mapT[K comparable, V any] struct {
	object map[K]V
}

// Result sort、filter 等操作后的值
func (receiver *mapT[K, V]) Result() map[K]V {
	return receiver.object
}

// Keys 获取所有键
func (receiver *mapT[K, V]) Keys() []K {
	var keys = make([]K, 0)
	for k := range receiver.object {
		keys = append(keys, k)
	}
	return keys
}

// Values 获取所有值
func (receiver *mapT[K, V]) Values() []V {
	var values = make([]V, 0)
	for _, v := range receiver.object {
		values = append(values, v)
	}
	return values
}

// Find 找到第一个符合条件的索引和值, 如果没有找到则返回false
func (receiver *mapT[K, V]) Find(fn func(K, V) bool) (K, V, bool) {
	for k, v := range receiver.object {
		if fn(k, v) {
			return k, v, true
		}
	}
	var k K
	var v V
	return k, v, false
}

// Filter 过滤
func (receiver *mapT[K, V]) Filter(fn func(K, V) bool) *mapT[K, V] {
	var newMap = make(map[K]V)
	for k, v := range receiver.object {
		if fn(k, v) {
			newMap[k] = v
		}
	}
	receiver.object = newMap
	return receiver
}

// Contains 是否包含
func (receiver *mapT[K, V]) Contains(k K) bool {
	_, ok := receiver.object[k]
	return ok
}

// Remove 删除
func (receiver *mapT[K, V]) Remove(fn func(K, V) bool) *mapT[K, V] {
	var newMap = make(map[K]V)
	for k, v := range receiver.object {
		if !fn(k, v) {
			newMap[k] = v
		}
	}
	receiver.object = newMap
	return receiver
}

// Merge 合并
func (receiver *mapT[K, V]) Merge(m map[K]V) *mapT[K, V] {
	for k, v := range m {
		receiver.object[k] = v
	}
	return receiver
}
