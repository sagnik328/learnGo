package main

import "fmt"

func main() {
	sum := 0
	for i := 0; i < 10; i++ {
		sum++
	}
	fmt.Println(sum)
	for sum < 1000 {
		sum += 1
		fmt.Println("ook")
	}
	fmt.Println(sum)
}
