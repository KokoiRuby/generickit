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
