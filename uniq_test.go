package uniq_test

import (
	"math/rand"
	"testing"

	"github.com/gochore/uniq"
)

func TestUniq(t *testing.T) {
	data := RandomIntSlice(1000, -100, 100)
	uniq.Uniq(&data)
	if !uniq.IsUniqed(data) {
		t.FailNow()
	}
}

func TestSlice(t *testing.T) {
	data := RandomIntSlice(1000, -100, 100)
	uniq.Slice(&data, func(i, j int) bool {
		return data[i] < data[j]
	}, func(i, j int) bool {
		return data[i] == data[j]
	})
	if !uniq.IsUniqed(data) {
		t.FailNow()
	}
}

func RandomIntSlice(length, min, max int) uniq.IntSlice {
	ret := make(uniq.IntSlice, length)
	for i := range ret {
		ret[i] = min + rand.Intn(max+1-min)
	}
	return ret
}
