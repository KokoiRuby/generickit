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

import "sync"

type Pool[T any] struct {
	p sync.Pool
}

func NewPool[T any](factory func() T) *Pool[T] {
	return &Pool[T]{
		p: sync.Pool{
			New: func() any {
				return factory()
			},
		},
	}
}

func (p *Pool[T]) Put(v T) {
	p.p.Put(v)
}

func (p *Pool[T]) Get() T {
	return p.p.Get().(T)
}
