package main

import "fmt"

func main() {
	var count = 1
	var num2 = 3
	for i := 1; i < 5; i++ {
		count *= i
		num2 *= i
	}
	fmt.Println(count, num2)
}
