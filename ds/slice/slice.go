package slice

import "fmt"

type slice[T any] []T

// New - create a new Slice
func New[T any]() slice[T] {
	return make(slice[T], 0)
}

// Append - append a value to the Slice
func (s slice[T]) Append(v T) slice[T] {
	return append(s, v)
}

// Prepend - prepend a value to the Slice
func (s slice[T]) Prepend(v T) slice[T] {
	return append([]T{v}, s...)
}

// Print - print the Slice
func (s slice[T]) Print() {
	fmt.Println(s)
}

// Len - get the length of the Slice
func (s slice[T]) Len() int {
	return len(s)
}

// Get - get the value at the given index
func (s slice[T]) Get(i int) T {
	return s[i]
}

// Set - set the value at the given index
func (s slice[T]) Set(i int, v T) slice[T] {
	s[i] = v
	return s
}

// Delete - delete the value at the given index
func (s slice[T]) Delete(i int) slice[T] {
	s = append(s[:i], s[i+1:]...)
	return s
}

// DeleteLast - delete the last value in the Slice
func (s slice[T]) DeleteLast() slice[T] {
	s = s[:len(s)-1]
	return s
}

// DeleteFirst - delete the first value in the Slice
func (s slice[T]) DeleteFirst() slice[T] {
	s = s[1:]
	return s
}

// DeleteAll - delete all values in the Slice
func (s slice[T]) DeleteAll() slice[T] {
	s = s[:0]
	return s
}

// Reverse - reverse the order of the Slice
func (s slice[T]) Reverse() slice[T] {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	return s
}
