package main

import (
	"fmt"
	"sync"
)

func main() {
	var m sync.Map
	var wg sync.WaitGroup

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			m.Store("key", i)
		}(i)
	}

	wg.Wait()

	val, _ := m.Load("key")
	fmt.Println("sync.Map result:", val)

	normalMap := make(map[string]int)
	var mu sync.RWMutex

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			mu.Lock()
			normalMap["key"] = i
			mu.Unlock()
		}(i)
	}

	wg.Wait()

	fmt.Println("RWMutex result:", normalMap["key"])
}
