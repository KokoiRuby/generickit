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

package syncx

import "sync/atomic"

type Value[T any] struct {
	val atomic.Value
}

func NewValue[T any](t T) *Value[T] {
	at := atomic.Value{}
	at.Store(t)
	return &Value[T]{
		val: at,
	}
}

func (v *Value[T]) Load() T {
	return v.val.Load().(T)
}

func (v *Value[T]) Store(t T) {
	v.val.Store(t)
}

func (v *Value[T]) Swap(new T) (old T) {
	return v.val.Swap(new).(T)
}

func (v *Value[T]) CompareAndSwap(old, new any) (swapped bool) {
	return v.val.CompareAndSwap(old, new)
}
