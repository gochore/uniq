package uniq_test

import (
	"github.com/gochore/uniq"
	"math/rand"
)

func RandomIntSlice(length, min, max int) uniq.IntSlice {
	ret := make(uniq.IntSlice, length)
	for i := range ret {
		ret[i] = min + rand.Intn(max+1-min)
	}
	return ret
}

func RandomFloat64Slice(length, min, max int) uniq.Float64Slice {
	ret := make(uniq.Float64Slice, length)
	for i := range ret {
		ret[i] = min + rand.Intn(max+1-min)
	}
	return ret
}

func RandomStringSlice(length, min, max int) uniq.StringSlice {
	ret := make(uniq.StringSlice, length)
	for i := range ret {
		ret[i] = min + rand.Intn(max+1-min)
	}
	return ret
}