package set

import (
	"fmt"
	"sort"
	"strings"
)

// Set - Set is a data structure that stores a collection of unique values.

type set[T comparable] map[T]struct{}

// New - Creates a new Set.
func New[T comparable]() set[T] {
	return make(set[T], 0)
}

// Add - Adds a value to the Set.
func (s set[T]) Add(value T) {
	s[value] = struct{}{}
}

// Remove - Removes a value from the Set.
func (s set[T]) Remove(value T) {
	delete(s, value)
}

// Contains - Checks if a value is in the Set.
func (s set[T]) Contains(value T) bool {
	_, ok := s[value]
	return ok
}

// Union - Returns a new Set that is the union of two Sets.
func (s set[T]) Union(other set[T]) set[T] {
	union := make(set[T], 0)
	for value := range s {
		union.Add(value)
	}
	for value := range other {
		union.Add(value)
	}
	return union
}

// Intersection - Returns a new Set that is the intersection of two Sets.
func (s set[T]) Intersection(other set[T]) set[T] {
	intersection := make(set[T], 0)
	for value := range s {
		if other.Contains(value) {
			intersection.Add(value)
		}
	}
	return intersection
}

// Difference - Returns a new Set that is the difference of two Sets.
func (s set[T]) Difference(other set[T]) set[T] {
	difference := make(set[T], 0)
	for value := range s {
		if !other.Contains(value) {
			difference.Add(value)
		}
	}
	return difference
}

// Subset - Checks if one Set is a subset of another.
func (s set[T]) Subset(other set[T]) bool {
	for value := range s {
		if !other.Contains(value) {
			return false
		}
	}
	return true
}

// Equal - Checks if two Sets are equal.
func (s set[T]) Equal(other set[T]) bool {
	return s.Subset(other) && other.Subset(s)
}

// Empty - Checks if a Set is empty.
func (s set[T]) Empty() bool {
	return len(s) == 0
}

// Size - Returns the size of a Set.
func (s set[T]) Len() int {
	return len(s)
}

// Clear - Removes all values from a Set.
func (s set[T]) Clear() {
	for value := range s {
		s.Remove(value)
	}
}

// Values - Returns a slice of all values in a Set.
func (s set[T]) Values() []T {
	values := make([]T, 0, len(s))
	for value := range s {
		values = append(values, value)
	}
	return values
}

// String - Returns a string representation of a Set.
func (s set[T]) String() {
	sets := make([]string, 0, len(s))
	for value := range s {
		sets = append(sets, fmt.Sprintf("%v", value))
	}

	fmt.Println("{" + strings.Join(sets, ", ") + "}")
}

// ForEach - Calls a function for each value in a Set.
func (s set[T]) ForEach(f func(T)) {
	for value := range s {
		f(value)
	}
}

// Map - Returns a new Set that is the result of calling a function on each value in a Set.
func (s set[T]) Map(f func(T) T) set[T] {
	mapped := make(set[T], 0)
	for value := range s {
		mapped.Add(f(value))
	}
	return mapped
}

// Filter - Returns a new Set that is the result of calling a function on each value in a Set.
func (s set[T]) Filter(f func(T) bool) set[T] {
	filtered := make(set[T], 0)
	for value := range s {
		if f(value) {
			filtered.Add(value)
		}
	}
	return filtered
}

// Sort - Sorts a Set.
func (s set[T]) Sort(less func(T, T) bool) []T {
	values := s.Values()
	sort.Slice(s.Values(), func(i, j int) bool {
		return less(s.Values()[i], s.Values()[j])
	})
	return values
}
