package main

import "fmt"

func values(arg int) (int, int) {
	return arg, arg + 1
}

func main() {
	fmt.Println(values(3))
}
