package main

import "fmt"

func main() {
	for i := 0; i <= 10; i++ {
		if isEven(i) {
			fmt.Println(i, " is even")
			continue
		}
		fmt.Println(i, " is odd")
	}
}

func isEven(i int) bool {
	return i%2 == 0
}
