package uniq

// Uint8s sorts and deduplicates a slice of uint8 in increasing order.
func Uint8s(s []uint8) int {
	return Slice(s, func(i int, j int) bool {
		return s[i] < s[j]
	})
}

// Uint16s sorts and deduplicates a slice of uint16 in increasing order.
func Uint16s(s []uint16) int {
	return Slice(s, func(i int, j int) bool {
		return s[i] < s[j]
	})
}

// Uint32s sorts and deduplicates a slice of uint32 in increasing order.
func Uint32s(s []uint32) int {
	return Slice(s, func(i int, j int) bool {
		return s[i] < s[j]
	})
}

// Uint64s sorts and deduplicates a slice of uint64 in increasing order.
func Uint64s(s []uint64) int {
	return Slice(s, func(i int, j int) bool {
		return s[i] < s[j]
	})
}

// Int8s sorts and deduplicates a slice of int8 in increasing order.
func Int8s(s []int8) int {
	return Slice(s, func(i int, j int) bool {
		return s[i] < s[j]
	})
}

// Int16s sorts and deduplicates a slice of int16 in increasing order.
func Int16s(s []int16) int {
	return Slice(s, func(i int, j int) bool {
		return s[i] < s[j]
	})
}

// Int32s sorts and deduplicates a slice of int32 in increasing order.
func Int32s(s []int32) int {
	return Slice(s, func(i int, j int) bool {
		return s[i] < s[j]
	})
}

// Int64s sorts and deduplicates a slice of int64 in increasing order.
func Int64s(s []int64) int {
	return Slice(s, func(i int, j int) bool {
		return s[i] < s[j]
	})
}

// Float32s sorts and deduplicates a slice of float32 in increasing order.
func Float32s(s []float32) int {
	return Slice(s, func(i int, j int) bool {
		return s[i] < s[j]
	})
}

// Uints sorts and deduplicates a slice of uint in increasing order.
func Uints(s []uint) int {
	return Slice(s, func(i int, j int) bool {
		return s[i] < s[j]
	})
}

// Uintptrs sorts and deduplicates a slice of uintptr in increasing order.
func Uintptrs(s []uintptr) int {
	return Slice(s, func(i int, j int) bool {
		return s[i] < s[j]
	})
}

// Bytes sorts and deduplicates a slice of byte in increasing order.
func Bytes(s []byte) int {
	return Slice(s, func(i int, j int) bool {
		return s[i] < s[j]
	})
}

// Runes sorts and deduplicates a slice of rune in increasing order.
func Runes(s []rune) int {
	return Slice(s, func(i int, j int) bool {
		return s[i] < s[j]
	})
}


// Uint8s tests whether a slice of uint8 is sorted and deduplicated in increasing order.
func Uint8sAreSorted(s []uint8) bool {
	return SliceIsSorted(s, func(i int, j int) bool {
		return s[i] < s[j]
	})
}

// Uint16s tests whether a slice of uint16 is sorted and deduplicated in increasing order.
func Uint16sAreSorted(s []uint16) bool {
	return SliceIsSorted(s, func(i int, j int) bool {
		return s[i] < s[j]
	})
}

// Uint32s tests whether a slice of uint32 is sorted and deduplicated in increasing order.
func Uint32sAreSorted(s []uint32) bool {
	return SliceIsSorted(s, func(i int, j int) bool {
		return s[i] < s[j]
	})
}

// Uint64s tests whether a slice of uint64 is sorted and deduplicated in increasing order.
func Uint64sAreSorted(s []uint64) bool {
	return SliceIsSorted(s, func(i int, j int) bool {
		return s[i] < s[j]
	})
}

// Int8s tests whether a slice of int8 is sorted and deduplicated in increasing order.
func Int8sAreSorted(s []int8) bool {
	return SliceIsSorted(s, func(i int, j int) bool {
		return s[i] < s[j]
	})
}

// Int16s tests whether a slice of int16 is sorted and deduplicated in increasing order.
func Int16sAreSorted(s []int16) bool {
	return SliceIsSorted(s, func(i int, j int) bool {
		return s[i] < s[j]
	})
}

// Int32s tests whether a slice of int32 is sorted and deduplicated in increasing order.
func Int32sAreSorted(s []int32) bool {
	return SliceIsSorted(s, func(i int, j int) bool {
		return s[i] < s[j]
	})
}

// Int64s tests whether a slice of int64 is sorted and deduplicated in increasing order.
func Int64sAreSorted(s []int64) bool {
	return SliceIsSorted(s, func(i int, j int) bool {
		return s[i] < s[j]
	})
}

// Float32s tests whether a slice of float32 is sorted and deduplicated in increasing order.
func Float32sAreSorted(s []float32) bool {
	return SliceIsSorted(s, func(i int, j int) bool {
		return s[i] < s[j]
	})
}

// Uints tests whether a slice of uint is sorted and deduplicated in increasing order.
func UintsAreSorted(s []uint) bool {
	return SliceIsSorted(s, func(i int, j int) bool {
		return s[i] < s[j]
	})
}

// Uintptrs tests whether a slice of uintptr is sorted and deduplicated in increasing order.
func UintptrsAreSorted(s []uintptr) bool {
	return SliceIsSorted(s, func(i int, j int) bool {
		return s[i] < s[j]
	})
}

// Bytes tests whether a slice of byte is sorted and deduplicated in increasing order.
func BytesAreSorted(s []byte) bool {
	return SliceIsSorted(s, func(i int, j int) bool {
		return s[i] < s[j]
	})
}

// Runes tests whether a slice of rune is sorted and deduplicated in increasing order.
func RunesAreSorted(s []rune) bool {
	return SliceIsSorted(s, func(i int, j int) bool {
		return s[i] < s[j]
	})
}
