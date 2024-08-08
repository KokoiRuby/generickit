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

import (
	"golang.org/x/exp/constraints"
)

func Delete[T any](dst []T, idx int) (T, []T, error) {
	length := len(dst)
	var t T
	if idx < 0 || idx > length-1 {
		return t, dst, ErrIdxOutOfRange(length, idx)
	}
	t = dst[idx]
	dst = append(dst[:idx], dst[idx+1:]...)
	return t, dst, nil
}

func DeleteAfter[T any](dst []T, idx int) ([]T, error) {
	length := len(dst)
	if idx < 0 || idx > length-1 {
		return dst, ErrIdxOutOfRange(length, idx)
	}
	if idx == 0 {
		return []T{}, nil
	}
	return dst[:idx], nil
}

func DeleteRange[T any](dst []T, r int, idx int) ([]T, error) {
	length := len(dst)
	if idx < 0 || idx > length-1 {
		return dst, ErrIdxOutOfRange(length, idx)
	}
	if r > length-idx+1 {
		return DeleteAfter(dst, idx)
	}
	dst = append(dst[0:idx], dst[idx+r:]...)
	return dst, nil

}

func DeleteVal[T constraints.Ordered](dst []T, val T) ([]T, error) {
	found := false
	for i, v := range dst {
		if v == val {
			dst = append(dst[:i], dst[i+1:]...)
			found = true
			break
		}
	}
	if !found {
		return dst, ErrElemNotFound()
	}
	return dst, nil
}
