package main

import (
	"fmt"
	"sync"
)

func main() {
	var waitGroup sync.WaitGroup
	nCount := 2
	waitGroup.Add(nCount)

	for i := 0; i < nCount; i++ {
		go func(wg *sync.WaitGroup) {
			defer wg.Done()
			fmt.Println("work: finish")
		}(&waitGroup)
	}

	fmt.Println("main:begin wait")
	waitGroup.Wait()
	fmt.Println("main:after wait")
}
