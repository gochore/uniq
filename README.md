# uniq

[![Go Action](https://github.com/gochore/uniq/workflows/Go/badge.svg)](https://github.com/gochore/uniq/actions)
[![Go Report Card](https://goreportcard.com/badge/github.com/gochore/uniq)](https://goreportcard.com/report/github.com/gochore/uniq)
[![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/gochore/uniq)](https://github.com/gochore/uniq/blob/master/go.mod)
[![GitHub tag (latest by date)](https://img.shields.io/github/v/tag/gochore/uniq)](https://github.com/gochore/uniq/releases)

Sort and deduplicate data.

## Example

```go
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
```

## Document

#### func  Float64s

```go
func Float64s(s []float64) int
```
Float64s sorts and deduplicates a slice of float64s in increasing order
(not-a-number values are treated as less than other values).

#### func  Float64sAreSorted

```go
func Float64sAreSorted(a []float64) bool
```
Float64sAreSorted tests whether a slice of float64s is sorted and deduplicated
in increasing order (not-a-number values are treated as less than other values).

#### func  Ints

```go
func Ints(s []int) int
```
Ints sorts and deduplicates a slice of ints in increasing order.

#### func  IntsAreSorted

```go
func IntsAreSorted(a []int) bool
```
IntsAreSorted tests whether a slice of ints is sorted and deduplicated in
increasing order.

#### func  IsSorted

```go
func IsSorted(data sort.Interface) bool
```
IsSorted reports if data is sorted and deduplicated.

#### func  Slice

```go
func Slice(slice interface{}, less func(i, j int) bool) int
```
Slice sorts and deduplicates the provided slice given the provided less
function.

#### func  SliceIsSorted

```go
func SliceIsSorted(data interface{}, less func(i, j int) bool) bool
```
SliceIsSorted reports if slice is sorted and deduplicated.

#### func  SliceStable

```go
func SliceStable(slice interface{}, less func(i, j int) bool) int
```
SliceStable sorts and deduplicates the provided slice given the provided less
function while keeping the original order of equal elements.

#### func  Sort

```go
func Sort(data sort.Interface) int
```
Sort sorts and deduplicated data.

#### func  Stable

```go
func Stable(data sort.Interface) int
```
Stable sorts and deduplicates data while keeping the original order of equal
elements.

#### func  Strings

```go
func Strings(s []string) int
```
Strings sorts and deduplicates a slice of strings in increasing order.

#### func  StringsAreSorted

```go
func StringsAreSorted(a []string) bool
```
StringsAreSorted tests whether a slice of strings is sorted and deduplicated in
increasing order.

