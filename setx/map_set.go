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

package setx

import (
	"sync"
)

type MapSet[T comparable] struct {
	m       map[T]struct{}
	rwMutex sync.RWMutex
}

func NewMapSet[T comparable]() *MapSet[T] {
	return &MapSet[T]{m: make(map[T]struct{}), rwMutex: sync.RWMutex{}}
}

func (s *MapSet[T]) Add(key T) {
	s.rwMutex.Lock()
	defer s.rwMutex.Unlock()
	s.m[key] = struct{}{}
}

func (s *MapSet[T]) Remove(key T) {
	s.rwMutex.Lock()
	defer s.rwMutex.Unlock()
	delete(s.m, key)
}

func (s *MapSet[T]) Contains(key T) bool {
	s.rwMutex.RLock()
	defer s.rwMutex.RUnlock()

	_, ok := s.m[key]
	return ok
}

func (s *MapSet[T]) Size() int {
	s.rwMutex.RLock()
	defer s.rwMutex.RUnlock()
	return len(s.m)
}

func (s *MapSet[T]) Elements() []T {
	s.rwMutex.RLock()
	defer s.rwMutex.RUnlock()
	elements := make([]T, 0, len(s.m))
	for key := range s.m {
		elements = append(elements, key)
	}
	return elements
}

func (s *MapSet[T]) Clear() {
	s.rwMutex.Lock()
	defer s.rwMutex.Unlock()
	s.m = make(map[T]struct{})
}
