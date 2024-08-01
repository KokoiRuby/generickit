package slice

func Delete[T any](dst []T, idx int) ([]T, error) {
	length := len(dst)
	if idx < 0 || idx > length-1 {
		return dst, ErrIdxOutOfRange(length, idx)
	}
	copy(dst[idx:], dst[idx+1:]) // mv 1 step backward
	dst = dst[:len(dst)-1]       // cut-off last
	return dst, nil
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

func DeleteVal[T any](dst []T, val T) ([]T, error) {
	return dst, nil
}
