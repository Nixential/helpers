package helpers

import (
	"golang.org/x/exp/constraints"
)

func ForEach[T constraints.Ordered](slice []T, fn func(index int, value T)) {
	for index, value := range slice {
		fn(index, value)
	}
}

func Map[T constraints.Ordered](slice []T, fn func(element T, index int, s []T) T) []T {

	var results []T

	for index, value := range slice {
		var result T = fn(value, index, slice)
		results = append(results, result)
	}

	return results
}

func Filter[T any](slice []T, fn func(T) bool) []T {
	var results []T

	for _, value := range slice {
		if fn(value) {
			results = append(results, value)
		}
	}

	return results
}
