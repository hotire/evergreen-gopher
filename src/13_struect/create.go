package main
 
import "fmt"
 
// struct 정의
type person struct {
    name string
    age  int
}
 
func main() {
	var p1 person 
	p1 = person{"Bob", 20}
	p2 := person{name: "Sean", age: 50}
}