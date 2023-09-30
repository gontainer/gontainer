package maps

import (
	"sort"
)

// Keys returns sorted keys of the input map.
func Keys[K ~string, V any](input map[K]V) []K {
	keys := make([]K, 0, len(input))
	for k := range input {
		keys = append(keys, k)
	}
	sort.Slice(keys, func(i, j int) bool {
		return keys[i] < keys[j]
	})
	return keys
}

// Iterate iterates over a map in the order determined by Keys.
func Iterate[K ~string, V any](input map[K]V, fn func(k K, v V)) {
	for _, n := range Keys(input) {
		fn(n, input[n])
	}
}
