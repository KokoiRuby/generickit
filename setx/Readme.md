## Set

Golang does not provide built-in set.

The community also provides [golang-set](https://github.com/deckarep/golang-set): thread-safe and non-thread-safe high-performance sets for Go.

### mapSet

A thread-safe set implemented by built-in map (value is empty struct) & RW mutex.

Note: elements must be comparable

> func NewMapSet

```go
func NewMapSet[T comparable]() *MapSet[T]
```

Constructor.

Example:

```go
ms = setx.NewMapSet[int]()
```

> func Add

```go
func (s *MapSet[T]) Add(key T)
```

Add a element into set. This function is `O(1)`.

Example:

```go
ms = setx.NewMapSet[int]()
ms.Add(1)
```

> func Remove

```go
func (s *MapSet[T]) Remove(key T)
```

Remove a element from set. This function is `O(1)`.

Example:

```go
ms = setx.NewMapSet[int]()
ms.Add(1)
ms.Remove(1)
```

> func Contains

```go
func (s *MapSet[T]) Contains(key T) bool
```

Check if set contains an element. This function is `O(1)`.

Example:

```go
ms = setx.NewMapSet[int]()
ms.Add(1)
ms.Contains(1) // true
```

> func Size

```go
func (s *MapSet[T]) Size() int
```

Return the number of elements in set. This function is `O(1)`.

Example:

```go
ms = setx.NewMapSet[int]()
ms.Add(1)
ms.Size() // 1
```

> func Elements

```go
func (s *MapSet[T]) Elements() []T {
```

Return a slice contains all elements in set. This function is `O(len(MapSet.m))`.

Example:

```go
ms = setx.NewMapSet[int]()
ms.Add(1)
ms.Elements // [1]
```

> func Clear

Clear the set. This function is `O(1)`.

```go
func (s *MapSet[T]) Clear()
```

Example:

```go
ms = setx.NewMapSet[int]()
ms.Add(1)
ms.Clear() // []
```

