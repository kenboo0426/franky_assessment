package common

func Ptr[T any](value T) *T {
	return &value
}

func Val[T any](value *T) T {
	if value == nil {
		return *new(T)
	}
	return *value
}
