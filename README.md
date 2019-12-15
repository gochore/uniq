# uniq

[![Go Action](https://github.com/gochore/uniq/workflows/Go/badge.svg)](https://github.com/gochore/uniq/actions)
[![Go Report Card](https://goreportcard.com/badge/github.com/gochore/uniq)](https://goreportcard.com/report/github.com/gochore/uniq)
[![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/gochore/uniq)](https://github.com/gochore/uniq/blob/master/go.mod)
[![GitHub tag (latest by date)](https://img.shields.io/github/v/tag/gochore/uniq)](https://github.com/gochore/uniq/releases)

Sort and deduplicate data.

```go
package main

import (
	"fmt"

	"github.com/gochore/uniq"
)

func main() {
	ints := []int{1, 2, 1, 2, 4, 4, 4, 3, 3}
	fmt.Println(ints)
	// [1 2 1 2 4 4 4 3 3]
	uniq.Ints(&ints)
	fmt.Println(ints)
	// [1 2 3 4]

	kvs := []KV{{"a", 1}, {"b", 2}, {"a", 1}, {"b", 4}, {"c", 5}}
	less := func(i, j int) bool {
		return kvs[i].K < kvs[j].K
	}
	equal := func(i, j int) bool {
		return kvs[i].K == kvs[j].K && kvs[i].V == kvs[j].V
	}
	fmt.Println(kvs)
	// [{a 1} {b 2} {a 1} {b 4} {c 5}]
	uniq.Slice(&kvs, less, equal)
	fmt.Println(kvs)
	//[{a 1} {b 2} {b 4} {c 5}]
}

type KV struct {
	K string
	V int
}
```
