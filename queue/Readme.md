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

### Concurrent Linked Blocking Queue

Leverage Go built-in sync primitive [sync.RWMutex](https://pkg.go.dev/sync#RWMutex) & [sync.Cond](https://pkg.go.dev/sync#Cond) to secure concurrency.

> func NewConcurrentLinkedBlockingQueue

```go
func NewConcurrentLinkedBlockingQueue[T any](capacity int) *ConcurrentLinkedBlockingQueue[T]
```

Constructor. Whether queue is bound is controlled by `capacity` where negative val indicates unbound.

Examples:

```go
func ExampleConcurrentLinkedBlockingQueue() {
	q := queue.NewConcurrentLinkedBlockingQueue[int](3)
	_ = q.Enqueue(1)
	val, err := q.Dequeue()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(val)
	// Output:
	// 1
}
```



