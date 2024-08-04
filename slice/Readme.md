## Slice

Slice is a dynamically-sized, flexible **view** into the elements of an fixed-length array.

stdlib [slices](https://pkg.go.dev/slices) also defines various functions useful with slices of any type.

**Before getting started, u need to import these packages in src.**

```go
import (
    "github.com/KokoiRuby/generickit/slice"
    "golang.org/x/exp/constraints"
)
```

### Insert

> func Insert

```go
func Insert[T any](dst []T, value T, idx int) ([]T, error)
```

Insert a `value` at `idx` of `dst` slice, returning a modified slice. This function is `O(len(dst))`.

Example:

```go
func ExampleInsert() {
	res, _ := slice.Insert[int]([]int{1, 2, 3}, 0, 1)
	fmt.Println(res)
	// Output:
	// [1 0 2 3]
}
```

> func InsertSlice

```go
func InsertSlice[T any](dst []T, src []T, idx int) ([]T, error)
```

Insert a `src` slice at `idx` of `dst` slice, returning a modified slice. This function is `O(len(src) + len(dst))`.

Example:

```go
func ExampleInsertSlice() {
	res, _ := slice.InsertSlice[int]([]int{1, 2, 3}, []int{4, 5}, 1)
	fmt.Println(res)
	// Output:
	// [1 4 5 2 3]
}
```

### Delete

> func Delete

```go
Delete[T any](dst []T, idx int) ([]T, error)
```

Delete a element at `idx` of `dst` slice, returning a modified slice. This function is `O(len(dst))`.

Example:

```go 
func ExampleDelete() {
	res, _ := slice.Delete[int]([]int{1, 2, 3}, 1)
	fmt.Println(res)
	// output:
	// [1 3]
}
```

> func DeleteAfter

```go
func DeleteAfter[T any](dst []T, idx int) ([]T, error)
```

Delete all elements after `idx` of `dst` slice, returning a modified slice. This function is `O(len(dst))`.

Example:

```go
func ExampleDeleteAfter() {
	res, _ := slice.DeleteAfter[int]([]int{1, 2, 3, 4, 5}, 2)
	fmt.Println(res)
	// output:
	// [1 2]
}
```

> func DeleteRange

```go
func DeleteRange[T any](dst []T, r int, idx int) ([]T, error)
```

Delete a range of elements starting from `idx` of `dst` slice, returning a modified slice. This function is `O(len(dst))`.

Example:

```go
func ExampleDeleteRange() {
	res, _ := slice.DeleteRange[int]([]int{1, 2, 3, 4, 5}, 3, 1)
	fmt.Println(res)
	// output:
	// [1 5]
}
```

> func DeleteVal

```go
func DeleteVal[T constraints.Ordered](dst []T, val T) ([]T, error)
```

Delete a specific `val` in `dst` slice, returning a modified slice. This function is `O(len(dst))`.

Notes: `T` must be `constraints.Ordered`.

```go
type Ordered interface {
	Integer | Float | ~string
}
```

Example:

```go
func ExampleDeleteVal() {
	res, _ := slice.DeleteVal[int]([]int{1, 2, 3, 4, 5}, 2)
	fmt.Println(res)
	// output:
	// [1 3 4 5]
}
```

### Aggregate

Note: `T` in Max/Min must be `RealNumber`, and `Number` in Sum.

Note: U must handle precision by urself if using floating-point.

```go
type RealNumber interface {
	~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 |
		~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~float32 | ~float64
}

type Number interface {
	RealNumber | ~complex64 | ~complex128
}
```

> func Max

```go
func Max[T generickit.RealNumber](sl []T) T
```

Returning the maximum element in slice. This function is `O(len(sl))`.

Example:

```go
func ExampleMax() {
	res := slice.Max[int]([]int{1, 2, 3, 4, 5})
	fmt.Println(res)
	// output:
	// 5
}
```

> func Min

```go
func Min[T generickit.RealNumber](sl []T) T
```

Returning the minimum element in slice. This function is `O(len(sl))`.

Example:

```go
func ExampleMin() {
	res := slice.Min[int]([]int{1, 2, 3, 4, 5})
	fmt.Println(res)
	// output:
	// 1
}
```

> func Sum

```go
func Sum[T generickit.Number](sl []T) T
```

Returning the sum of all element in slice. This function is `O(len(sl))`.

Example:

```go
func ExampleSum() {
	res := slice.Sum[int]([]int{1, 2, 3, 4, 5})
	fmt.Println(res)
	// output:
	// 15
}
```

### Reverse

> func Reverse

```go
func Reverse[T any](s []T) []T
```

Reverse the order of slice **in-place**, returning a modified slice. This function is `O(len(s))`.

Example:

```go
func ExampleReverse() {
	res := slice.Reverse[int]([]int{5, 4, 3, 2, 1})
	fmt.Println(res)
	// output:
	// [1 2 3 4 5]
}
```

### Shrink

Note: Golang [Slice](https://pkg.go.dev/slices) does not support shrinking primitively, it only supports growth where strategy: 2x for small slices (< 256) & 1.25x for large slices (>= 256).

> func Shrink

```go
func Shrink[T any](src []T) []T
```

Shrink `src` slice in-place, returning the slice with shrinked capacity. This function is `O(cap(s))`.

Example:

```go
func ExampleShrink() {
	sl := make([]int, 100, 300)
	sl = slice.Shrink[int](sl)
	fmt.Println(cap(sl))
	sl = make([]int, 300, 480)
	sl = slice.Shrink[int](sl)
	fmt.Println(cap(sl))
	// output:
	// 150
	// 384
}
```

### Generator

> func Generator

```go
func Generator[T any](sl []T) <-chan T
```

A Python featured yield-like generator implmented by Go using goroutine & channel.

Example:

```go
func ExampleGenerator() {
	intGenerator := slice.Generator[int]([]int{1, 2, 3})
	fmt.Println(<-intGenerator)
	fmt.Println(<-intGenerator)
	fmt.Println(<-intGenerator)
	fmt.Println(<-intGenerator)
	// output:
	// 1
	// 2
	// 3
	// 0
}
```

### Contains

> func Contains

```go
func Contains[T comparable](sl []T, val T) bool
```

Check if `sl` contains `val`. This function is `O(len(sl))`.

Example:

```go
func ExampleContains() {
	fmt.Println(slice.Contains[int]([]int{1, 2, 3, 4, 5}, 5))
	// output:
	// true
}
```

### Find

> func Find

```go
func Find[T comparable](sl []T, val T) (idx int, isFound bool)
```

Find `val` in `sl`, returing the 1st idx & flag. This function is `O(len(sl))`.

Example:

```go
func ExampleFind() {
	idx, isFound := slice.Find[int]([]int{1, 2, 3, 4, 5, 5}, 3)
	fmt.Println(idx)
	fmt.Println(isFound)
	// output:
	// 2
	// true
}
```

> func FindLast

```go
func FindLast[T comparable](sl []T, val T) (idx int, isFound bool) {
```

Find `val` in `sl`, returing the last idx & flag. This function is `O(len(sl))`.

Example:

```go
func ExampleFindLast() {
	idx, isFound := slice.Find[int]([]int{1, 2, 3, 4, 5, 5}, 5)
	fmt.Println(idx)
	fmt.Println(isFound)
	// output:
	// 5
	// true
}
```

> func FindAll

```go
func FindAll[T comparable](sl []T, val T) (idx []int, isFound bool)
```

Find all val in sl, returning idx slice & flag. This function is `O(len(sl))`.

Example:

```go
func ExampleFindAll() {
	idxSlice, isFound := slice.FindAll[int]([]int{1, 2, 3, 4, 5, 5}, 5)
	fmt.Println(idxSlice)
	fmt.Println(isFound)
	// output:
	// [4 5]
	// true
}
```

### MapReduce

**MapReduce** is a programming model used for parallel processing of large datasets.

- **Map Phase**: the input dataset is divided into smaller subsets, which are processed in parallel. Each subset is processed by a Map function, transforming it **into key-value pairs**.
- **Reduce Phase**: In this stage, the key-value pairs produced by the Map phase are **aggregated based on keys**, and the values for the same key are merged to produce the final output.

> func Map

```go
func Map[Src any, Dst any](src []Src, mapFunc func(src Src) Dst) []Dst
```

Map a `Src` type of `src` slice to another slice with type `Dst`. This function is `O(len(src))`.

Example:

```go
func ExampleMap() {
	after := slice.Map[int, string]([]int{1, 2, 3, 4, 5}, func(src int) string { 
        return strconv.Itoa(src) 
    })
	fmt.Println(after)
	// output:
	// [1 2 3 4 5]
}
```

> func Filter

```go
func Filter[Src any, Dst any](src []Src, mapFunc func(src Src) (dst Dst, filter bool)) []Dst
```

Map a `Src` type of `src` slice to another slice with type `Dst` given a filter. This function is `O(len(src))`.

Example:

```go
func ExampleFilterMap() {
	after := slice.Filter[int, string]([]int{1, 2, 3, 4, 5}, func(src int) (string, bool) { 
        return strconv.Itoa(src), src >= 3 
    })
	fmt.Println(after)
	// output:
	// [3 4 5]
}
```

> func Reduce

```go
func Reduce[Src any, Dst Number](src []Src, reducer func(src Src) Dst) Dst
```

Reduce a `Src` type of `src` slice to a Dst type. This function is `O(len(src))`.

Note: Reduce usually for **aggregation**, Dst in `Number` type in order to support operators.

```go
type Number interface {
	constraints.Integer | constraints.Float
}
```

Example:

```go
func ExampleReduce() {
	after := slice.Reduce[int, int]([]int{1, 2, 3, 4, 5}, func(src int) int { return src })
	fmt.Println(after)
	// output:
	// 15
}
```

