package main

import "fmt"

func main() {
	i := 1
	for i <= 3 {
		fmt.Println(i)
		i++
	}

	for j := 1; j <= i; j++ {
		fmt.Println(j)
	}

	for {
		fmt.Println("break")
		break
	}
}
