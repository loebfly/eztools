package eztools

import (
	"fmt"
	"testing"
)

func TestArray(t *testing.T) {
	var srcArr = []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	arr := Array(srcArr)

	var Result = arr.Filter(func(i int) bool {
		return i%2 == 0
	}).Sort(func(a, b int) bool {
		return a < b
	}).Result()

	fmt.Println(Result)

	var index, one = arr.Find(func(i int) bool {
		return i == 3
	})
	fmt.Println(index, one)
}
