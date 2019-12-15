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
	data.Slice(0, deduplicate(data.Len(), data.Swap, data.Equal))
}

func IsUniqed(data EqualSorter) bool {
	n := data.Len()
	for i := n - 1; i > 0; i-- {
		if data.Less(i, i-1) || data.Equal(i, i-1) {
			return false
		}
	}
	return true
}

func Slice(data interface{}, less func(i, j int) bool, equal func(i, j int) bool) {
	value := reflect.ValueOf(data).Elem()
	slice := value.Interface()
	sort.Slice(slice, less)
	swapper := reflect.Swapper(slice)
	value.Set(value.Slice(0, deduplicate(value.Len(), swapper, equal)))
}

func deduplicate(length int, swap func(i, j int), equal func(i, j int) bool) int {
	j := 0
	for i := 1; i < length; i++ {
		if equal(i, j) {
			continue
		}
		j++
		swap(i, j)
	}
	return j + 1
}
