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
			name: "empty",
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

func TestStable(t *testing.T) {
	type args struct {
		data sort.Interface
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "empty",
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
			tt.args.data = rv.Slice(0, uniq.Stable(tt.args.data)).Interface().(sort.Interface)
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
		slice interface{}
		less  func(i, j int) bool
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "empty IntSlice",
			args: args{
				slice: sort.IntSlice{},
				less:  nil,
			},
		},
		{
			name: "IntSlice",
			args: args{
				slice: intSlice,
				less:  intSlice.Less,
			},
		},
		{
			name: "[]int",
			args: args{
				slice: ints,
				less: func(i, j int) bool {
					return ints[i] < ints[j]
				},
			},
		},
		{
			name: "empty []int",
			args: args{
				slice: []int{},
				less:  nil,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Log(tt.args.slice)
			defer func() {
				t.Log(tt.args.slice)
			}()
			rv := reflect.ValueOf(tt.args.slice)
			tt.args.slice = rv.Slice(0, uniq.Slice(tt.args.slice, tt.args.less)).Interface()
			if !uniq.SliceIsSorted(tt.args.slice, tt.args.less) {
				t.Fail()
				t.Log(tt.args.slice)
			}
		})
	}
}

func TestSliceStable(t *testing.T) {
	intSlice := sort.IntSlice{1, 9, 9, 4, 0, 9, 2, 6, 2, 0, 1, 4, 0, 7, 1, 3}
	ints := []int{1, 9, 9, 4, 0, 9, 2, 6, 2, 0, 1, 4, 0, 7, 1, 3}

	type args struct {
		slice interface{}
		less  func(i, j int) bool
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "empty IntSlice",
			args: args{
				slice: sort.IntSlice{},
				less:  nil,
			},
		},
		{
			name: "IntSlice",
			args: args{
				slice: intSlice,
				less:  intSlice.Less,
			},
		},
		{
			name: "[]int",
			args: args{
				slice: ints,
				less: func(i, j int) bool {
					return ints[i] < ints[j]
				},
			},
		},
		{
			name: "empty []int",
			args: args{
				slice: []int{},
				less:  nil,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Log(tt.args.slice)
			defer func() {
				t.Log(tt.args.slice)
			}()
			rv := reflect.ValueOf(tt.args.slice)
			tt.args.slice = rv.Slice(0, uniq.SliceStable(tt.args.slice, tt.args.less)).Interface()
			if !uniq.SliceIsSorted(tt.args.slice, tt.args.less) {
				t.Fail()
				t.Log(tt.args.slice)
			}
		})
	}
}

func BenchmarkSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		uniq.Sort(&sort.IntSlice{1, 9, 9, 4, 0, 9, 2, 6, 2, 0, 1, 4, 0, 7, 1, 3})
	}
}

func TestIsSorted(t *testing.T) {
	ints := []int{1, 9, 9, 4, 0, 9, 2, 6, 2, 0, 1, 4, 0, 7, 1, 3}
	soredInts := []int{1, 2, 4, 5, 6, 7, 8, 9}

	type args struct {
		data sort.Interface
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "not sorted",
			args: args{
				data: sort.IntSlice(ints),
			},
			want: false,
		},
		{
			name: "sorted",
			args: args{
				data: sort.IntSlice(soredInts),
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := uniq.IsSorted(tt.args.data); got != tt.want {
				t.Errorf("IsSorted() = %v, want %v", got, tt.want)
			}
		})
	}
}
