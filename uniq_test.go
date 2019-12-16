package uniq_test

import (
	"reflect"
	"sort"
	"testing"

	"github.com/gochore/uniq"
)

func TestSort(t *testing.T) {
	type args struct {
		data sort.Interface
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "emtpy",
			args: args{
				data: sort.IntSlice{},
			},
		},
		{
			name: "IntSlice",
			args: args{
				data: sort.IntSlice{1, 9, 9, 4, 0, 9, 2, 6, 2, 0, 1, 4, 0, 7, 1, 3},
			},
		},
		{
			name: "Float64Slice",
			args: args{
				data: sort.Float64Slice{1, 9, 9, 4, 0, 9, 2, 6, 2, 0, 1, 4, 0, 7, 1, 3},
			},
		},
		{
			name: "StringSlice",
			args: args{
				data: sort.StringSlice{"1", "9", "9", "4", "0", "9", "2", "6", "2", "0", "1", "4", "0", "7", "1", "3"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Log(tt.args.data)
			defer func() {
				t.Log(tt.args.data)
			}()
			uniq.Sort(tt.args.data)
			rv := reflect.ValueOf(tt.args.data)
			tt.args.data = rv.Slice(0, uniq.Sort(tt.args.data)).Interface().(sort.Interface)
			if !uniq.IsSorted(tt.args.data) {
				t.Fail()
				t.Log(tt.args.data)
			}
		})
	}
}

func TestSlice(t *testing.T) {
	intSlice := sort.IntSlice{1, 9, 9, 4, 0, 9, 2, 6, 2, 0, 1, 4, 0, 7, 1, 3}
	ints := []int{1, 9, 9, 4, 0, 9, 2, 6, 2, 0, 1, 4, 0, 7, 1, 3}

	type args struct {
		data interface{}
		less func(i, j int) bool
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "empty IntSlice",
			args: args{
				data: sort.IntSlice{},
				less: nil,
			},
		},
		{
			name: "IntSlice",
			args: args{
				data: intSlice,
				less: intSlice.Less,
			},
		},
		{
			name: "[]int",
			args: args{
				data: ints,
				less: func(i, j int) bool {
					return ints[i] < ints[j]
				},
			},
		},
		{
			name: "empty []int",
			args: args{
				data: []int{},
				less: nil,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Log(tt.args.data)
			defer func() {
				t.Log(tt.args.data)
			}()
			rv := reflect.ValueOf(tt.args.data)
			tt.args.data = rv.Slice(0, uniq.Slice(tt.args.data, tt.args.less)).Interface()
			if !uniq.SliceIsSorted(tt.args.data, tt.args.less) {
				t.Fail()
				t.Log(tt.args.data)
			}
		})
	}
}

func BenchmarkSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		uniq.Sort(&sort.IntSlice{1, 9, 9, 4, 0, 9, 2, 6, 2, 0, 1, 4, 0, 7, 1, 3})
	}
}
