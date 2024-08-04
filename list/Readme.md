## List

list represents a collection of elements, where each is connected to the next through references or pointers.

stdlib container/[list](https://pkg.go.dev/container/list) implements a doubly linked list.

**Before getting started, u need to import these packages in src.**

```go
import (
    "github.com/KokoiRuby/generickit/list"
)
```

List Interface

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

	// ToSlice transform List to a slice, []T if no element
	ToSlice() []T

	// Generator yield-like
	Generator() <-chan T
}
```

### ArrayList



### LinkedList



### Doubly LinkedList



### Circular Doubly LinkedList



### Concurrent List