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

package queue

import (
	"context"
	"unsafe"
)

// Queue is non-blocking
type Queue[T any] interface {
	Enqueue(t T) error
	Dequeue() (T, error)
}

// BlockingQueue using ctx timeout
type BlockingQueue[T any] interface {
	Enqueue(ctx context.Context, t T) error
	Dequeue(ctx context.Context) (T, error)
}

type node[T any] struct {
	val T
	// *node[T], unsafe.Pointer can be transformed into pointer of any type
	next unsafe.Pointer
}
