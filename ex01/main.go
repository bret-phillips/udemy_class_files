package main

import "fmt"

func main() {
	var count = 1
	for i := 1; i < 5; i++ {
		count *= i
	}
	fmt.Println(count)
}
