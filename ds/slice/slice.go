package slice

import "fmt"

type Slice[T any] []T

// New - create a new Slice
func (s Slice[T]) New(v T) Slice[T] {
	return append(s, v)
}

// Append - append a value to the Slice
func (s Slice[T]) Append(v T) Slice[T] {
	return append(s, v)
}

// Prepend - prepend a value to the Slice
func (s Slice[T]) Prepend(v T) Slice[T] {
	return append([]T{v}, s...)
}

// Print - print the Slice
func (s Slice[T]) Print() {
	fmt.Println(s)
}

// Len - get the length of the Slice
func (s Slice[T]) Len() int {
	return len(s)
}

// Get - get the value at the given index
func (s Slice[T]) Get(i int) T {
	return s[i]
}

// Set - set the value at the given index
func (s Slice[T]) Set(i int, v T) Slice[T] {
	s[i] = v
	return s
}

// Delete - delete the value at the given index
func (s Slice[T]) Delete(i int) Slice[T] {
	s = append(s[:i], s[i+1:]...)
	return s
}

// DeleteLast - delete the last value in the Slice
func (s Slice[T]) DeleteLast() Slice[T] {
	s = s[:len(s)-1]
	return s
}

// DeleteFirst - delete the first value in the Slice
func (s Slice[T]) DeleteFirst() Slice[T] {
	s = s[1:]
	return s
}

// DeleteAll - delete all values in the Slice
func (s Slice[T]) DeleteAll() Slice[T] {
	s = s[:0]
	return s
}

// Reverse - reverse the order of the Slice
func (s Slice[T]) Reverse() Slice[T] {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	return s
}
