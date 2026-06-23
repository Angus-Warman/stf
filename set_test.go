package stf

import (
	"encoding/json"
	"slices"
	"testing"
)

func TestNewSet(t *testing.T) {
	s := NewSet(1, 2, 2, 3)

	assertEqual(t, 3, s.Len())
}

func TestSetContains(t *testing.T) {
	s := NewSet[int]()

	if s.Contains(123) {
		t.Fatal("unexpected value")
	}

	s.Add(123)

	if !s.Contains(123) {
		t.Error("expected value")
	}
}

func TestRemove(t *testing.T) {
	s := NewSet(1, 2, 3)

	s.Remove(2)

	if s.Contains(2) {
		t.Fatal("value should have been removed")
	}

	assertEqual(t, 2, s.Len())
}

func TestSetDuplicateAdd(t *testing.T) {
	s := NewSet[int]()

	s.Add(5)
	s.Add(5)
	s.Add(5)

	assertEqual(t, 1, s.Len())
}

func TestSetClear(t *testing.T) {
	s := NewSet(1, 2, 3)

	s.Clear()

	assertEqual(t, 0, s.Len())

	if s.Contains(1) {
		t.Fatal("set should be empty")
	}
}

func TestSetValues(t *testing.T) {
	s := NewSet(3, 1, 2)

	values := s.Values()
	slices.Sort(values)

	expected := []int{1, 2, 3}

	if !slices.Equal(values, expected) {
		t.Fatalf("expected %v, got %v", expected, values)
	}
}

func TestSetClone(t *testing.T) {
	a := NewSet(1, 2, 3)
	b := a.Clone()

	if !a.Equal(b) {
		t.Errorf("expected sets to be equal")
	}

	b.Add(4)

	if a.Contains(4) {
		t.Fatal("clone should be independent")
	}
}

func TestSetEqual(t *testing.T) {
	a := NewSet(1, 2, 3)
	b := NewSet(3, 2, 1)
	c := NewSet(1, 2)

	if !a.Equal(b) {
		t.Fatal("sets should be equal")
	}

	if a.Equal(c) {
		t.Fatal("sets should not be equal")
	}
}

func TestSetZeroValue(t *testing.T) {
	var s Set[string]

	s.Add("hello")

	if !s.Contains("hello") {
		t.Fatal("zero-value set should be usable")
	}

	assertEqual(t, 1, s.Len())
}

func TestSetRemoveMissing(t *testing.T) {
	s := NewSet(1, 2)

	s.Remove(999)

	assertEqual(t, 2, s.Len())
}

func TestSetIntersection(t *testing.T) {
	a := NewSet(1, 2, 3)
	b := NewSet(2, 3, 4)

	c := a.Intersection(b)

	assertEqual(t, 2, c.Len())
}

func TestSetUnion(t *testing.T) {
	a := NewSet(1, 2, 3)
	b := NewSet(2, 3, 4)

	c := a.Union(b)

	assertEqual(t, 4, c.Len())
}

func TestSetDifference(t *testing.T) {
	setXYZ := NewSet("x", "y", "z")
	setXYQ := NewSet("x", "y", "q")

	setZ := setXYZ.Difference(setXYQ)
	setQ := setXYQ.Difference(setXYZ)

	assertEqual(t, 1, setZ.Len())
	assertEqual(t, 1, setQ.Len())

	zValue, _ := setZ.Pop()
	qValue, _ := setQ.Pop()

	assertEqual(t, zValue, "z")
	assertEqual(t, qValue, "q")
}

func TestSetJSONRoundTrip(t *testing.T) {
	original := NewSet(1, 2, 3)

	data, err := json.Marshal(original)
	if err != nil {
		t.Fatalf("Marshal failed: %v", err)
	}

	var decoded Set[int]
	if err := json.Unmarshal(data, &decoded); err != nil {
		t.Fatalf("Unmarshal failed: %v", err)
	}

	if !original.Equal(&decoded) {
		t.Fatalf("expected %v, got %v", original.Values(), decoded.Values())
	}
}

func assertEqual[V comparable](t *testing.T, expected, got V) {
	t.Helper()

	if expected != got {
		t.Errorf("expected %v, got %v", expected, got)
	}
}
