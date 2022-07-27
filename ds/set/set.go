package set

import (
	"fmt"
	"strings"
)

// Set - Set is a data structure that stores a collection of unique values.

type Set[T comparable] map[T]struct{}

// New - Creates a new Set.
func (s Set[T]) New() Set[T] {
	return make(Set[T])
}

// Add - Adds a value to the Set.
func (s Set[T]) Add(value T) {
	s[value] = struct{}{}
}

// Remove - Removes a value from the Set.
func (s Set[T]) Remove(value T) {
	delete(s, value)
}

// Contains - Checks if a value is in the Set.
func (s Set[T]) Contains(value T) bool {
	_, ok := s[value]
	return ok
}

// Union - Returns a new Set that is the union of two Sets.
func (s Set[T]) Union(other Set[T]) Set[T] {
	union := s.New()
	for value := range s {
		union.Add(value)
	}
	for value := range other {
		union.Add(value)
	}
	return union
}

// Intersection - Returns a new Set that is the intersection of two Sets.
func (s Set[T]) Intersection(other Set[T]) Set[T] {
	intersection := s.New()
	for value := range s {
		if other.Contains(value) {
			intersection.Add(value)
		}
	}
	return intersection
}

// Difference - Returns a new Set that is the difference of two Sets.
func (s Set[T]) Difference(other Set[T]) Set[T] {
	difference := s.New()
	for value := range s {
		if !other.Contains(value) {
			difference.Add(value)
		}
	}
	return difference
}

// Subset - Checks if one Set is a subset of another.
func (s Set[T]) Subset(other Set[T]) bool {
	for value := range s {
		if !other.Contains(value) {
			return false
		}
	}
	return true
}

// Equal - Checks if two Sets are equal.
func (s Set[T]) Equal(other Set[T]) bool {
	return s.Subset(other) && other.Subset(s)
}

// Empty - Checks if a Set is empty.
func (s Set[T]) Empty() bool {
	return len(s) == 0
}

// Size - Returns the size of a Set.
func (s Set[T]) Size() int {
	return len(s)
}

// Clear - Removes all values from a Set.
func (s Set[T]) Clear() {
	for value := range s {
		s.Remove(value)
	}
}

// Values - Returns a slice of all values in a Set.
func (s Set[T]) Values() []T {
	values := make([]T, 0, len(s))
	for value := range s {
		values = append(values, value)
	}
	return values
}

// String - Returns a string representation of a Set.
func (s Set[T]) String() string {
	sets := make([]string, 0, len(s))
	for value := range s {
		sets = append(sets, fmt.Sprintf("%v", value))
	}
	return "{" + strings.Join(sets, ", ") + "}"
}

// ForEach - Calls a function for each value in a Set.
func (s Set[T]) ForEach(f func(T)) {
	for value := range s {
		f(value)
	}
}

// Map - Returns a new Set that is the result of calling a function on each value in a Set.
func (s Set[T]) Map(f func(T) T) Set[T] {
	mapped := s.New()
	for value := range s {
		mapped.Add(f(value))
	}
	return mapped
}

// Filter - Returns a new Set that is the result of calling a function on each value in a Set.
func (s Set[T]) Filter(f func(T) bool) Set[T] {
	filtered := s.New()
	for value := range s {
		if f(value) {
			filtered.Add(value)
		}
	}
	return filtered
}

// MapToSlice - Returns a slice of all values in a Set.
func (s Set[T]) MapToSlice(f func(T) interface{}) []interface{} {
	values := make([]any, 0, len(s))
	for value := range s {
		values = append(values, f(value))
	}
	return values
}
