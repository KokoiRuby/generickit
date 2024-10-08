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

func Contains[T comparable](sl []T, val T) bool {
	return ContainsFunc[T](sl, func(v T) bool {
		return v == val
	})
}

func ContainsFunc[T any](sl []T, match matchFunc[T]) bool {
	for _, v := range sl {
		if match(v) {
			return true
		}
	}
	return false
}
