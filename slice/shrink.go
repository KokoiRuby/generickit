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

package slice

func calNewCapacity(length, capacity int) (int, bool) {
	if length >= 256 && float32(capacity)/float32(length) > 1.25 {
		return int(float32(capacity) / 1.25), true
	}
	if length < 256 && float32(capacity)/float32(length) > 2 {
		return int(float32(capacity) / 2.0), true
	}
	return capacity, false
}

func Shrink[T any](src []T) []T {
	length, capacity := len(src), cap(src)
	newCapacity, ok := calNewCapacity(length, capacity)
	if !ok {
		return src
	}
	s := make([]T, 0, newCapacity)
	s = append(s, src...)
	return s

}
