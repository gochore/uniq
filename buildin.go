package uniq

import "math"

// IntSlice attaches the methods of Interface to []int, sorting in increasing order.
type IntSlice []int

func (s IntSlice) Len() int            { return len(s) }
func (s IntSlice) Less(i, j int) bool  { return s[i] < s[j] }
func (s IntSlice) Swap(i, j int)       { s[i], s[j] = s[j], s[i] }
func (s IntSlice) Equal(i, j int) bool { return s[i] == s[j] }
func (s *IntSlice) Slice(i, j int)     { *s = (*s)[i:j] }

// Float64Slice attaches the methods of Interface to []float64, sorting in increasing order.
type Float64Slice []float64

func (s Float64Slice) Len() int            { return len(s) }
func (s Float64Slice) Less(i, j int) bool  { return s[i] < s[j] || math.IsNaN(s[i]) && !math.IsNaN(s[j]) }
func (s Float64Slice) Swap(i, j int)       { s[i], s[j] = s[j], s[i] }
func (s Float64Slice) Equal(i, j int) bool { return s[i] == s[j] }
func (s *Float64Slice) Slice(i, j int)     { *s = (*s)[i:j] }

// StringSlice attaches the methods of Interface to []string, sorting in increasing order.
type StringSlice []string

func (s StringSlice) Len() int            { return len(s) }
func (s StringSlice) Less(i, j int) bool  { return s[i] < s[j] }
func (s StringSlice) Swap(i, j int)       { s[i], s[j] = s[j], s[i] }
func (s StringSlice) Equal(i, j int) bool { return s[i] == s[j] }
func (s *StringSlice) Slice(i, j int)     { *s = (*s)[i:j] }
