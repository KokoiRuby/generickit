package slice

// Reverse in-place
func Reverse[T any](s []T) []T {
	low := 0
	high := len(s) - 1
	for low < high {
		s[low], s[high] = s[high], s[low]
		low++
		high--
	}
	return s
}
