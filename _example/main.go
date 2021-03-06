package main

import (
	"fmt"
	"sort"

	"github.com/gochore/uniq"
)

func main() {
	ints1 := []int{1, 2, 1, 2, 4, 4, 4, 3, 3}
	ints2 := make([]int, len(ints1))
	copy(ints2, ints1)

	fmt.Println(ints1)
	// [1 2 1 2 4 4 4 3 3]
	sort.Ints(ints1)
	fmt.Println(ints1)
	// [1 1 2 2 3 3 4 4 4]
	ints2 = ints2[:uniq.Ints(ints2)]
	fmt.Println(ints2)
	// [1 2 3 4]

	kvs1 := []KV{{"a", 5}, {"b", 2}, {"a", 5}, {"b", 4}, {"c", 9}}
	less1 := func(i, j int) bool {
		return kvs1[i].K < kvs1[j].K
	}
	kvs2 := make([]KV, len(kvs1))
	copy(kvs2, kvs1)
	less2 := func(i, j int) bool {
		return kvs2[i].K < kvs2[j].K || kvs2[i].K == kvs2[j].K && kvs2[i].V < kvs2[j].V
	}
	fmt.Println(kvs1)
	// [{a 5} {b 2} {a 5} {b 4} {c 9}]
	sort.Slice(kvs1, less1)
	fmt.Println(kvs1)
	// [{a 5} {a 5} {b 2} {b 4} {c 9}]
	kvs2 = kvs2[:uniq.Slice(kvs2, less2)]
	fmt.Println(kvs2)
	// [{a 5} {b 2} {b 4} {c 9}]
}

type KV struct {
	K string
	V int
}
