package main

import (
	"encoding/json"
	"fmt"
)

type Response struct {
	Status int
	Body string
}

func main() {
	bolB, _ := json.Marshal(true)
	fmt.Println(string(bolB))
}