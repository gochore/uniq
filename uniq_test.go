package uniq_test

import (
	"testing"

	"github.com/gochore/uniq"
)

func TestSort(t *testing.T) {
	type args struct {
		data uniq.Interface
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "nil",
			args: args{
				data: nil,
			},
		},
		{
			name: "emtpy",
			args: args{
				data: &uniq.IntSlice{},
			},
		},
		{
			name: "IntSlice",
			args: args{
				data: &uniq.IntSlice{1, 9, 9, 4, 0, 9, 2, 6, 2, 0, 1, 4, 0, 7, 1, 3},
			},
		},
		{
			name: "Float64Slice",
			args: args{
				data: &uniq.Float64Slice{1, 9, 9, 4, 0, 9, 2, 6, 2, 0, 1, 4, 0, 7, 1, 3},
			},
		},
		{
			name: "StringSlice",
			args: args{
				data: &uniq.StringSlice{"1", "9", "9", "4", "0", "9", "2", "6", "2", "0", "1", "4", "0", "7", "1", "3"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			uniq.Sort(tt.args.data)
			if !uniq.IsSorted(tt.args.data) {
				t.Fail()
			}
		})
	}
}

func TestSlice(t *testing.T) {
	intSlice := uniq.IntSlice{1, 9, 9, 4, 0, 9, 2, 6, 2, 0, 1, 4, 0, 7, 1, 3}
	ints := []int{1, 9, 9, 4, 0, 9, 2, 6, 2, 0, 1, 4, 0, 7, 1, 3}

	type args struct {
		data  interface{}
		less  func(i, j int) bool
		equal func(i, j int) bool
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "nil",
			args: args{
				data:  nil,
				less:  nil,
				equal: nil,
			},
		},
		{
			name: "empty IntSlice",
			args: args{
				data:  &uniq.IntSlice{},
				less:  nil,
				equal: nil,
			},
		},
		{
			name: "IntSlice",
			args: args{
				data:  &intSlice,
				less:  intSlice.Less,
				equal: intSlice.Equal,
			},
		},
		{
			name: "[]int",
			args: args{
				data: &ints,
				less: func(i, j int) bool {
					return ints[i] < ints[j]
				},
				equal: func(i, j int) bool {
					return ints[i] == ints[j]
				},
			},
		},
		{
			name: "empty []int",
			args: args{
				data: func() *[]int {
					t := []int{}
					return &t
				}(),
				less:  nil,
				equal: nil,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			uniq.Slice(tt.args.data, tt.args.less, tt.args.equal)
			if !uniq.IsSliceSorted(tt.args.data, tt.args.less, tt.args.equal) {
				t.Fail()
			}
		})
	}
}

func BenchmarkSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		uniq.Sort(&uniq.IntSlice{1, 9, 9, 4, 0, 9, 2, 6, 2, 0, 1, 4, 0, 7, 1, 3})
	}
}
