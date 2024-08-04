## Syncx

stdlib [sync](https://pkg.go.dev/sync#pkg-overview) provides basic synchronization primitives.

While syncx encapsulates with [generic](https://go.dev/blog/intro-generics).

### Pool

stblib [Pool](https://pkg.go.dev/sync#Pool) is a set of temporary objects that may be individually saved and retrieved.

This pool is encapsulated with [generic](https://go.dev/blog/intro-generics). (Avoid type assert when dealing with `sync.Pool.Get()`)

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

