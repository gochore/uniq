package uniq_test

import (
	"math"
	"testing"

	"github.com/gochore/uniq"
)

func TestInts(t *testing.T) {
	type args struct {
		s []int
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "nil",
			args: args{
				s: nil,
			},
		},
		{
			name: "empty",
			args: args{
				s: []int{},
			},
		},
		{
			name: "[]int",
			args: args{
				s: []int{1, 9, 9, 4, 0, 9, 2, 6, 2, 0, 1, 4, 0, 7, 1, 3},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Log(tt.args.s)
			defer func() {
				t.Log(tt.args.s)
			}()
			tt.args.s = tt.args.s[:uniq.Ints(tt.args.s)]
			if !uniq.IntsAreSorted(tt.args.s) {
				t.Fail()
				t.Log(tt.args.s)
			}
		})
	}
}

func TestFloat64s(t *testing.T) {
	type args struct {
		s []float64
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "nil",
			args: args{
				s: nil,
			},
		},
		{
			name: "empty",
			args: args{
				s: []float64{},
			},
		},
		{
			name: "[]float64",
			args: args{
				s: []float64{1, 9, 9, 4, 0, 9, 2, 6, 2, 0, 1, 4, 0, 7, 1, 3},
			},
		},
		{
			name: "[]float64 with NaN",
			args: args{
				s: []float64{math.NaN(), 1, 9, 9, 4, 0, 9, 2, 6, math.NaN(), 2, 0, 1, 4, 0, 7, 1, 3, math.NaN()},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Log(tt.args.s)
			defer func() {
				t.Log(tt.args.s)
			}()
			tt.args.s = tt.args.s[:uniq.Float64s(tt.args.s)]
			if !uniq.Float64sAreSorted(tt.args.s) {
				t.Fail()
				t.Log(tt.args.s)
			}
		})
	}
}

func TestStrings(t *testing.T) {
	type args struct {
		s []string
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "nil",
			args: args{
				s: nil,
			},
		},
		{
			name: "empty",
			args: args{
				s: []string{},
			},
		},
		{
			name: "[]string",
			args: args{
				s: []string{"1", "9", "9", "4", "0", "9", "2", "6", "2", "0", "1", "4", "0", "7", "1", "3"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Log(tt.args.s)
			defer func() {
				t.Log(tt.args.s)
			}()
			tt.args.s = tt.args.s[:uniq.Strings(tt.args.s)]
			if !uniq.StringsAreSorted(tt.args.s) {
				t.Fail()
				t.Log(tt.args.s)
			}
		})
	}
}
