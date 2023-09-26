package ptr

func New[T any](i T) *T {
	return &i
}

func Dereference[T any](ptr *T, default_ T) T {
	if ptr != nil {
		return *ptr
	}
	return default_
}
