package uniq_test

import (
	"math/rand"

	"github.com/gochore/uniq"
)

func RandomIntSlice(length, min, max int) *uniq.IntSlice {
	ret := make(uniq.IntSlice, length)
	for i := range ret {
		ret[i] = min + rand.Intn(max+1-min)
	}
	return &ret
}

func RandomFloat64Slice(length int) *uniq.Float64Slice {
	ret := make(uniq.Float64Slice, length)
	for i := range ret {
		ret[i] = rand.Float64()
	}
	return &ret
}
