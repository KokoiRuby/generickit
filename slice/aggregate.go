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

import "generickit"

func Max[T generickit.RealNumber](sl []T) T {
	res := sl[0]
	for _, v := range sl {
		if v > res {
			res = v
		}
	}
	return res
}

func Min[T generickit.RealNumber](sl []T) T {
	res := sl[0]
	for _, v := range sl {
		if v < res {
			res = v
		}
	}
	return res
}

func Sum[T generickit.Number](sl []T) T {
	var sum T
	for _, v := range sl {
		sum += v
	}
	return sum
}
