package main

import (
	"fmt"
	"sync"
)

type SafeMap struct {
	mu sync.RWMutex
	m  map[string]int
}

func NewSafeMap() *SafeMap {
	return &SafeMap{m: make(map[string]int)}
}

func (s *SafeMap) Set(k string, v int) {
	s.mu.Lock()
	s.m[k] = v
	s.mu.Unlock()
}

func (s *SafeMap) Get(k string) (int, bool) {
	s.mu.RLock()
	v, ok := s.m[k]
	s.mu.RUnlock()
	return v, ok
}

func (s *SafeMap) Len() int {
	s.mu.RLock()
	n := len(s.m)
	s.mu.RUnlock()
	return n
}

func main() {
	sm := NewSafeMap()

	var wg sync.WaitGroup
	const workers = 8
	const perWorker = 50000

	wg.Add(workers)
	for w := 0; w < workers; w++ {
		go func(workerID int) {
			defer wg.Done()
			start := workerID * perWorker
			for i := start; i < start+perWorker; i++ {
				// разные ключи, чтобы не спорить за один и тот же
				key := fmt.Sprintf("k_%d", i)
				sm.Set(key, i)
			}
		}(w)
	}
	wg.Wait()

	fmt.Println("len:", sm.Len())
	v, ok := sm.Get("k_123")
	fmt.Println("get k_123:", v, ok)
}
