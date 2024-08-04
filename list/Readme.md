## List

list represents a collection of elements, where each is connected to the next through references or pointers.

stdlib container/[list](https://pkg.go.dev/container/list) implements a doubly linked list.

**Before getting started, u need to import these packages in src.**

```go
import (
    "github.com/KokoiRuby/generickit/list"
)
```

List Interface:

```go
type List[T any] interface {
	// Get the element given idx, err if out of range
	Get(idx int) T

	// Append a element(s)
	Append(ts ...T)

	// Add an element at idx, err if out of range
	// Append if idx == Len()
	Add(idx int, t T) error

	// Set the element given idx, err if out of range
	Set(idx int, t T)

	// Delete the element given idx, err if out of range
	Delete(idx int) error

	// Len Length of
	Len() int

	// Cap Capacity of
	Cap() int

	// Range iteration
	Range(fn func(idx int, t T) error) error

	// ToSlice transform List to a slice, empty []T if no element
	ToSlice() []T

	// Generator yield-like
	Generator() <-chan T
}
```

### ArrayList

Implmented by Go built-in slice.

Note: **Add/Delete/ToSlice** actually generate/return a new slice, be careful in loop call.

| func                                       | Time Complexity |
| ------------------------------------------ | --------------- |
| `Get(idx int) T`                           | O(1)            |
| `Append(ts ...T)`                          | O(1)            |
| `Add(idx int, t T) error`                  | O(n)            |
| `Set(idx int, t T)`                        | O(1)            |
| `Delete(idx int) error`                    | O(n)            |
| `Len() int`                                | O(1)            |
| `Cap() int`                                | O(1)            |
| `Range(fn func(idx int, t T) error) error` | O(n)            |
| `ToSlice() []T`                            | O(1)            |
| `Generator() <-chan T`                     | O(n)            |

> func NewArrayList

```go
func NewArrayList[T any](cap int) *ArrayList[T]
```

Constructor given capacity `cap`.

Example:

```go
l := list.NewArrayList[int](5)
```

> func NewArrayListFrom

```go
func NewArrayListFrom[T any](sl []T) *ArrayList[T]
```

Constructor given a slice `sl`.

Example:

```go
l := list.NewArrayListFrom[int]([]int{1, 2, 3, 4, 5})
```

### LinkedList

A **circular** doubly linked list.

Note: head/tail as **guard** node.

```go
// head(-1) ←→ 0 ←→ ... ←→ n-1 ←→ tail(n)
//  ↑↓                             ↑↓
//  |_______________________________|
```

| func                                       | Time Complexity |
| ------------------------------------------ | --------------- |
| `Get(idx int) T`                           | O(n)            |
| `Append(ts ...T)`                          | O(1)            |
| `Add(idx int, t T) error`                  | O(n)            |
| `Set(idx int, t T)`                        | O(n)            |
| `Delete(idx int) error`                    | O(n)            |
| `Len() int`                                | O(1)            |
| `Cap() int`                                | O(1)            |
| `Range(fn func(idx int, t T) error) error` | O(n)            |
| `ToSlice() []T`                            | O(n)            |
| `Generator() <-chan T`                     | O(n)            |

> func NewLinkedList

```go
func NewLinkedList[T any]() *LinkedList[T]
```

Constructor.

Example:

```go
l := list.NewLinkedList[int]()
```

> func NewLinkedListFrom

```go
func NewLinkedListFrom[T any](sl []T) *LinkedList[T]
```

Constructor given a slice `sl`.

Example:

```go
l := list.NewLinkedListFrom[int]([]int{1, 2, 3, 4, 5})
```