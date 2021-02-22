package main

import (
	"os"
)

func main() {
	// 잘못된 파일명을 넣음
	defer println("defer 1")
	println("panic start")
	openFile("Invalid.txt")

	// openFile() 안에서 panic이 실행되면
	// 아래 println 문장은 실행 안됨
	println("panic end")
}

func openFile(fn string) {
	defer println("defer 2")
	f, err := os.Open(fn)
	if err != nil {
		panic(err)
	}
	defer println("defer 3")
	defer f.Close()
}
