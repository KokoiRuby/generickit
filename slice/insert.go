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
