package main
 
import "fmt"
 
// struct 정의
type person struct {
    name string
    age  int
}
 
func main() {
    // person 객체 생성
    p := person{}
     
    // 필드값 설정
    p.name = "Lee"
    p.age = 10
     
    fmt.Println(p)
}