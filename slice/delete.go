package slice

import (
	"golang.org/x/exp/constraints"
)

func Delete[T any](dst []T, idx int) ([]T, error) {
	length := len(dst)
	if idx < 0 || idx > length-1 {
		return dst, ErrIdxOutOfRange(length, idx)
	}
	dst = append(dst[:idx], dst[idx+1:]...)
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
