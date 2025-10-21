package main

import (
	"fmt"
	"sync"
)

func foo() {
	panic("foo")
}

func bar(i int) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in bar:", r, "i:", i)
		}
	}()
	foo()
}

func main() {
	var wg sync.WaitGroup
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			fmt.Println("goroutine start", i)
			bar(i)
			i++

		}()
	}
	wg.Add(1)

	go func() {
		defer wg.Done()
		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main:", r)
			}
		}()
		foo()
	}()
	wg.Wait()
}
