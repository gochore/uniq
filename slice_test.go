package uniq_test

import (
	"math"
	"testing"

	"github.com/gochore/uniq"
)

func TestInts(t *testing.T) {
	type args struct {
		s *[]int
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
				s: func() *[]int {
					t := []int{}
					return &t
				}(),
			},
		},
		{
			name: "[]int",
			args: args{
				s: func() *[]int {
					t := []int{1, 9, 9, 4, 0, 9, 2, 6, 2, 0, 1, 4, 0, 7, 1, 3}
					return &t
				}(),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			uniq.Ints(tt.args.s)
			if tt.args.s != nil {
				if !uniq.IntsAreSorted(*tt.args.s) {
					t.Fail()
				}
			}
		})
	}
}

func TestFloat64s(t *testing.T) {
	type args struct {
		s *[]float64
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
				s: func() *[]float64 {
					t := []float64{}
					return &t
				}(),
			},
		},
		{
			name: "[]float64",
			args: args{
				s: func() *[]float64 {
					t := []float64{1, 9, 9, 4, 0, 9, 2, 6, 2, 0, 1, 4, 0, 7, 1, 3}
					return &t
				}(),
			},
		},
		{
			name: "[]float64 with NaN",
			args: args{
				s: func() *[]float64 {
					t := []float64{math.NaN(), 1, 9, 9, 4, 0, 9, 2, 6, math.NaN(), 2, 0, 1, 4, 0, 7, 1, 3, math.NaN()}
					return &t
				}(),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			uniq.Float64s(tt.args.s)
			if tt.args.s != nil {
				if !uniq.Float64sAreSorted(*tt.args.s) {
					t.Fail()
				}
			}
		})
	}
}

func TestStrings(t *testing.T) {
	type args struct {
		s *[]string
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
				s: func() *[]string {
					t := []string{}
					return &t
				}(),
			},
		},
		{
			name: "[]string",
			args: args{
				s: func() *[]string {
					t := []string{"1", "9", "9", "4", "0", "9", "2", "6", "2", "0", "1", "4", "0", "7", "1", "3"}
					return &t
				}(),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			uniq.Strings(tt.args.s)
			if tt.args.s != nil {
				if !uniq.StringsAreSorted(*tt.args.s) {
					t.Fail()
				}
			}
		})
	}
}
