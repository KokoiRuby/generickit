## Bean

"Bean" typically refers to a class/model that follows certain conventions.

- **Properties**: private instance variables with public getter & setter to access & modify.
- **Serializable**: supports serialization, can be converted to a stream of bytes.
- **Constructor**: for instantiation.

Golang usually uses **struct** to represent a Bean.

**Before getting started, u need to import these packages in src.**

```go
import (
    "github.com/KokoiRuby/generickit/bean/option"
)
```

### Option

A Builder-like design pattern provides an elegant, extensible & human-readble way to construct a complicated object with options that are not all necessary.

> func Apply

```go
func Apply[T any](t *T, opts ...Option[T]) error
```

`Option` is a type redef of a setter function of struct `T`.

```go
type Option[T any] func(t *T) error
```

Example:

```go
func ExampleApply() {
	srv, _ := option.NewServer("127.0.0.1:8080", 8080)
	fmt.Printf("%+v\n", srv)
	srv, _ = option.NewServer("127.0.0.1:8080", 8080, WithProtocol("udp"), WithTimeout(10*time.Second), WithMaxConn(50))
	fmt.Printf("%+v\n", srv)
	// output:
	// &{Addr:127.0.0.1:8080 Port:8080 Protocol:tcp Timeout:5s MaxConn:100 TLS:<nil>}
	// &{Addr:127.0.0.1:8080 Port:8080 Protocol:udp Timeout:10s MaxConn:50 TLS:<nil>}
}
```

### Copier