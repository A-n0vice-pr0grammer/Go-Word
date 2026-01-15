package main

import (
	"fmt"
	"sync"
	"time"
)

var Wait sync.WaitGroup

func Shop(num string) {
	fmt.Println("xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx")
	fmt.Println(num, "口开始结账")

	Ch := make(chan int, 3)
	All := 0
	Ch <- 5
	Sum := <-Ch

	All = All + Sum
	fmt.Println(num, "口结账金额为", All)
	fmt.Println("xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx")
	Wait.Done()
}

func main() {

	Wait.Add(3)

	go Shop("a")
	time.Sleep(2 * time.Second)
	go Shop("b")
	time.Sleep(2 * time.Second)
	go Shop("c")
	Wait.Wait()
}
