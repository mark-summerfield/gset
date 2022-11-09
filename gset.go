// Copyright © 2022 Mark Summerfield. All rights reserved.
// License: Apache-2.0

// Set is a generic set type.
//
// In addition to its methods it also supports the built-in len() function.
// And since Set is built on map you can use map methods (e.g., to iterate;
// see [ToSlice]).
//
// See [New] for how to create empty or populated sets.
package gset

import (
	_ "embed"
	"fmt"
	"sort"
)

//go:embed Version.dat
var Version string // This module's version.

type null struct{}

// Set is a generic set type. (`null` is an alias for `struct{}`.)
type Set[T comparable] map[T]null

// New returns a new set containing the given elements (if any).
// If no elements are given, the type must be specified since it can't be
// inferred.
//
// Examples:
//
//	a := New[int]() // Type must be specified since it can't be inferred.
//	b := New("a string")
//	c := New(1, 2, 4, 8)
//	s := []string{"A", "B", "C", "De", "Fgh"}
//	d := New(s...)
func New[T comparable](elements ...T) Set[T] {
	set := make(Set[T], len(elements))
	for _, element := range elements {
		set[element] = null{}
	}
	return set
}

// String returns a human readable string representation of the set.
func (me Set[T]) String() string {
	elements := make([]T, 0, len(me))
	for element := range me {
		elements = append(elements, element)
	}
	sort.Slice(elements, func(i, j int) bool {
		return less(elements[i], elements[j])
	})
	s := "{"
	sep := ""
	for _, element := range elements {
		s += sep + asStr(element)
		sep = " "
	}
	return s + "}"
}

func asStr(x any) string {
	if s, ok := x.(string); ok {
		return fmt.Sprintf("%q", s)
	}
	return fmt.Sprintf("%v", x)
}

func less(a, b any) bool {
	switch x := a.(type) {
	case int:
		return x < b.(int)
	case float64:
		return x < b.(float64)
	case string:
		return x < b.(string)
	default:
		return fmt.Sprintf("%v", a) < fmt.Sprintf("%v", b)
	}
}

// ToSlice returns this set's elements as a slice.
// For iteration either use this, or if you only need one value at a time,
// use map syntax with a for loop.
// Example:
//
//	s := gset.New(2, 3, 5, 7, 11, 13)
//	slice := s.ToSlice() // Copies the lot
//	for _, v := range slice {
//		fmt.Println(v)
//	}
//	// Alternatively, one value at a time:
//	for v := range s {
//		fmt.Println(v)
//	}
func (me Set[T]) ToSlice() []T {
	result := make([]T, 0, len(me))
	for element := range me {
		result = append(result, element)
	}
	return result
}

// Add adds the given element(s) to the set.
func (me Set[T]) Add(elements ...T) {
	for _, element := range elements {
		me[element] = null{}
	}
}

// Delete deletes the given element(s) from the set.
func (me Set[T]) Delete(elements ...T) {
	for _, element := range elements {
		delete(me, element)
	}
}

// Clear deletes all the elements to make this an empty set.
func (me Set[T]) Clear() {
	for element := range me {
		delete(me, element)
	}
}

// Contains returns true if element is in the set; otherwise returns false.
// Alternatively, use map syntax.
// Example:
//
//	s := gset.New("X", "Y", "Z")
//	if s.Contains("Y") {
//		fmt.Println("Got Y")
//	}
//	// Alternatively, use map syntax
//	if _, ok := s["Y"]; ok {
//		fmt.Println("Got Y")
//	}
func (me Set[T]) Contains(element T) bool {
	_, found := me[element]
	return found
}

// Difference returns a new set that contains the elements which are in this
// set that are not in the other set.
func (me Set[T]) Difference(other Set[T]) Set[T] {
	diff := make(Set[T])
	for element := range me {
		if !other.Contains(element) {
			diff[element] = null{}
		}
	}
	return diff
}

// SymmetricDifference returns a new set that contains the elements which
// are in this set or the other set—but not in both sets.
func (me Set[T]) SymmetricDifference(other Set[T]) Set[T] {
	diff := make(Set[T])
	for element := range me {
		if !other.Contains(element) {
			diff[element] = null{}
		}
	}
	for element := range other {
		if !me.Contains(element) {
			diff[element] = null{}
		}
	}
	return diff
}

// Intersection returns a new set that contains the elements this set has in
// common with the other set.
func (me Set[T]) Intersection(other Set[T]) Set[T] {
	intersection := make(Set[T])
	for element := range me {
		if other.Contains(element) {
			intersection[element] = null{}
		}
	}
	for element := range other {
		if me.Contains(element) {
			intersection[element] = null{}
		}
	}
	return intersection
}

// Union returns a new set that contains the elements from this set and from
// the other set (with no duplicates of course).
func (me Set[T]) Union(other Set[T]) Set[T] {
	union := make(Set[T], len(me))
	for element := range me {
		union[element] = null{}
	}
	for element := range other {
		union[element] = null{}
	}
	return union
}

// Unite adds all the elements from other that aren't already in this set to
// this set.
func (me Set[T]) Unite(other Set[T]) {
	for element := range other {
		me[element] = null{}
	}
}

// Copy returns a copy of this set.
func (me Set[T]) Copy() Set[T] {
	other := make(Set[T], len(me))
	for element := range me {
		other[element] = null{}
	}
	return other
}

// Equal returns true if this set has the same elements as the other set;
// otherwise returns false.
func (me Set[T]) Equal(other Set[T]) bool {
	if len(me) != len(other) {
		return false
	}
	// If they have the same number of elements then if any element in this
	// is not in the other then they're different.
	for element := range me {
		if !other.Contains(element) {
			return false
		}
	}
	return true
}

// IsDisjoint returns true if this set has no elements in common with the
// other set; otherwise returns false.
func (me Set[T]) IsDisjoint(other Set[T]) bool {
	for element := range me {
		if other.Contains(element) {
			return false
		}
	}
	for element := range other {
		if me.Contains(element) {
			return false
		}
	}
	return true
}
