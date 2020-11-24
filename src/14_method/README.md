# Method

http://golang.site/go/article/17-Go-%EB%A9%94%EC%84%9C%EB%93%9C

앞에서(Go 구조체) 언급했듯이 Go 언어는 객체지향 프로그래밍(OOP)을 고유의 방식으로 지원한다. 타 언어의 OOP의 클래스가 필드와 메서드를 함께 갖는 것과 달리 Go 언어에서는 struct가 필드만을 가지며, 메서드는 별도로 분리되어 정의된다.

Go 메서드는 특별한 형태의 func 함수이다. 메서드는 함수 정의에서 func 키워드와 함수명 사이에 "그 함수가 어떤 struct를 위한 메서드인지"를 표시하게 된다. 흔히 receiver로 불리우는 이 부분은 메서드가 속한 struct 타입과 struct 변수명을 지정하는데, struct 변수명은 함수 내에서 마치 입력 파라미터처럼 사용된다. 

~~~go
package main
 
//Rect - struct 정의
type Rect struct {
    width, height int
}
 
//Rect의 area() 메소드
func (r Rect) area() int {
    return r.width * r.height   
}
 
func main() {
    rect := Rect{10, 20}
    area := rect.area() //메서드 호출
    println(area)
}
~~~


### Value vs 포인터 receiver

Value receiver는 struct의 데이타를 복사(copy)하여 전달하며, 포인터 receiver는 struct의 포인터만을 전달한다. Value receiver의 경우 만약 메서드내에서 그 struct의 필드값이 변경되더라도 호출자의 데이타는 변경되지 않는 반면, 포인터 receiver는 메서드 내의 필드값 변경이 그대로 호출자에서 반영된다.


### Go 인터페이스

구조체(struct)가 필드들의 집합체라면, interface는 메서드들의 집합체이다. interface는 타입(type)이 구현해야 하는 메서드 원형(prototype)들을 정의한다. 하나의 사용자 정의 타입이 interface를 구현하기 위해서는 단순히 그 인터페이스가 갖는 모든 메서드들을 구현하면 된다.

인터페이스는 struct와 마찬가지로 type 문을 사용하여 정의한다.