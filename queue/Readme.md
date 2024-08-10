## Queue

A queue is a linear structure that follows **FIFO** order in which the operations are performed.

Below queues are provided:

- Concurrent [Blocking] Linked Queue
- ...

**Before getting started, u need to import these packages in src.**

```go
import (
    "github.com/KokoiRuby/generickit/queue"
    "github.com/KokoiRuby/generickit/list"
)
```

### Concurrent Linked Queue

This queue is unbound, using [sync.atomic](https://pkg.go.dev/sync/atomic) & CAS operation to secure concurrency.

```go
// head(val + next) → head.next(val + next) → ...  → tail (val + next)     Node Layer
//
//
// &head                  &(head.next)                    &(tail)          Pointer Layer
//
//   ↑                                                       ↑
//
// q.head                                                  q.tail          ConcurrentLinkedQueue
```

> func NewConcurrentLinkedQueue

```go
func NewConcurrentLinkedQueue[T any]() *ConcurrentLinkedQueue[T]
```

Constructor.

Example:

```go
func ExampleConcurrentLinkedQueue() {
    q := queue.NewConcurrentLinkedQueue[int]()
    _ = q.Enqueue(1)
    val, _ := q.Dequeue()
    fmt.Println(val)
    // output:
    // 1
}
```

### Concurrent Linked Blocking Queue

Leverage Go built-in sync primitive [sync.RWMutex](https://pkg.go.dev/sync#RWMutex) & [sync.Cond](https://pkg.go.dev/sync#Cond) to secure concurrency & blocking.

Leverage Go built-in [context](https://pkg.go.dev/context) to achieve timeout control.

> func NewConcurrentLinkedBlockingQueue

```go
func NewConcurrentLinkedBlockingQueue[T any](capacity int) *ConcurrentLinkedBlockingQueue[T]
```

Constructor. Whether queue is bound is controlled by `capacity` where negative val indicates unbound.

Example:

```go
func ExampleConcurrentLinkedBlockingQueue() {
	q := queue.NewConcurrentLinkedBlockingQueue[int](3)
    ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
    
	_ = q.Enqueue(1)
	val, err := q.Dequeue()
    switch err {
	case context.Canceled:
        // cancel() called
	case context.DeadlineExceeded:
		// timeout
	case nil:
		fmt.Println(val)
	default:
		// TODO
	}
	// Output:
	// 1
}
```



