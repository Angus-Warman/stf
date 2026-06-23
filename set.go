package stf

import "encoding/json"

type Set[T comparable] struct {
	m map[T]struct{}
}

func NewSet[T comparable](values ...T) *Set[T] {
	s := &Set[T]{
		m: make(map[T]struct{}, len(values)),
	}

	for _, v := range values {
		s.m[v] = struct{}{}
	}

	return s
}

func (s *Set[T]) Add(value T) {
	if s.m == nil {
		s.m = make(map[T]struct{})
	}

	s.m[value] = struct{}{}
}

func (s *Set[T]) Remove(value T) {
	delete(s.m, value)
}

func (s *Set[T]) Contains(value T) bool {
	_, ok := s.m[value]

	return ok
}

func (s *Set[T]) Len() int {
	return len(s.m)
}

func (s *Set[T]) Clear() {
	clear(s.m)
}

// Values returns the contents in unspecified order.
func (s *Set[T]) Values() []T {
	values := make([]T, 0, len(s.m))

	for v := range s.m {
		values = append(values, v)
	}

	return values
}

func (s *Set[T]) Clone() *Set[T] {
	clone := &Set[T]{
		m: make(map[T]struct{}, s.Len()),
	}

	for v := range s.m {
		clone.m[v] = struct{}{}
	}

	return clone
}

func (s *Set[T]) Equal(other *Set[T]) bool {
	if s.Len() != other.Len() {
		return false
	}

	for v := range s.m {
		if !other.Contains(v) {
			return false
		}
	}

	return true
}

func (s *Set[T]) AddAll(values ...T) {
	if s.m == nil {
		s.m = make(map[T]struct{}, len(values))
	}

	for _, v := range values {
		s.m[v] = struct{}{}
	}
}

func (s *Set[T]) Pop() (T, bool) {
	for v := range s.m {
		delete(s.m, v)
		return v, true
	}

	var zero T
	return zero, false
}

func (s *Set[T]) Peek() (T, bool) {
	for v := range s.m {
		return v, true
	}

	var zero T
	return zero, false
}

func (s *Set[T]) Union(other *Set[T]) *Set[T] {
	out := s.Clone()

	for k := range other.m {
		out.m[k] = struct{}{}
	}

	return out
}

func (s *Set[T]) Intersection(other *Set[T]) *Set[T] {
	out := NewSet[T]()

	for v := range s.m {
		if other.Contains(v) {
			out.m[v] = struct{}{}
		}
	}

	return out
}

func (s *Set[T]) Difference(other *Set[T]) *Set[T] {
	out := NewSet[T]()

	for v := range s.m {
		if !other.Contains(v) {
			out.m[v] = struct{}{}
		}
	}

	return out
}

func (s *Set[T]) MarshalJSON() ([]byte, error) {
	values := s.Values()

	return json.Marshal(values)
}

func (s *Set[T]) UnmarshalJSON(data []byte) error {
	var values []T

	if err := json.Unmarshal(data, &values); err != nil {
		return err
	}

	s.AddAll(values...)

	return nil
}
