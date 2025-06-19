package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	var shared int

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func() {
			shared++ // Вот здесь будет гонка на запись
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println("Shared =", shared)
}

func MaxInt(a, b int) int {
	if a >= b {
		return a
	}

	return b
}
