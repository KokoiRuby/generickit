package slice

import "generickit"

func Max[T generickit.RealNumber](sl []T) T {
	res := sl[0]
	for _, v := range sl {
		if v > res {
			res = v
		}
	}
	return res
}

func Min[T generickit.RealNumber](sl []T) T {
	res := sl[0]
	for _, v := range sl {
		if v < res {
			res = v
		}
	}
	return res
}

func Sum[T generickit.Number](sl []T) T {
	var sum T
	for _, v := range sl {
		sum += v
	}
	return sum
}
