## Slice

[Slice](https://pkg.go.dev/slices) is a dynamically-sized, flexible **view** into the elements of an fixed-length array.

Before getting started, u need to import this package in src.

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

### Shirnk

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
	sl = slice.Shrink(sl)
	fmt.Println(cap(sl))
	sl = make([]int, 300, 480)
	sl = slice.Shrink(sl)
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

A Python featured yield-like generator implmented by Go using Channel.

Example:

```go
func ExampleGenerator() {
	intGenerator := slice.Generator([]int{1, 2, 3})
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
	fmt.Println(Contains([]int{1, 2, 3, 4, 5}, 5))
	// output:
	// true
}
```

### Find

> func Find

```go
func Find[T comparable](sl []T, val T) (idx int, isFound bool)
```

Find `val` in `sl`, returing 1st idx & flag. This function is `O(len(sl))`.

Example:

```go
func ExampleFind() {
	idx, isFound := Find([]int{1, 2, 3, 4, 5, 5}, 3)
	fmt.Println(idx)
	fmt.Println(isFound)
	// output:
	// 2
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
	idxSlice, isFound := FindAll([]int{1, 2, 3, 4, 5, 5}, 5)
	fmt.Println(idxSlice)
	fmt.Println(isFound)
	// output:
	// [4 5]
	// true
}
```

