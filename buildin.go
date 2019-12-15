package uniq

type IntSlice []int

func (s IntSlice) Len() int            { return len(s) }
func (s IntSlice) Less(i, j int) bool  { return s[i] < s[j] }
func (s IntSlice) Swap(i, j int)       { s[i], s[j] = s[j], s[i] }
func (s IntSlice) Equal(i, j int) bool { return s[i] == s[j] }
func (s *IntSlice) Slice(i, j int)     { *s = (*s)[i:j] }

type Float64Slice []int

func (s Float64Slice) Len() int            { return len(s) }
func (s Float64Slice) Less(i, j int) bool  { return s[i] < s[j] }
func (s Float64Slice) Swap(i, j int)       { s[i], s[j] = s[j], s[i] }
func (s Float64Slice) Equal(i, j int) bool { return s[i] == s[j] }
func (s *Float64Slice) Slice(i, j int)     { *s = (*s)[i:j] }

type StringSlice []int

func (s StringSlice) Len() int            { return len(s) }
func (s StringSlice) Less(i, j int) bool  { return s[i] < s[j] }
func (s StringSlice) Swap(i, j int)       { s[i], s[j] = s[j], s[i] }
func (s StringSlice) Equal(i, j int) bool { return s[i] == s[j] }
func (s *StringSlice) Slice(i, j int)     { *s = (*s)[i:j] }
