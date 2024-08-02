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

func Find[T comparable](sl []T, val T) (idx int, isFound bool) {
	return FindFunc(sl, func(v T) bool {
		return v == val
	})
}

func FindLast[T comparable](sl []T, val T) (idx int, isFound bool) {
	return FindLastFunc(sl, func(v T) bool {
		return v == val
	})
}

func FindAll[T comparable](sl []T, val T) (idx []int, isFound bool) {
	return FindAllFunc(sl, func(v T) bool {
		return v == val
	})
}

// FindFunc accepts a anonymous func to match
func FindFunc[T any](slice []T, match matchFunc[T]) (idx int, isFound bool) {
	for i, v := range slice {
		if match(v) {
			return i, true
		}
	}
	return -1, false
}

// FindLastFunc accepts a anonymous func to match
func FindLastFunc[T any](slice []T, match matchFunc[T]) (idx int, isFound bool) {
	for i := len(slice) - 1; i >= 0; i-- {
		if match(slice[i]) {
			return i, true
		}
	}
	return -1, false
}

// FindAllFunc accepts a anonymous func to match
func FindAllFunc[T any](sl []T, match matchFunc[T]) (idx []int, isFound bool) {
	// hard to decide capacity right here...
	tmpSlice := make([]int, 0, len(sl))
	flag := false
	for i, v := range sl {
		if match(v) {
			tmpSlice = append(tmpSlice, i)
			flag = true
		}
	}
	return tmpSlice, flag
}
