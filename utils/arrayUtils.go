package utils

func IsEmptyArray[T any](data []T) bool {
	return len(data) == 0
}
