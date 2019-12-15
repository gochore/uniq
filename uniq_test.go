package uniq_test

import (
	"testing"

	"github.com/gochore/uniq"
)

func TestUniq(t *testing.T) {
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
			name: "IntSlice",
			args: args{
				data: RandomIntSlice(1000, -100, 100),
			},
		},
		{
			name: "Float64Slice",
			args: args{
				data: RandomFloat64Slice(1000),
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
