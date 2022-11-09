package gset_test

import (
	"fmt"
	"github.com/mark-summerfield/gset"
)

func ExampleNew_a() {
	a := gset.New[int]() // Type must be specified since it can't be inferred.
	fmt.Println(a)
	// Output: {}
}

func ExampleNew_b() {
	b := gset.New(8, 7, 1, 2, 4, 3, 8)
	fmt.Println(b)
	// Output: {1 2 3 4 7 8}
}

func ExampleNew_c() {
	s := []string{"De", "A", "Fgh", "C", "B"}
	c := gset.New(s...)
	fmt.Println(c)
	// Output: {"A" "B" "C" "De" "Fgh"}
}

func ExampleSet_String() {
	s := gset.New("one", "two", "three", "four", "five", "six")
	fmt.Println(s)
	// Output: {"five" "four" "one" "six" "three" "two"}
}

func ExampleSet_ToSlice() {
	total1 := 0
	s := gset.New(2, 3, 5, 7, 11, 13)
	slice := s.ToSlice() // Copies the lot
	for _, v := range slice {
		total1 += v
	}
	total2 := 0
	// Alternatively, one value at a time using map syntax:
	for v := range s {
		total2 += v
	}
	fmt.Println(total1, total2, total1 == total2)
	// Output: 41 41 true
}

func ExampleSet_Contains() {
	count := 0
	s := gset.New("X", "Y", "Z")
	if s.Contains("Y") {
		count += 1
	}
	// Alternatively, use map syntax:
	if _, ok := s["Y"]; ok {
		count += 1
	}
	fmt.Println(count, count == 2)
	// Output: 2 true
}

func ExampleSet_Union() {
	s := gset.New(23, 19, 17, 13, 11)
	t := gset.New(2, 3, 5, 7, 11, 13)
	u := s.Union(t)
	v := t.Union(s)
	fmt.Println(u.Equal(v), u, t)
	// Output: true {2 3 5 7 11 13 17 19 23} {2 3 5 7 11 13}
}

func ExampleSet_Unite() {
	s := gset.New(23, 19, 17, 13, 11)
	s.Unite(gset.New(2, 3, 5, 7, 11, 13))
	fmt.Println(s)
	// Output: {2 3 5 7 11 13 17 19 23}
}
