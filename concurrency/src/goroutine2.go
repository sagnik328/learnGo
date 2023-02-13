package main

import (
	"fmt"
	"time"
)

//goroutines are light weight threads
// take little memory and all managerd by the go scheduler.

func printSomething(s string) {
	fmt.Println(s)
}

func main() {
	//this spawns its own go routine
	//go printSomething("This is the first thing to be printed")
	//printSomething("This is the second thing to be printed")
	words := []string{
		"alpha",
		"beta",
		"gamma",
		"delta",
		"epsilon",
		"pi",
	}

	for i, word := range words {
		//they will not be in order
		// it doesnt matter in what order we spanwed them
		//better to use waitgroups
		go printSomething(fmt.Sprintf("%d: %s", i, word))
	}
	time.Sleep(1 * time.Second)
	printSomething("This is the third thing to be printed")
}
