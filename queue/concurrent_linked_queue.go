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
	"errors"
	"sync/atomic"
	"unsafe"
)

type ConcurrentLinkedQueue[T any] struct {
	head unsafe.Pointer
	tail unsafe.Pointer
}

func NewConcurrentLinkedQueue[T any]() *ConcurrentLinkedQueue[T] {
	head := &node[T]{}
	ptr := unsafe.Pointer(head)
	return &ConcurrentLinkedQueue[T]{head: ptr, tail: ptr}
}

func (q *ConcurrentLinkedQueue[T]) Enqueue(value T) error {
	newNode := &node[T]{val: value}
	newPtr := unsafe.Pointer(newNode)

	// for-loop + atomic → spin lock (cpu-costly), return/break if cond is met
	for {
		// get pointer of queue tail
		tailPtr := atomic.LoadPointer(&q.tail)
		// type conversion from unsafe.Pointer to node struct
		tail := (*node[T])(tailPtr)
		// get pointer of queue tail next
		tailNextPtr := atomic.LoadPointer(&tail.next)
		// there are other ongoing Enqueue
		if tailNextPtr != nil {
			continue
		}
		// I got this!
		// tail points to new node
		if atomic.CompareAndSwapPointer(&tail.next, tailNextPtr, newPtr) {
			// q.tail points to new node pointer
			atomic.CompareAndSwapPointer(&q.tail, tailPtr, newPtr)
			return nil
		}
	}
}

func (q *ConcurrentLinkedQueue[T]) Dequeue() (T, error) {
	// for-loop + atomic → spin lock (cpu-costly), return/break if cond is met
	for {
		// get head/tail pointer & node
		headPtr := atomic.LoadPointer(&q.head)
		head := (*node[T])(headPtr)
		tailPtr := atomic.LoadPointer(&q.tail)
		tail := (*node[T])(tailPtr)

		if head == tail {
			var t T
			return t, errors.New("empty queue")
		}

		// q.head points to head.next
		headNextPtr := atomic.LoadPointer(&head.next)
		// I got this = no other ongoing Dequeue
		if atomic.CompareAndSwapPointer(&q.head, headPtr, headNextPtr) {
			headNext := (*node[T])(headNextPtr)
			return headNext.val, nil
		}
	}
}

func (q *ConcurrentLinkedQueue[T]) ToSlice() []T {
	var sl []T
	cur := (*node[T])((*node[T])(q.head).next)
	for cur != nil {
		sl = append(sl, cur.val)
		cur = (*node[T])(cur.next)
	}
	return sl
}
