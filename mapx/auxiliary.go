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

import "fmt"

func Keys[K comparable, V any](m map[K]V) []K {
	kSl := make([]K, 0, len(m))
	for k, _ := range m {
		kSl = append(kSl, k)
	}
	return kSl
}

func Values[K comparable, V any](m map[K]V) []V {
	vSl := make([]V, 0, len(m))
	for _, v := range m {
		vSl = append(vSl, v)
	}
	return vSl
}

func KeysValues[K comparable, V any](m map[K]V) ([]K, []V) {
	kSl := make([]K, 0, len(m))
	vSl := make([]V, 0, len(m))
	for k, v := range m {
		kSl = append(kSl, k)
		vSl = append(vSl, v)
	}
	return kSl, vSl
}

func ToMap[K comparable, V any](kSl []K, vSl []V) (map[K]V, error) {
	if kSl == nil || vSl == nil {
		return nil, fmt.Errorf("key slice or value slice is nil")
	}
	keyLen := len(kSl)
	valLen := len(vSl)
	if keyLen != valLen {
		return nil, fmt.Errorf("key slice or value slice length not equal")
	}
	m := make(map[K]V)
	for i, k := range kSl {
		m[k] = vSl[i]
	}
	return m, nil

}
