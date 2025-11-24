package state

import "sync"

var Shared *State

type State struct {
	Data map[string]any
	m    sync.RWMutex
}

func NewState() *State {
	return &State{Data: make(map[string]any)}
}

func (s *State) Get(key string) (any, bool) {
	s.m.RLock()
	defer s.m.RUnlock()
	v, ok := s.Data[key]

	return v, ok
}

func (s *State) Set(key string, value any) {
	s.m.Lock()
	s.Data[key] = value
	s.m.Unlock()
}

func (s *State) Delete(key string) {
	s.m.Lock()
	delete(s.Data, key)
	s.m.Unlock()
}

func (s *State) Reset() {
	s.m.Lock()
	s.Data = make(map[string]any)
	s.m.Unlock()
}

func (s *State) Keys() []string {
	keys := make([]string, 0, len(s.Data))
	for k := range s.Data {
		keys = append(keys, k)
	}

	return keys
}

func Get[T any](s *State, key string) (value T, ok bool) {
	s.m.RLock()
	defer s.m.RUnlock()

	v, ok := s.Data[key]
	if !ok {
		var z T
		return z, ok
	}

	value, ok = v.(T)

	return value, ok
}

func init() {
	Shared = &State{Data: make(map[string]any)}
}
