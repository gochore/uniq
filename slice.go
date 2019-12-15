package uniq

import "math"

// Ints sorts and deduplicates a slice of ints in increasing order.
func Ints(s *[]int) {
	Slice(s, func(i, j int) bool {
		return (*s)[i] < (*s)[j]
	}, func(i, j int) bool {
		return (*s)[i] == (*s)[j]
	})
}

// Float64s sorts and deduplicates a slice of float64s in increasing order
// (not-a-number values are treated as less than other values).
func Float64s(s *[]float64) {
	Slice(s, func(i, j int) bool {
		return (*s)[i] < (*s)[j] || math.IsNaN((*s)[i]) && !math.IsNaN((*s)[j])
	}, func(i, j int) bool {
		return (*s)[i] == (*s)[j]
	})
}

// Strings sorts and deduplicates a slice of strings in increasing order.
func Strings(s *[]string) {
	Slice(s, func(i, j int) bool {
		return (*s)[i] < (*s)[j]
	}, func(i, j int) bool {
		return (*s)[i] == (*s)[j]
	})
}

// IntsAreSorted tests whether a slice of ints is sorted and deduplicated in increasing order.
func IntsAreSorted(a []int) bool {
	return IsSorted(IntSlice(a))
}

// Float64sAreSorted tests whether a slice of float64s is sorted and deduplicated in increasing order
// (not-a-number values are treated as less than other values).
func Float64sAreSorted(a []float64) bool {
	return IsSorted(Float64Slice(a))
}

// StringsAreSorted tests whether a slice of strings is sorted and deduplicated in increasing order.
func StringsAreSorted(a []string) bool {
	return IsSorted(StringSlice(a))
}
