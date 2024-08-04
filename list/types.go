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

type List[T any] interface {
	// Get the element given idx, err if out of range
	Get(idx int) (T, error)

	// Append a element(s)
	Append(ts ...T)

	// Add an element at idx, err if out of range
	// Append if idx == Len()
	Add(idx int, t T) error

	// Set the element given idx, err if out of range
	Set(idx int, t T) error

	// Delete the element given idx, err if out of range
	Delete(idx int) error

	// Len Length of
	Len() int

	// Cap Capacity of
	Cap() int

	// Range iteration
	Range(fn func(idx int, t T) error) error

	// ToSlice transform List to a slice, []T if no element
	ToSlice() []T

	// Generator yield-like
	Generator() <-chan T
}
