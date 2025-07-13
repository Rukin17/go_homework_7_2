package main

import (
	"fmt"
	"sync"
)

func numGenerator(c chan<- int) {
	defer close(c)
	for i := 1; i <= 10; i++ {
		c <- i
	}

}

func numPrinter(c <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for num := range c {
		fmt.Println(num)
	}
}

func main() {
	nums := make(chan int)
	var wg sync.WaitGroup
	wg.Add(1)
	go numGenerator(nums)
	go numPrinter(nums, &wg)
	wg.Wait()

}
