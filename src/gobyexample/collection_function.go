package main

import (
	"fmt"
	"github.com/thoas/go-funk"
)

func main()  {
	fmt.Println(funk.Contains([]string{"foo", "bar"}, "bar"))

	fmt.Println(funk.Find([]int{1, 2, 3, 4}, func(x int) bool {
		return x%2 == 0
	}))
}