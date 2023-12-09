package helpers

import "golang.org/x/exp/constraints"

func ForEach[T constraints.Ordered](slice []T, fn func(index int, value T)) {
	for index, value := range slice {
		fn(index, value)
	}
}
