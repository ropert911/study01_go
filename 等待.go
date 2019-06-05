package main

import (
	"fmt"
	"sync"
)

func main() {
	var waitGroup sync.WaitGroup
	n_count := 2
	waitGroup.Add(n_count)

	for i := 0; i < n_count; i++ {
		go func(wg *sync.WaitGroup) {
			defer wg.Done()
			fmt.Println("work: finish")
		}(&waitGroup)
	}

	fmt.Println("main:begin wait")
	waitGroup.Wait()
	fmt.Println("main:after wait")
}
