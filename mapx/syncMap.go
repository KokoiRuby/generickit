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

package mapx

import "sync"

// SyncMap decorate built-in map with RW Lock
type SyncMap[K comparable, V any] struct {
	lock sync.RWMutex
	m    map[K]V
}

func NewSyncMap[K comparable, V any]() *SyncMap[K, V] {
	return &SyncMap[K, V]{m: make(map[K]V)}
}

func (m *SyncMap[K, V]) Put(key K, val V) error {
	m.lock.Lock()
	defer m.lock.Unlock()
	m.m[key] = val
	return nil
}

func (m *SyncMap[K, V]) Get(key K) (V, bool) {
	m.lock.RLock()
	defer m.lock.RUnlock()
	val, ok := m.m[key]
	return val, ok
}

func (m *SyncMap[K, V]) Delete(k K) (V, bool) {
	m.lock.Lock()
	defer m.lock.Unlock()
	val, ok := m.m[k]
	delete(m.m, k)
	return val, ok
}

func (m *SyncMap[K, V]) Keys() []K {
	return Keys[K, V](m.m)
}

func (m *SyncMap[K, V]) Values() []V {
	return Values[K, V](m.m)
}

func (m *SyncMap[K, V]) Len() uint64 {
	return uint64(len(m.m))
}
