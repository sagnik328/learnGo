package main

import (
	"fmt"
	"sync"
)

func printSomething(s string, wg *sync.WaitGroup) {
	defer wg.Done() //decrement eg by 1 after exiting the function
	fmt.Println(s)
}

func main() {
	var wg sync.WaitGroup
	words := []string{
		"alpha",
		"beta",
		"gamma",
		"delta",
		"epsilon",
		"pi",
	}
	wg.Add(len(words))
	//once we created a wait group we shouldnt copy them
	for i, word := range words {
		go printSomething(fmt.Sprintf("%d: %s", i, word), &wg)
	}
	wg.Wait() //will wait till the wg is set to 0
	wg.Add(1)
	printSomething("This is the second thing to be printed", &wg)
}
