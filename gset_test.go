package gset

import (
	"fmt"
	"sort"
	"testing"
)

func check(act string, actSize int, exp string, expSize int, t *testing.T) {
	if actSize != expSize {
		t.Errorf("expected %d elements, got %d", expSize, actSize)
	}
	if exp != act {
		t.Errorf("expected %s, got %s", exp, act)
	}
}

func TestNew(t *testing.T) {
	s1 := New[int]()
	check(s1.String(), len(s1), "{}", 0, t)
	s2 := New(5)
	check(s2.String(), len(s2), "{5}", 1, t)
	s3 := New(50, 35, 78)
	check(s3.String(), len(s3), "{35 50 78}", 3, t)
	s4 := New("one", "two")
	check(s4.String(), len(s4), "{\"one\" \"two\"}", 2, t)
	a := New[int]()
	check(a.String(), len(a), "{}", 0, t)
	b := New("a string")
	check(b.String(), len(b), "{\"a string\"}", 1, t)
	c := New(19, 21, 1, 2, 4, 8)
	check(c.String(), len(c), "{1 2 4 8 19 21}", 6, t)
	s := []string{"A", "B", "C", "De", "Fgh"}
	d := New(s...)
	check(d.String(), len(d), "{\"A\" \"B\" \"C\" \"De\" \"Fgh\"}", len(s),
		t)
}

func TestToSlice(t *testing.T) {
	s := New(19, 21, 1, 2, 4, 8)
	u := s.ToSlice()
	sort.Ints(u)
	check(fmt.Sprintf("%v", u), len(u), "[1 2 4 8 19 21]", len(s), t)
}

func TestAdd(t *testing.T) {
	s := New(19, 21, 1, 2, 4, 8)
	s.Add(5, 7, 1, 19)
	check(s.String(), len(s), "{1 2 4 5 7 8 19 21}", 8, t)
}

func TestDelete(t *testing.T) {
	s := New(19, 21, 1, 2, 5, 4, 8, 9, 11, 13, 7)
	s.Delete(5, 7, 1, 19)
	check(s.String(), len(s), "{2 4 8 9 11 13 21}", 7, t)
}

func TestClear(t *testing.T) {
	s := New(19, 21, 1, 2, 5, 4, 8, 9, 11, 13, 7)
	s.Clear()
	check(s.String(), len(s), "{}", 0, t)
}

func TestContains(t *testing.T) {
	s := New(19, 21, 1, 2, 5, 4, 8, 9, 11, 13, 7)
	if !s.Contains(11) {
		t.Error("expected set to contain 11")
	}
	if s.Contains(23) {
		t.Error("expected set not to contain 23")
	}
}

func TestDifference(t *testing.T) {
	s := New(0, 1, 2, 3, 4, 5, 6, 7, 8, 9)
	u := New(2, 4, 6, 8)
	d := s.Difference(u)
	check(d.String(), len(d), "{0 1 3 5 7 9}", 6, t)
	d = u.Difference(s)
	check(d.String(), len(d), "{}", 0, t)
}

func TestSymmetricDifference(t *testing.T) {
	s := New(0, 1, 2, 3, 4, 5, 6, 7, 8, 9)
	u := New(2, 4, 6, 8)
	d := s.SymmetricDifference(u)
	check(d.String(), len(d), "{0 1 3 5 7 9}", 6, t)
	d = u.SymmetricDifference(s)
	check(d.String(), len(d), "{0 1 3 5 7 9}", 6, t)
}

func TestIntersection(t *testing.T) {
	s := New(0, 1, 2, 3, 4, 5, 6, 7, 8, 9)
	u := New(2, 4, 6, 8)
	x := s.Intersection(u)
	check(x.String(), len(x), "{2 4 6 8}", 4, t)
}

func TestUnion(t *testing.T) {
	s := New(0, 1, 2, 3, 4, 5, 6, 7, 8, 9)
	u := New(2, 4, 6, 8, 10, 12)
	x := s.Union(u)
	check(x.String(), len(x), "{0 1 2 3 4 5 6 7 8 9 10 12}", 12, t)
}

func TestUnite(t *testing.T) {
	s := New(0, 1, 2, 3, 4, 5, 6, 7, 8, 9)
	s.Unite(New(2, 4, 6, 8, 10, 12))
	check(s.String(), len(s), "{0 1 2 3 4 5 6 7 8 9 10 12}", 12, t)
}

func TestCopy(t *testing.T) {
	s := New(0, 1, 2, 3, 4, 5, 6, 7, 8, 9)
	u := s.Copy()
	check(s.String(), len(s), u.String(), len(u), t)
}

func TestEqual(t *testing.T) {
	s := New(0, 1, 2, 3, 4, 5, 6, 7, 8, 9)
	u := s.Copy()
	if !s.Equal(u) {
		t.Errorf("%s != %s", s, u)
	}
	u.Add(-3)
	if s.Equal(u) {
		t.Errorf("%s == %s", s, u)
	}
}
func TestIsDisjoing(t *testing.T) {
	s := New(0, 1, 2, 3, 4, 5, 6, 7, 8, 9)
	u := s.Copy()
	if s.IsDisjoint(u) {
		t.Error("unexpectedly disjoint")
	}
	w := New(10, 11, 12)
	if !u.IsDisjoint(w) {
		t.Error("unexpectedly not disjoint")
	}
}
