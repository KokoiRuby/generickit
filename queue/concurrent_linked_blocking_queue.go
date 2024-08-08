/*
 * Copyright 2024 KoKoiRuby
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package queue

import (
	"context"
	"github.com/KokoiRuby/generickit/list"
	"sync"
)

// ConcurrentLinkedBlockingQueue bounded if capacity > 0, otherwise un-bounded = enqueue always succeed
type ConcurrentLinkedBlockingQueue[T any] struct {
	lock     *sync.RWMutex
	capacity int
	list     *list.LinkedList[T]

	enqueueCond *sync.Cond
	dequeueCond *sync.Cond
}

func NewConcurrentLinkedBlockingQueue[T any](capacity int) *ConcurrentLinkedBlockingQueue[T] {
	mutex := &sync.RWMutex{}
	return &ConcurrentLinkedBlockingQueue[T]{
		lock:        mutex,
		capacity:    capacity,
		list:        list.NewLinkedList[T](),
		enqueueCond: sync.NewCond(mutex),
		dequeueCond: sync.NewCond(mutex.RLocker()),
	}
}

func (q *ConcurrentLinkedBlockingQueue[T]) Enqueue(ctx context.Context, t T) error {
	if ctx.Err() != nil {
		return ctx.Err()
	}
	q.lock.Lock()
	defer q.lock.Unlock()
	// queue is full
	for q.capacity > 0 && q.list.Len() == q.capacity {
		select {
		// return if ctx timeout or canceled, otherwise wait to enqueue
		case <-ctx.Done():
			return ctx.Err()
		default:
			q.enqueueCond.Wait()
		}
	}
	q.list.Append(t)
	// notify to dequeue
	q.dequeueCond.Broadcast()
	return nil
}

func (q *ConcurrentLinkedBlockingQueue[T]) Dequeue(ctx context.Context) (T, error) {
	if ctx.Err() != nil {
		var t T
		return t, ctx.Err()
	}
	q.lock.Lock()
	defer q.lock.Unlock()
	// queue is empty
	for q.list.Len() == 0 {
		select {
		case <-ctx.Done():
			var t T
			return t, ctx.Err()
		default:
			q.dequeueCond.Wait()
		}
	}
	v, err := q.list.Delete(0)
	q.enqueueCond.Broadcast()
	return v, err
}

func (q *ConcurrentLinkedBlockingQueue[T]) Len() int {
	q.lock.RLock()
	defer q.lock.RUnlock()
	return q.list.Len()
}

func (q *ConcurrentLinkedBlockingQueue[T]) ToSlice() []T {
	q.lock.RLock()
	defer q.lock.RUnlock()
	return q.list.ToSlice()
}
