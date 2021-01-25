package main

import (
	"fmt"
	"github.com/thoas/go-funk"
)

func main()  {
	result := funk.Contains([]string{"foo", "bar"}, "bar")
	fmt.Println(result)
}