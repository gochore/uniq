package uniq

import (
	"reflect"
	"sort"
)

// Interface can be sorted and deduplicated by the routines in this package.
type Interface interface {
	EqualSorter
	// Slice writes receiver by v[i:j].
	Slice(i, j int)
}

// EqualSorter can be reported if sorted and deduplicated by the routines in this package.
type EqualSorter interface {
	sort.Interface
	// Equal reports if the element with index i is equal with the element with index j.
	Equal(i, j int) bool
}

// Sort sorts and deduplicated data.
func Sort(data Interface) {
	rv := reflect.ValueOf(data)
	if rv.Kind() == reflect.Invalid {
		return
	}
	if rv.IsZero() {
		return
	}
	if rv.Kind() != reflect.Ptr {
		panic("uniq: Sort(non-pointer " + reflect.TypeOf(data).String() + ")")
	}
	sort.Sort(data)
	data.Slice(0, deduplicate(data.Len(), data.Swap, data.Equal))
}

// IsSorted reports if data is sorted and deduplicated.
func IsSorted(data EqualSorter) bool {
	rv := reflect.ValueOf(data)
	if rv.Kind() == reflect.Invalid {
		return true
	}
	if rv.IsZero() {
		return true
	}
	return isSortedAndDeduplicated(data.Len(), data.Less, data.Equal)
}

// Slice sorts and deduplicates the provided slice given the provided less function and equal function.
func Slice(data interface{}, less func(i, j int) bool, equal func(i, j int) bool) {
	rv := reflect.ValueOf(data)
	if rv.Kind() == reflect.Invalid {
		return
	}
	if rv.IsZero() {
		return
	}
	if rv.Kind() != reflect.Ptr {
		panic("uniq: Slice(non-slice-pointer " + reflect.TypeOf(data).String() + ")")
	}
	rve := rv.Elem()
	if rve.Kind() != reflect.Slice {
		panic("uniq: Slice(non-slice-pointer " + reflect.TypeOf(data).String() + ")")
	}

	slice := rve.Interface()
	sort.Slice(slice, less)
	swapper := reflect.Swapper(slice)
	rve.Set(rve.Slice(0, deduplicate(rve.Len(), swapper, equal)))
}

// IsSliceSorted reports if slice is sorted and deduplicated.
func IsSliceSorted(data interface{}, less func(i, j int) bool, equal func(i, j int) bool) bool {
	rv := reflect.ValueOf(data)
	if rv.Kind() == reflect.Invalid {
		return true
	}
	if rv.IsZero() {
		return true
	}
	if rv.Kind() == reflect.Ptr {
		rv = rv.Elem()
		if rv.IsZero() {
			return true
		}
	}
	if rv.Kind() != reflect.Slice {
		panic("uniq: IsSliceSorted(non-slice " + reflect.TypeOf(data).String() + ")")
	}
	return isSortedAndDeduplicated(rv.Len(), less, equal)
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

func isSortedAndDeduplicated(length int, less func(i, j int) bool, equal func(i, j int) bool) bool {
	for i := length - 1; i > 0; i-- {
		if less(i, i-1) || equal(i, i-1) {
			return false
		}
	}
	return true
}
