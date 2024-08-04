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

import "errors"

type node[T any] struct {
	next, prev *node[T]
	val        T
}

type LinkedList[T any] struct {
	head, tail *node[T]
	length     int
}

func NewLinkedList[T any]() *LinkedList[T] {
	head := &node[T]{}
	tail := &node[T]{
		prev: head,
		next: head,
	}
	head.prev = tail
	head.next = tail
	return &LinkedList[T]{
		head: head,
		tail: tail,
	}
}

func NewLinkedListFrom[T any](sl []T) *LinkedList[T] {
	list := NewLinkedList[T]()
	list.Append(sl...)
	return list
}

func (l *LinkedList[T]) find(idx int) *node[T] {
	var cur *node[T]
	if idx < l.Len()/2 {
		cur = l.head.next
		for i := 0; i < idx; i++ {
			cur = cur.next
		}
	} else {
		cur = l.tail
		for i := l.length; i > idx; i-- {
			cur = cur.prev
		}
	}
	return cur
}

// FindByVal requires T is comparable
//func (l *LinkedList[T]) FindByVal(val T) *node[T] {
//	var cur *node[T]
//	for cur = l.head; cur != l.tail; cur = cur.next {
//		if cur.GetVal() == val {
//			return cur
//		}
//	}
//	return nil
//}

func (l *LinkedList[T]) Get(idx int) (T, error) {
	if l.IsOutOfRange(idx) {
		var zeroVal T
		return zeroVal, errors.New("index out of range")
	}
	return l.find(idx).val, nil
}

func (l *LinkedList[T]) IsOutOfRange(idx int) bool {
	return idx < 0 || idx >= l.Len()
}

func (l *LinkedList[T]) Append(vals ...T) {
	for _, val := range vals {
		node := &node[T]{
			prev: l.tail.prev,
			next: l.tail,
			val:  val,
		}
		node.prev.next, node.next.prev = node, node
		l.length++
	}
}

func (l *LinkedList[T]) Add(idx int, val T) error {
	if l.IsOutOfRange(idx) {
		return errors.New("index out of range")
	}
	if idx == l.length {
		l.Append(val)
	}
	nxt := l.find(idx)
	toAdd := &node[T]{
		prev: nxt.prev,
		next: nxt,
		val:  val,
	}
	toAdd.prev.next, toAdd.next.prev = toAdd, toAdd
	l.length++
	return nil
}

func (l *LinkedList[T]) Set(idx int, val T) error {
	if l.IsOutOfRange(idx) {
		return errors.New("index out of range")
	}
	toSet := l.find(idx)
	toSet.val = val
	return nil
}

func (l *LinkedList[T]) Delete(idx int) error {
	if l.IsOutOfRange(idx) {
		return errors.New("index out of range")
	}
	toDelete := l.find(idx)
	toDelete.prev.next = toDelete.next
	toDelete.next.prev = toDelete.prev
	toDelete.next, toDelete.prev = nil, nil
	l.length--
	return nil
}

func (l *LinkedList[T]) Len() int {
	return l.length
}

func (l *LinkedList[T]) Cap() int {
	return l.Len()
}

func (l *LinkedList[T]) Range(f func(idx int, val T) error) error {
	for cur, i := l.head.next, 0; i < l.Len(); i++ {
		err := f(i, cur.val)
		if err != nil {
			return err
		}
		cur = cur.next
	}
	return nil
}

func (l *LinkedList[T]) ToSlice() []T {
	sl := make([]T, l.length)
	for cur, i := l.head.next, 0; i < l.Len(); i++ {
		sl[i] = cur.val
		cur = cur.next
	}
	return sl
}

func (l *LinkedList[T]) Generator() <-chan T {
	ch := make(chan T)
	go func() {
		defer close(ch)
		for cur, i := l.head.next, 0; i < l.Len(); i++ {
			ch <- cur.val
			cur = cur.next
		}
	}()
	return ch
}

/*
Extra
*/
