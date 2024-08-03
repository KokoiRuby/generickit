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

type Number interface {
	constraints.Integer | constraints.Float
}

func Map[Src any, Dst any](src []Src, mapper func(src Src) Dst) []Dst {
	dst := make([]Dst, 0)
	for _, v := range src {
		dst = append(dst, mapper(v))
	}
	return dst
}

func Filter[Src any, Dst any](src []Src, filter func(src Src) (dst Dst, filter bool)) []Dst {
	dst := make([]Dst, 0)
	for _, v := range src {
		if val, ok := filter(v); ok {
			dst = append(dst, val)
		}
	}
	return dst
}

// Reduce usually for aggregation, Dst in Number type in order to support operators.
func Reduce[Src any, Dst Number](src []Src, reducer func(src Src) Dst) Dst {
	var sum Dst
	for _, v := range src {
		sum += reducer(v)
	}
	return sum
}
