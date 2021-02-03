package uniq

import "testing"

func TestBytes(t *testing.T) {
	type args struct {
		s []byte
	}

	tests := []struct {
		name string
		args args
	}{
		{
			name: "regular",
			args: args{
				s: []byte{1, 9, 9, 4, 0, 9, 2, 6, 2, 0, 1, 4, 0, 7, 1, 3},
			},
		},
		{
			name: "nil",
			args: args{
				s: nil,
			},
		},
		{
			name: "empty",
			args: args{
				s: []byte{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Log(tt.args.s)
			defer func() {
				t.Log(tt.args.s)
			}()
			tt.args.s = tt.args.s[:Bytes(tt.args.s)]
			if !BytesAreSorted(tt.args.s) {
				t.Fail()
				t.Log(tt.args.s)
			}
		})
	}
}

func TestFloat32s(t *testing.T) {
	type args struct {
		s []float32
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "regular",
			args: args{
				s: []float32{1, 9, 9, 4, 0, 9, 2, 6, 2, 0, 1, 4, 0, 7, 1, 3},
			},
		},
		{
			name: "nil",
			args: args{
				s: nil,
			},
		},
		{
			name: "empty",
			args: args{
				s: []float32{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Log(tt.args.s)
			defer func() {
				t.Log(tt.args.s)
			}()
			tt.args.s = tt.args.s[:Float32s(tt.args.s)]
			if !Float32sAreSorted(tt.args.s) {
				t.Fail()
				t.Log(tt.args.s)
			}
		})
	}
}

func TestInt16s(t *testing.T) {
	type args struct {
		s []int16
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "regular",
			args: args{
				s: []int16{1, 9, 9, 4, 0, 9, 2, 6, 2, 0, 1, 4, 0, 7, 1, 3},
			},
		},
		{
			name: "nil",
			args: args{
				s: nil,
			},
		},
		{
			name: "empty",
			args: args{
				s: []int16{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Log(tt.args.s)
			defer func() {
				t.Log(tt.args.s)
			}()
			tt.args.s = tt.args.s[:Int16s(tt.args.s)]
			if !Int16sAreSorted(tt.args.s) {
				t.Fail()
				t.Log(tt.args.s)
			}
		})
	}
}

func TestInt32s(t *testing.T) {
	type args struct {
		s []int32
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "regular",
			args: args{
				s: []int32{1, 9, 9, 4, 0, 9, 2, 6, 2, 0, 1, 4, 0, 7, 1, 3},
			},
		},
		{
			name: "nil",
			args: args{
				s: nil,
			},
		},
		{
			name: "empty",
			args: args{
				s: []int32{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Log(tt.args.s)
			defer func() {
				t.Log(tt.args.s)
			}()
			tt.args.s = tt.args.s[:Int32s(tt.args.s)]
			if !Int32sAreSorted(tt.args.s) {
				t.Fail()
				t.Log(tt.args.s)
			}
		})
	}
}

func TestInt64s(t *testing.T) {
	type args struct {
		s []int64
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "regular",
			args: args{
				s: []int64{1, 9, 9, 4, 0, 9, 2, 6, 2, 0, 1, 4, 0, 7, 1, 3},
			},
		},
		{
			name: "nil",
			args: args{
				s: nil,
			},
		},
		{
			name: "empty",
			args: args{
				s: []int64{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Log(tt.args.s)
			defer func() {
				t.Log(tt.args.s)
			}()
			tt.args.s = tt.args.s[:Int64s(tt.args.s)]
			if !Int64sAreSorted(tt.args.s) {
				t.Fail()
				t.Log(tt.args.s)
			}
		})
	}
}

func TestInt8s(t *testing.T) {
	type args struct {
		s []int8
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "regular",
			args: args{
				s: []int8{1, 9, 9, 4, 0, 9, 2, 6, 2, 0, 1, 4, 0, 7, 1, 3},
			},
		},
		{
			name: "nil",
			args: args{
				s: nil,
			},
		},
		{
			name: "empty",
			args: args{
				s: []int8{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Log(tt.args.s)
			defer func() {
				t.Log(tt.args.s)
			}()
			tt.args.s = tt.args.s[:Int8s(tt.args.s)]
			if !Int8sAreSorted(tt.args.s) {
				t.Fail()
				t.Log(tt.args.s)
			}
		})
	}
}

func TestRunes(t *testing.T) {
	type args struct {
		s []rune
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "regular",
			args: args{
				s: []rune{1, 9, 9, 4, 0, 9, 2, 6, 2, 0, 1, 4, 0, 7, 1, 3},
			},
		},
		{
			name: "nil",
			args: args{
				s: nil,
			},
		},
		{
			name: "empty",
			args: args{
				s: []rune{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Log(tt.args.s)
			defer func() {
				t.Log(tt.args.s)
			}()
			tt.args.s = tt.args.s[:Runes(tt.args.s)]
			if !RunesAreSorted(tt.args.s) {
				t.Fail()
				t.Log(tt.args.s)
			}
		})
	}
}

func TestUint16s(t *testing.T) {
	type args struct {
		s []uint16
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "regular",
			args: args{
				s: []uint16{1, 9, 9, 4, 0, 9, 2, 6, 2, 0, 1, 4, 0, 7, 1, 3},
			},
		},
		{
			name: "nil",
			args: args{
				s: nil,
			},
		},
		{
			name: "empty",
			args: args{
				s: []uint16{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Log(tt.args.s)
			defer func() {
				t.Log(tt.args.s)
			}()
			tt.args.s = tt.args.s[:Uint16s(tt.args.s)]
			if !Uint16sAreSorted(tt.args.s) {
				t.Fail()
				t.Log(tt.args.s)
			}
		})
	}
}

func TestUint32s(t *testing.T) {
	type args struct {
		s []uint32
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "regular",
			args: args{
				s: []uint32{1, 9, 9, 4, 0, 9, 2, 6, 2, 0, 1, 4, 0, 7, 1, 3},
			},
		},
		{
			name: "nil",
			args: args{
				s: nil,
			},
		},
		{
			name: "empty",
			args: args{
				s: []uint32{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Log(tt.args.s)
			defer func() {
				t.Log(tt.args.s)
			}()
			tt.args.s = tt.args.s[:Uint32s(tt.args.s)]
			if !Uint32sAreSorted(tt.args.s) {
				t.Fail()
				t.Log(tt.args.s)
			}
		})
	}
}

func TestUint64s(t *testing.T) {
	type args struct {
		s []uint64
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "regular",
			args: args{
				s: []uint64{1, 9, 9, 4, 0, 9, 2, 6, 2, 0, 1, 4, 0, 7, 1, 3},
			},
		},
		{
			name: "nil",
			args: args{
				s: nil,
			},
		},
		{
			name: "empty",
			args: args{
				s: []uint64{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Log(tt.args.s)
			defer func() {
				t.Log(tt.args.s)
			}()
			tt.args.s = tt.args.s[:Uint64s(tt.args.s)]
			if !Uint64sAreSorted(tt.args.s) {
				t.Fail()
				t.Log(tt.args.s)
			}
		})
	}
}

func TestUint8s(t *testing.T) {
	type args struct {
		s []uint8
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "regular",
			args: args{
				s: []byte{1, 9, 9, 4, 0, 9, 2, 6, 2, 0, 1, 4, 0, 7, 1, 3},
			},
		},
		{
			name: "nil",
			args: args{
				s: nil,
			},
		},
		{
			name: "empty",
			args: args{
				s: []byte{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Log(tt.args.s)
			defer func() {
				t.Log(tt.args.s)
			}()
			tt.args.s = tt.args.s[:Uint8s(tt.args.s)]
			if !Uint8sAreSorted(tt.args.s) {
				t.Fail()
				t.Log(tt.args.s)
			}
		})
	}
}

func TestUintptrs(t *testing.T) {
	type args struct {
		s []uintptr
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "regular",
			args: args{
				s: []uintptr{1, 9, 9, 4, 0, 9, 2, 6, 2, 0, 1, 4, 0, 7, 1, 3},
			},
		},
		{
			name: "nil",
			args: args{
				s: nil,
			},
		},
		{
			name: "empty",
			args: args{
				s: []uintptr{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Log(tt.args.s)
			defer func() {
				t.Log(tt.args.s)
			}()
			tt.args.s = tt.args.s[:Uintptrs(tt.args.s)]
			if !UintptrsAreSorted(tt.args.s) {
				t.Fail()
				t.Log(tt.args.s)
			}
		})
	}
}

func TestUints(t *testing.T) {
	type args struct {
		s []uint
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "regular",
			args: args{
				s: []uint{1, 9, 9, 4, 0, 9, 2, 6, 2, 0, 1, 4, 0, 7, 1, 3},
			},
		},
		{
			name: "nil",
			args: args{
				s: nil,
			},
		},
		{
			name: "empty",
			args: args{
				s: []uint{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Log(tt.args.s)
			defer func() {
				t.Log(tt.args.s)
			}()
			tt.args.s = tt.args.s[:Uints(tt.args.s)]
			if !UintsAreSorted(tt.args.s) {
				t.Fail()
				t.Log(tt.args.s)
			}
		})
	}
}
