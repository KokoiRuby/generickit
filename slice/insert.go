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

func Insert[T any](dst []T, value T, idx int) ([]T, error) {
	length := len(dst)
	if idx < 0 || idx > length-1 {
		return dst, ErrIdxOutOfRange(length, idx)
	}
	var zVal T
	dst = append(dst, zVal)
	copy(dst[idx+1:], dst[idx:]) // mv 1 step forward
	dst[idx] = value

	return dst, nil
}

func InsertSlice[T any](dst []T, src []T, idx int) ([]T, error) {
	length := len(dst)
	if idx < 0 || idx > length-1 {
		return dst, ErrIdxOutOfRange(length, idx)
	}
	length = len(src)
	if length == 0 {
		return dst, nil
	}
	tmp := make([]T, length)
	dst = append(dst, tmp...)
	copy(dst[idx+length:], dst[idx:]) // mv len(src) step forward
	copy(dst[idx:], src)
	return dst, nil
}
