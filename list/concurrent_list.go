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

package list

import (
	"sync"
)

// ConcurrentList by decorator
type ConcurrentList[T any] struct {
	List[T]
	lock sync.RWMutex
}

func NewConcurrentList[T any](cap int) *ConcurrentList[T] {
	return &ConcurrentList[T]{List: NewArrayList[T](cap)}
}

func NewConcurrentListFrom[T any](sl []T) *ConcurrentList[T] {
	return &ConcurrentList[T]{List: NewArrayListFrom[T](sl)}
}

/*
methods proxy to List[T]
*/

func (l *ConcurrentList[T]) Get(idx int) T {
	l.lock.RLock()
	defer l.lock.RUnlock()
	return l.List.Get(idx)
}

func (l *ConcurrentList[T]) Append(ts ...T) {
	l.lock.Lock()
	defer l.lock.Unlock()
	l.List.Append(ts...)
}

// Add proxy to slice.Insert
func (l *ConcurrentList[T]) Add(idx int, t T) (err error) {
	l.lock.Lock()
	defer l.lock.Unlock()
	err = l.List.Add(idx, t)
	if err != nil {
		return err
	}
	return nil
}

func (l *ConcurrentList[T]) Set(idx int, t T) {
	l.lock.Lock()
	defer l.lock.Unlock()
	l.List.Set(idx, t)
}

// Delete proxy to slice.Delete & slice.Shrink
func (l *ConcurrentList[T]) Delete(idx int) (err error) {
	l.lock.Lock()
	defer l.lock.Unlock()
	err = l.List.Delete(idx)
	return err
}

func (l *ConcurrentList[T]) Len() int {
	l.lock.RLock()
	defer l.lock.RUnlock()
	return l.List.Len()
}

func (l *ConcurrentList[T]) Cap() int {
	l.lock.RLock()
	defer l.lock.RUnlock()
	return l.List.Cap()
}

func (l *ConcurrentList[T]) Range(fn func(idx int, t T) error) error {
	l.lock.RLock()
	defer l.lock.RUnlock()
	err := l.List.Range(fn)
	if err != nil {
		return err
	}
	return nil
}

func (l *ConcurrentList[T]) ToSlice() []T {
	l.lock.RLock()
	defer l.lock.RUnlock()
	return l.List.ToSlice()
}

// Generator proxy to slice.Generator
func (l *ConcurrentList[T]) Generator() <-chan T {
	l.lock.RLock()
	defer l.lock.RUnlock()
	return l.List.Generator()
}
