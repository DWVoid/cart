package conc

import "sync"

type Set[K comparable] struct {
	m sync.Map
}

func (m *Set[K]) Has(key K) bool {
	_, ok := m.m.Load(key)
	return ok
}

func (m *Set[K]) Remove(key K) bool {
	_, was := m.m.LoadAndDelete(key)
	return was
}

func (m *Set[K]) Set(key K) bool {
	_, was := m.m.LoadOrStore(key, true)
	return was
}

func (m *Set[K]) Range(f func(key K) bool) {
	m.m.Range(func(key, value any) bool { return f(key.(K)) })
}
