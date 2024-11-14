package tools

// Contains function checks if a given entry exits in a slice of any comparable
// type.
func Contains[T comparable](slice []T, entry T) (int, bool) {
	for index, item := range slice {
		if item == entry {
			return index, true
		}
	}
	return -1, false
}

// Removes function removes a given entry if exists in a slice of any
// comparable type.
func Removes[T comparable](slice []T, entry T) []T {
	if index, ok := Contains(slice, entry); ok {
		slice = append(slice[:index], slice[index+1:]...)
	}
	return slice
}

// Equal function checks if two values of comparable type T are equal.
func Equal[T comparable](a, b T) bool {
	return a == b
}

// ContainsAny function checks if a given entry exist in a slice of a type
// using the given equal function to compare entries.
func ContainsAny[T any](slice []T, entry T, equal func(T, T) bool) (int, bool) {
	for index, item := range slice {
		if equal(item, entry) {
			return index, true
		}
	}
	return -1, false
}

// RemovesAny function removes a given entry if exist in a slice of any type
// using the givem equal functon to compare entries.
func RemovesAny[T any](slice []T, entry T, equal func(T, T) bool) []T {
	if index, ok := ContainsAny(slice, entry, equal); ok {
		slice = append(slice[:index], slice[index+1:]...)
	}
	return slice
}
