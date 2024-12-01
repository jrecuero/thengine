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

// ContainsGenericAny function checks if a given entry exists in a slice where
// slice and entry have different type and the give equal function is used to
// match the equality.
func ContainsGenericAny[T, W any](slice []T, entry W, equal func(T, W) bool) (int, bool) {
	for index, item := range slice {
		if equal(item, entry) {
			return index, true
		}
	}
	return -1, false
}

// ContainsAny function checks if a given entry exists in a slice of a type
// using the given equal function to compare entries.
func ContainsAny[T any](slice []T, entry T, equal func(T, T) bool) (int, bool) {
	for index, item := range slice {
		if equal(item, entry) {
			return index, true
		}
	}
	return -1, false
}

// Equal function checks if two values of comparable type T are equal.
func Equal[T comparable](a, b T) bool {
	return a == b
}

// Removes function removes a given entry if exists in a slice of any
// comparable type.
func Removes[T comparable](slice []T, entry T) []T {
	if index, ok := Contains(slice, entry); ok {
		slice = append(slice[:index], slice[index+1:]...)
	}
	return slice
}

// RemovesGeneric function removes a given entry if exists int a slice of any
// type which could be different from the entry, and the give equal function is
// used for checking equality.
func RemovesGenericAny[T, W any](slice []T, entry W, equal func(T, W) bool) []T {
	if index, ok := ContainsGenericAny(slice, entry, equal); ok {
		slice = append(slice[:index], slice[index+1:]...)
	}
	return slice
}

// RemovesAny function removes a given entry if exists in a slice of any type
// using the given equal functon to compare entries.
func RemovesAny[T any](slice []T, entry T, equal func(T, T) bool) []T {
	if index, ok := ContainsAny(slice, entry, equal); ok {
		slice = append(slice[:index], slice[index+1:]...)
	}
	return slice
}

// Retrieve function retrieves an entry in the slice if it matches the given
// input parameter.
func Retrieve[T comparable](slice []T, entry T) (T, bool) {
	if index, ok := Contains(slice, entry); ok {
		return slice[index], true
	}
	var zero T
	return zero, false
}

// RetriveGenericAny function retrieves an entry in the slice if it matches
// base on the give equal function, used for checking functionality. Slice and
// entry type can be different.
func RetrieveGenericAny[T, W any](slice []T, entry W, equal func(T, W) bool) (T, bool) {
	if index, ok := ContainsGenericAny(slice, entry, equal); ok {
		return slice[index], true
	}
	var zero T
	return zero, false
}

// RetrieveAny function retrieves an entry in the slice if it matches
// base on the give equal function, used for checking functionality. Slice and
// entry type should be the same.
func RetrieveAny[T any](slice []T, entry T, equal func(T, T) bool) (T, bool) {
	if index, ok := ContainsAny(slice, entry, equal); ok {
		return slice[index], true
	}
	var zero T
	return zero, false
}
