package main

import (
	"fmt"
	"sync"
	"time"
	"runtime"
)

var wg sync.WaitGroup

func init(){
	runtime.GOMAXPROCS(runtime.NumCPU() ) 
}

func foo(){
	for i := 0 ; i < 45 ; i++ {
		fmt.Printf("Foo: %d\n", i)
		time.Sleep(time.Duration(3 * time.Millisecond))
	}
	wg.Done()
}

func bar(){
	for i := 0 ; i < 45 ; i++ {
		fmt.Printf("Bar: %d\n", i)
		time.Sleep(time.Duration(20 * time.Millisecond))
	}
	wg.Done()
}

func main(){
	wg.Add(2)
	go foo()
	go bar()
	wg.Wait()
}