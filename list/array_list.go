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
	"errors"
	"github.com/KokoiRuby/generickit/slice"
)

type ArrayList[T any] struct {
	vals []T
}

func NewArrayList[T any](cap int) *ArrayList[T] {
	return &ArrayList[T]{vals: make([]T, 0, cap)}
}

func NewArrayListFrom[T any](sl []T) *ArrayList[T] {
	return &ArrayList[T]{vals: sl}
}

//func (l *ArrayList[T]) GetVals() []T {
//	return l.vals
//}

func (l *ArrayList[T]) Get(idx int) (T, error) {
	var v T
	if l.IsOutOfRange(idx) {
		return v, errors.New("index out of range")
	}
	return l.vals[idx], nil
}

func (l *ArrayList[T]) IsOutOfRange(idx int) bool {
	return idx <= 0 || idx >= l.Len()
}

func (l *ArrayList[T]) Append(ts ...T) {
	l.vals = append(l.vals, ts...)
}

// Add proxy to slice.Insert
func (l *ArrayList[T]) Add(idx int, t T) (err error) {
	l.vals, err = slice.Insert(l.vals, t, idx)
	return err
}

func (l *ArrayList[T]) Set(idx int, t T) error {
	if l.IsOutOfRange(idx) {
		return errors.New("index out of range")
	}
	l.vals[idx] = t
	return nil
}

// Delete proxy to slice.Delete & slice.Shrink
func (l *ArrayList[T]) Delete(idx int) (err error) {
	l.vals, err = slice.Delete(l.vals, idx)
	l.vals = slice.Shrink(l.vals)
	return err
}

func (l *ArrayList[T]) Len() int {
	return len(l.vals)
}

func (l *ArrayList[T]) Cap() int {
	return cap(l.vals)
}

func (l *ArrayList[T]) Range(fn func(idx int, t T) error) error {
	for i, v := range l.vals {
		if err := fn(i, v); err != nil {
			return err
		}
	}
	return nil
}

func (l *ArrayList[T]) ToSlice() []T {
	tmp := make([]T, len(l.vals))
	copy(tmp, l.vals)
	return tmp
}

// Generator proxy to slice.Generator
func (l *ArrayList[T]) Generator() <-chan T {
	return slice.Generator(l.vals)
}
