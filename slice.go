package uniq

import (
	"sort"
)

// Ints sorts and deduplicates a slice of ints in increasing order.
func Ints(s []int) int {
	return Sort(sort.IntSlice(s))
}

// Float64s sorts and deduplicates a slice of float64s in increasing order
// (not-a-number values are treated as less than other values).
func Float64s(s []float64) int {
	return Sort(sort.Float64Slice(s))
}

// Strings sorts and deduplicates a slice of strings in increasing order.
func Strings(s []string) int {
	return Sort(sort.StringSlice(s))
}

// IntsAreSorted tests whether a slice of ints is sorted and deduplicated in increasing order.
func IntsAreSorted(a []int) bool {
	return IsSorted(sort.IntSlice(a))
}

// Float64sAreSorted tests whether a slice of float64s is sorted and deduplicated in increasing order
// (not-a-number values are treated as less than other values).
func Float64sAreSorted(a []float64) bool {
	return IsSorted(sort.Float64Slice(a))
}

// StringsAreSorted tests whether a slice of strings is sorted and deduplicated in increasing order.
func StringsAreSorted(a []string) bool {
	return IsSorted(sort.StringSlice(a))
}
