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

import (
	"github.com/KokoiRuby/generickit/syncx"
)

// HashAble shall be implemented by Key
type HashAble interface {
	Hash() uint64
	Equal(k any) bool
}

type node[K HashAble, V any] struct {
	key   K
	value V
	next  *node[K, V] // singly
}

func (n *node[K, V]) clear() {
	var k K
	var v V
	n.key = k
	n.value = v
	n.next = nil
}

func (m *HashMap[K, V]) newNode(k K, val V) *node[K, V] {
	newNode := m.pool.Get()
	newNode.key = k
	newNode.value = val
	return newNode
}

type HashMap[K HashAble, V any] struct {
	hashmap map[uint64]*node[K, V]
	pool    *syncx.Pool[*node[K, V]] // for re-usage, K usually is a struct, save cost of constructing.
}

func NewHashMap[K HashAble, V any](size uint) *HashMap[K, V] {
	return &HashMap[K, V]{
		hashmap: make(map[uint64]*node[K, V], size),
		// pool with factory
		pool: syncx.NewPool[*node[K, V]](func() *node[K, V] {
			return &node[K, V]{}
		}),
	}
}

func (m *HashMap[K, V]) Put(key K, val V) error {
	hash := key.Hash()
	root, ok := m.hashmap[hash]
	// bucket not taken up yet
	if !ok {
		m.hashmap[hash] = m.newNode(key, val)
		return nil
	}
	// bucket taken then iter singly
	pre := root
	for root != nil {
		if root.key.Equal(key) {
			root.value = val
			return nil
		}
		// step forward
		pre = root
		root = root.next
	}
	// no key is matched, append
	pre.next = m.newNode(key, val)
	return nil
}

func (m *HashMap[K, V]) Get(key K) (V, bool) {
	hash := key.Hash()
	var val V
	if root, ok := m.hashmap[hash]; ok {
		// bucket taken then iter singly
		for root != nil {
			if root.key.Equal(key) {
				return root.value, true
			}
			// step forward
			root = root.next
		}
	} else {
		// bucket not taken yet
		return val, false
	}
	return val, false
}

func (m *HashMap[K, V]) Delete(k K) (V, bool) {
	root, ok := m.hashmap[k.Hash()]
	var v V
	// bucket not taken
	if !ok {
		return v, false
	}
	// bucket taken then iter singly、
	pre := root
	idx := 0
	for root != nil {
		if root.key.Equal(k) {
			if idx == 0 && root.next == nil {
				// only one node in bucket
				v = root.value
				delete(m.hashmap, k.Hash())
				return v, true
			} else if idx == 0 && root.next != nil {
				// 2 nodes in bucket
				m.hashmap[k.Hash()] = root.next
			} else {
				pre.next = root.next
			}
			v = root.value
			root.clear()
			m.pool.Put(root) // node put back to
			return v, true
		}
		idx++
		pre = root
		root = root.next
	}
	return v, false
}

func (m *HashMap[K, V]) Keys() []K {
	// not able to assume cap
	sl := make([]K, 0)
	for _, bucket := range m.hashmap {
		cur := bucket
		for cur != nil {
			sl = append(sl, cur.key)
			cur = cur.next
		}
	}
	return sl
}

func (m *HashMap[K, V]) Values() []V {
	// not able to assume cap
	sl := make([]V, 0)
	for _, bucket := range m.hashmap {
		cur := bucket
		for cur != nil {
			sl = append(sl, cur.value)
			cur = cur.next
		}
	}
	return sl
}

// Len # of buckets
func (m *HashMap[K, V]) Len() uint64 {
	return uint64(len(m.hashmap))
}
