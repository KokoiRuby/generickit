package slice

import "fmt"

func ErrIdxOutOfRange(length int, idx int) error {
	return fmt.Errorf("idx out of range, length: %d, idx: %d", length, idx)
}

func ErrElemNotFound() error {
	return fmt.Errorf("element not found")
}
