package uniq

import (
	"reflect"
	"sort"
)

type Interface interface {
	EqualSorter
	Slice(i, j int)
}

type EqualSorter interface {
	sort.Interface
	Equal(i, j int) bool
}

func Uniq(data Interface) {
	sort.Sort(data)
	data.Slice(0, uniq(data.Len, data.Swap, data.Equal))
}

func IsUniqed(data EqualSorter) bool {
	return isUniqed(data.Len, data.Less, data.Equal)
}

func Slice(data interface{}, less func(i, j int) bool, equal func(i, j int) bool) {
	value := reflect.ValueOf(data).Elem()
	slice := value.Interface()
	sort.Slice(slice, less)
	swapper := reflect.Swapper(slice)
	value.Set(value.Slice(0, uniq(value.Len, swapper, equal)))
}

func IsSliceUniqed(data interface{}, less func(i, j int) bool, equal func(i, j int) bool) bool {
	value := reflect.ValueOf(data).Elem()
	return isUniqed(value.Len, less, equal)
}

func uniq(len func() int, swap func(i, j int), equal func(i, j int) bool) int {
	j := 0
	for i := 1; i < len(); i++ {
		if equal(i, j) {
			continue
		}
		j++
		swap(i, j)
	}
	return j + 1
}

func isUniqed(len func() int, less func(i, j int) bool, equal func(i, j int) bool) bool {
	n := len()
	for i := n - 1; i > 0; i-- {
		if less(i, i-1) || equal(i, i-1) {
			return false
		}
	}
	return true
}
