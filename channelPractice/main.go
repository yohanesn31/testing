package main

import (
	"fmt"
	"sync"
)

type a struct {
	a int
}

func main() {

	input := genNumber(100)
	c0 := factorial(input)
	c1 := factorial(input)
	c2 := factorial(input)
	c3 := factorial(input)
	c4 := factorial(input)
	c5 := factorial(input)
	c6 := factorial(input)
	c7 := factorial(input)
	c8 := factorial(input)
	c9 := factorial(input)
	// for n := range factorial(genNumber(10)) {
	// 	fmt.Println("Total:", n)
	// }
	for n := range merge(c0, c1, c2, c3, c4, c5, c6, c7, c8, c9) {
		fmt.Println("Total: ", n)
	}

}

func merge(n ...<-chan int) <-chan int {
	var wg sync.WaitGroup
	out := make(chan int)
	wg.Add(len(n))
	output := func(val <-chan int) {
		for c := range val {
			out <- c
		}
		wg.Done()
	}
	for _, val := range n {
		go output(val)
	}
	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}

func genNumber(n int) <-chan int {
	out := make(chan int)
	go func() {
		for i := n; i > 0; i-- {
			for j := 3; j < 13; j++ {
				out <- j
			}
		}
		close(out)
	}()
	return out
}

func factorial(n <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for a := range n {
			total := 1
			for i := a; i > 0; i-- {
				total *= i
			}
			out <- total
		}
		close(out)
	}()
	return out
}
