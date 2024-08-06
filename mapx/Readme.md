## Map

[map](https://go.dev/blog/maps) is Golang built-in hash table. stdlib [maps](https://pkg.go.dev/maps) also defines various functions useful with maps of any type.

**Before getting started, u need to import these packages in src.**

```go
import (
    "github.com/KokoiRuby/generickit/mapx"
    "github.com/KokoiRuby/generickit/syncx"
)
```

map interface:

```go
type mapi[K any, V any] interface {
	Put(key K, val V) error
	Get(key K) (V, bool)
	Delete(k K) (V, bool)
	Keys() []K
	Values() []V
	Len() uint64
}
```

### Auxiliary

> func Keys

```go
func Keys[K comparable, V any](m map[K]V) []K 
```

Returning a slice contains all keys of `m`. No order guaranteed.

Example:

```go
m = make(map[string]int, 3)
m["a"] = 1
m["b"] = 2
m["c"] = 3

kSl := mapx.Keys[string, int](m) // could be [b c a] or [c a b]
```

> func Values

```go
func Values[K comparable, V any](m map[K]V) []V
```

Returning a slice contains all values of `m`. No order guaranteed.

Example:

```go
m = make(map[string]int, 3)
m["a"] = 1
m["b"] = 2
m["c"] = 3

vSl := mapx.Values[string, int](m) // could be [2 3 1] or [1 3 2]
```

> func KeysValues

```go
func KeysValues[K comparable, V any](m map[K]V) ([]K, []V)
```

Returning a slice contains all keys & a slice contains all values of `m`. No order guaranteed.

Example:

```go
m = make(map[string]int, 3)
m["a"] = 1
m["b"] = 2
m["c"] = 3

kSl, vSl := mapx.KeysValues[string, int](m)
```

> func ToMap

```go
func ToMap[K comparable, V any](kSl []K, vSl []V) (map[K]V, error)
```

Example:

```go
kSl = []string{"a", "b", "c"}
vSl = []int{1, 2, 3}

m, _ := mapx.ToMap[string, int](suite.kSl, suite.vSl) // map[a:1 b:2 c:3]
```

### SyncMap

A simple conccurent map using **decorator**.

**U should always firtly consider using stdlib [sync.Map](https://pkg.go.dev/sync#Map) which provides better performance.**

> func NewSyncMap

```go
func NewSyncMap[K comparable, V any]() *SyncMap[K, V]
```

Constructor.

Example:

```go
m := mapx.NewSyncMap[string, int]
```

### HashMap

Go built-in map where key only supports comparable.

**This HashMap supports complex data structure like Go struct as key.**

Implementation: built-in map + singly linkedlist.

```go
// hashVal → bucket[0] → { &node{k, v} as root → &node{k, v} → &node{k, v}... } 
//         → bucket[1] → { &node{k, v} as root → &node{k, v} → &node{k, v}... } 
//         → ...
//         → bucket[n] → { &node{k, v} as root → &node{k, v} → &node{k, v}... } 
```

> func NewHashMap

```go
func NewHashMap[K HashAble, V any](size uint) *HashMap[K, V]
```

Consturctor with size.

Example:

```go
// HashAble shall be implemented by Key
type HashAble interface {
	Hash() uint64
	Equal(k any) bool
}

type User struct {
	id   uint
	name string
	age  uint8
}

func (u User) Hash() uint64 {
	h := fnv.New64a()
	_, err := h.Write([]byte(fmt.Sprintf("%d", u.id)))
	if err != nil {
		return 0
	}
	return h.Sum64()
}

func (u User) Equal(k any) bool {
	// type assertion
	if ku, ok := k.(User); ok {
		if ku.id == u.id && ku.name == u.name && ku.age == u.age {
			return true
		}
	}
	return false
}
```

```go
m = mapx.NewHashMap[User, int](5)
m.Put(User{id: 0, name: "Alice", age: 10}, 0)
v, isFound := m.Get(User{id: 0, name: "Alice", age: 10})
fmt.Println(v)
fmt.Println(true)
v, isFound = m.Get(User{id: 0, name: "Alice", age: 100})
fmt.Println(v)
fmt.Println(true)
// output:
// 0
// true
// 0
// false
```

### LinkedMap

### MultiMap

