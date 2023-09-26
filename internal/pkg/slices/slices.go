package slices

func Copy[T any](i []T) []T {
	if i == nil {
		return nil
	}
	cp := make([]T, len(i))
	copy(cp, i)
	return cp
}
