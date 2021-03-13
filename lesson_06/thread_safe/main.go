package main

import (
	"fmt"
	"sync"
)

type mapSafe struct {
	mm map[string]string
	mu sync.RWMutex
}

func main() {

	// trace.Start(os.Stderr)
	// defer trace.Stop()

	s := mapSafe{mm: make(map[string]string, 2), mu: sync.RWMutex{}}

	wg := sync.WaitGroup{}

	wg.Add(2)

	go func() {
		defer wg.Done()
		s.mu.RLock()
		defer s.mu.RUnlock()
		s.mm["1"] = "2"
	}()

	go func() {
		defer wg.Done()
		s.mu.Lock()
		defer s.mu.Unlock()
		s.mm["2"] = "3"
	}()

	wg.Wait()

	s.mu.RLock()
	defer s.mu.RUnlock()
	for k, v := range s.mm {
		fmt.Println(k, v)
	}

}
