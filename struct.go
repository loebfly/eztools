package eztools

import (
	"reflect"
)

// StructToQuery 结构体转换为query参数
func StructToQuery(obj any, tag string) map[string]string {
	var (
		result = make(map[string]string)
	)

	if obj == nil {
		return result
	}

	objValue := reflect.ValueOf(obj).Elem()

	for i := 0; i < objValue.NumField(); i++ {
		key := objValue.Type().Field(i).Tag.Get(tag)
		if key == "" {
			continue
		}
		result[key] = Any(objValue.Field(i).Interface()).ToString()
	}

	return result
}

// CopyStruct 结构体复制, 忽略空值，暂不支持结构体内部map复制（有需要可扩展）
func CopyStruct[DST any](src any) DST {
	var dst DST
	dstValue := reflect.ValueOf(&dst).Elem()
	srcValue := reflect.ValueOf(src).Elem()

	if src == nil {
		return dst
	}

	// Recursively copy the original.
	copyStructRecursive(srcValue.Addr().Interface(), dstValue.Addr().Interface())

	return dst
}

// CopyStructTo 复制到目标结构体对象，忽略空值，暂不支持结构体内部map复制（有需要可扩展）
func CopyStructTo(dst any, src any) {
	srcValue := reflect.ValueOf(src).Elem()
	dstValue := reflect.ValueOf(dst).Elem()

	if src == nil {
		return
	}

	// Recursively copy the original.
	copyStructRecursive(srcValue.Addr().Interface(), dstValue.Addr().Interface())

	return
}

func copyStructRecursive(src, dst interface{}) {
	dstValue := reflect.ValueOf(dst).Elem()
	srcValue := reflect.ValueOf(src).Elem()

	for i := 0; i < srcValue.NumField(); i++ {
		srcField := srcValue.Field(i)
		srcName := srcValue.Type().Field(i).Name
		dstFieldByName := dstValue.FieldByName(srcName)
		if !dstFieldByName.IsValid() {
			continue
		}

		switch srcField.Kind() {
		case reflect.Slice:
			// Make a new slice and copy each element.
			if srcField.IsNil() {
				continue
			}
			dstFieldByName.Set(reflect.MakeSlice(dstFieldByName.Type(), srcField.Len(), srcField.Cap()))
			for j := 0; j < srcField.Len(); j++ {
				srcInterface := srcField.Index(j).Addr().Interface()
				dstInterface := dstFieldByName.Index(j).Addr().Interface()
				copyStructRecursive(srcInterface, dstInterface)
			}
		case reflect.Struct:
			copyStructRecursive(srcField.Addr().Interface(), dstFieldByName.Addr().Interface())
		default:
			if dstFieldByName.CanSet() && !isBlank(srcField) && dstFieldByName.Type() == srcField.Type() {
				dstFieldByName.Set(srcField)
			}
		}
	}
}

func isBlank(value reflect.Value) bool {
	switch value.Kind() {
	case reflect.String:
		return value.Len() == 0
	case reflect.Bool:
		return !value.Bool()
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return value.Int() == 0
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return value.Uint() == 0
	case reflect.Float32, reflect.Float64:
		return value.Float() == 0
	case reflect.Interface, reflect.Ptr:
		return value.IsNil()
	}
	return reflect.DeepEqual(value.Interface(), reflect.Zero(value.Type()).Interface())
}
