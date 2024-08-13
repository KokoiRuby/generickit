## Syncx

stdlib [sync](https://pkg.go.dev/sync#pkg-overview) provides basic synchronization primitives.

While syncx encapsulates with [generic](https://go.dev/blog/intro-generics).

### Pool

stblib [Pool](https://pkg.go.dev/sync#Pool) is a set of temporary objects that may be individually saved and retrieved.

This pool is encapsulated with [generic](https://go.dev/blog/intro-generics). (Avoid type assertion when dealing with such as`sync.Pool.Get()`)

> func NewPool

```go
func NewPool[T any](factory func() T) *Pool[T]
```

Constructor passed by a `factory` function (to produce obj to be stored in pool).

Example:

```go
pool = syncx.NewPool[[]byte](func() []byte {
	sl := make([]byte, 1, 3)
	sl[0] = 'a'
	return sl
})

fmt.Print(string(suite.pool.Get()))
// Output:
// a
```

### Map

stdlib [Map](https://pkg.go.dev/sync#Map) is like a Go map[any]any but is safe for concurrent use by multiple goroutines without additional locking or coordination. Loads, stores, and deletes run in amortized constant time.

This Map is encapsulated with [generic](https://go.dev/blog/intro-generics). (Avoid type assert when dealing with such as `sync.Map.Load()`)

> func NewMap

```go
func NewMap[K comparable, V any]() *Map[K, V]
```

Constructor.

Example:

```go
func ExampleMap_Load() {
	var m syncx.Map[string, int]
	m.Store("a", 1)
	val, ok := m.Load("a")
	if ok {
		fmt.Println(val)
	}
	// Output:
	// 123
}
```

### AtomicValue

A `Value` provides an atomic load and store of a consistently typed value. The zero value for a Value returns nil from Load. Once Store has been called, a Value must not be copied.

This Value is encapsulated with [generic](https://go.dev/blog/intro-generics). (Avoid type assert when dealing with such as `sync.atomic.Value.Load()`)

> func NewValue

```go
func NewValue[T any](t T) *Value[T]
```

Constructor.

Example:

```go
func ExampleMap_Load() {
	var val syncx.Value[int]
	val.Store(1)
    fmt.Print(val.Load())
	// Output:
	// 1
}
```

