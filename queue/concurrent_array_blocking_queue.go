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
	"sync"
)

type ConcurrentArrayBlockingQueue[T any] struct {
	data        []T
	lock        *sync.RWMutex
	head, tail  int
	size        int
	enqueueCond *sync.Cond
	dequeueCond *sync.Cond
	zero        T
}

func NewConcurrentArrayBlockingQueue[T any](capacity int) *ConcurrentArrayBlockingQueue[T] {
	mutex := &sync.RWMutex{}
	return &ConcurrentArrayBlockingQueue[T]{
		data:        make([]T, capacity),
		lock:        mutex,
		enqueueCond: sync.NewCond(mutex),
		dequeueCond: sync.NewCond(mutex),
	}
}

func (q *ConcurrentArrayBlockingQueue[T]) Enqueue(ctx context.Context, t T) error {
	if ctx.Err() != nil {
		return ctx.Err()
	}

	q.lock.Lock()
	defer q.lock.Unlock()

	// queue is full
	for q.size == cap(q.data) {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
			q.enqueueCond.Wait()
		}
	}
	q.data[q.tail] = t
	q.tail = (q.tail + 1) % cap(q.data)
	q.size++
	q.dequeueCond.Broadcast()
	return nil
}

func (q *ConcurrentArrayBlockingQueue[T]) Dequeue(ctx context.Context) (T, error) {
	if ctx.Err() != nil {
		var t T
		return t, ctx.Err()
	}

	q.lock.Lock()
	defer q.lock.Unlock()

	// queue is empty
	for q.size == 0 {
		select {
		case <-ctx.Done():
			var t T
			return t, ctx.Err()
		default:
			q.dequeueCond.Wait()
		}
	}

	t := q.data[q.head]
	// clear
	q.data[q.head] = q.zero
	q.head = (q.head + 1) % cap(q.data)
	q.size--
	q.enqueueCond.Broadcast()
	return t, nil
}

func (q *ConcurrentArrayBlockingQueue[T]) Len() int {
	q.lock.RLock()
	defer q.lock.RUnlock()
	return q.size
}

func (q *ConcurrentArrayBlockingQueue[T]) ToSlice() []T {
	q.lock.RLock()
	defer q.lock.RUnlock()
	sl := make([]T, 0, q.size)
	step := 0
	capacity := cap(q.data)
	for step < q.size {
		idx := (q.head + step) % capacity
		sl = append(sl, q.data[idx])
		step++
	}
	return sl
}
