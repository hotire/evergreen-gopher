package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	url := "https://google.com/robots.txt"

	resp, _ := http.Get(url)
	robots, _ := ioutil.ReadAll(resp.Body)
	resp.Body.Close()

	fmt.Printf("%s\n", robots)
}
