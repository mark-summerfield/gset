// Copyright © 2022 Mark Summerfield. All rights reserved.
// License: Apache-2.0

// Set is a generic set type based on a map.
//
// Set supports all the map methods, functions that apply to maps (e.g.,
// len()), and has its own often more convenient API.
//
// See [New] for how to create empty or populated sets.
package gset

import (
	_ "embed"
	"fmt"
	"sort"
	"strings"
)

//go:embed Version.dat
var Version string // This module's version.

type Set[T comparable] map[T]struct{}

// New returns a new set containing the given elements (if any).
// If no elements are given, the type must be specified since it can't be
// inferred.
func New[T comparable](elements ...T) Set[T] {
	set := make(Set[T], len(elements))
	for _, element := range elements {
		set[element] = struct{}{}
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
	var s strings.Builder
	s.WriteString("{")
	sep := ""
	for _, element := range elements {
		s.WriteString(sep)
		if selement, ok := any(element).(string); ok {
			fmt.Fprintf(&s, "%q", selement)
		} else {
			fmt.Fprintf(&s, "%v", element)
		}
		sep = " "
	}
	s.WriteString("}")
	return s.String()
}

func less(a, b any) bool {
	switch x := a.(type) {
	case byte:
		return x < b.(byte)
	case int8:
		return x < b.(int8)
	case int16:
		return x < b.(int16)
	case int32:
		return x < b.(int32)
	case int64:
		return x < b.(int64)
	case int:
		return x < b.(int)
	case uint16:
		return x < b.(uint16)
	case uint32:
		return x < b.(uint32)
	case uint64:
		return x < b.(uint64)
	case uint:
		return x < b.(uint)
	case float32:
		return x < b.(float32)
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
// See also [ToSortedSlice].
func (me Set[T]) ToSlice() []T {
	result := make([]T, 0, len(me))
	for element := range me {
		result = append(result, element)
	}
	return result
}

// ToSortedSlice returns this set's elements as a slice with the elements
// sorted using <.
// For iteration either use this, or if you only need one value at a time,
// use map syntax with a for loop.
// See also [ToSlice].
func (me Set[T]) ToSortedSlice() []T {
	result := make([]T, 0, len(me))
	for element := range me {
		result = append(result, element)
	}
	sort.Slice(result, func(i, j int) bool {
		return less(result[i], result[j])
	})
	return result
}

// Add adds the given element(s) to the set.
func (me Set[T]) Add(elements ...T) {
	for _, element := range elements {
		me[element] = struct{}{}
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

// IsEmpty returns true if the set is empty; otherwise returns false.
// This is just a convenience for len(s) == 0.
func (me Set[T]) IsEmpty() bool { return len(me) == 0 }

// Contains returns true if element is in the set; otherwise returns false.
// Alternatively, use map syntax.
func (me Set[T]) Contains(element T) bool {
	_, found := me[element]
	return found
}

// Difference returns a new set that contains the elements which are in this
// set that are not in the other set.
func (me Set[T]) Difference(other Set[T]) Set[T] {
	diff := Set[T]{}
	for element := range me {
		if !other.Contains(element) {
			diff[element] = struct{}{}
		}
	}
	return diff
}

// SymmetricDifference returns a new set that contains the elements which
// are in this set or the other set—but not in both sets.
func (me Set[T]) SymmetricDifference(other Set[T]) Set[T] {
	diff := Set[T]{}
	for element := range me {
		if !other.Contains(element) {
			diff[element] = struct{}{}
		}
	}
	for element := range other {
		if !me.Contains(element) {
			diff[element] = struct{}{}
		}
	}
	return diff
}

// Intersection returns a new set that contains the elements this set has in
// common with the other set.
func (me Set[T]) Intersection(other Set[T]) Set[T] {
	intersection := Set[T]{}
	for element := range me {
		if other.Contains(element) {
			intersection[element] = struct{}{}
		}
	}
	for element := range other {
		if me.Contains(element) {
			intersection[element] = struct{}{}
		}
	}
	return intersection
}

// Union returns a new set that contains the elements from this set and from
// the other set (with no duplicates of course).
// See also [Set.Unite].
func (me Set[T]) Union(other Set[T]) Set[T] {
	union := make(Set[T], len(me))
	for element := range me {
		union[element] = struct{}{}
	}
	for element := range other {
		union[element] = struct{}{}
	}
	return union
}

// Unite adds all the elements from other that aren't already in this set to
// this set.
// See also [Set.Union].
func (me Set[T]) Unite(other Set[T]) {
	for element := range other {
		me[element] = struct{}{}
	}
}

// Copy returns a copy of this set.
func (me Set[T]) Copy() Set[T] {
	other := make(Set[T], len(me))
	for element := range me {
		other[element] = struct{}{}
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
