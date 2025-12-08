package x

import "sync"

type typeSafeMap[T any] struct {
	Data map[string]T

	mu sync.RWMutex
}

func newtypeSafeMap[T any]() *typeSafeMap[T] { return &typeSafeMap[T]{Data: make(map[string]T)} }

func (s *typeSafeMap[T]) Get(key string) (T, bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	v, ok := s.Data[key]

	return v, ok
}

func (s *typeSafeMap[T]) Set(key string, value T) {
	s.mu.Lock()
	s.Data[key] = value
	s.mu.Unlock()
}

func (s *typeSafeMap[T]) Delete(key string) {
	s.mu.Lock()
	delete(s.Data, key)
	s.mu.Unlock()
}

func (s *typeSafeMap[T]) Clear() {
	s.mu.Lock()
	s.Data = make(map[string]T)
	s.mu.Unlock()
}

func (s *typeSafeMap[T]) Keys() []string {
	keys := make([]string, 0, len(s.Data))
	for k := range s.Data {
		keys = append(keys, k)
	}

	return keys
}

type TypeSafeMap[T any] struct{ *typeSafeMap[T] }

func NewTypeSafeMap[T any]() *TypeSafeMap[T] {
	return &TypeSafeMap[T]{typeSafeMap: newtypeSafeMap[T]()}
}

type SafeMap struct{ *typeSafeMap[any] }

func NewSafeMap() *SafeMap { return &SafeMap{typeSafeMap: newtypeSafeMap[any]()} }

func GetTypedFromLockedMap[T any](s *SafeMap, key string) (T, bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	v, ok := s.Data[key]
	if !ok {
		var z T
		return z, ok
	}

	value, ok := v.(T)

	return value, ok
}
