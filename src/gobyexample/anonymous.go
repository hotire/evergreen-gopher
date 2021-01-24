package main

import "fmt"

func main() {

	sum := func(n ...int) int {
		result := 0
		for _, i := range n {
			result += i
		}
		return result
	}

	fmt.Print(sum(4))
}