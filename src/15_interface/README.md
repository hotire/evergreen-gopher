# Interface

http://golang.site/go/article/18-Go-%EC%9D%B8%ED%84%B0%ED%8E%98%EC%9D%B4%EC%8A%A4

구조체(struct)가 필드들의 집합체라면, interface는 메서드들의 집합체이다. interface는 타입(type)이 구현해야 하는 메서드 원형(prototype)들을 정의한다. 하나의 사용자 정의 타입이 interface를 구현하기 위해서는 단순히 그 인터페이스가 갖는 모든 메서드들을 구현하면 된다.

인터페이스는 struct와 마찬가지로 type 문을 사용하여 정의한다.

~~~go
type Shape interface {
    area() float64
    perimeter() float64
}
~~~

### 인터페이스 구현

인터페이스를 구현하기 위해서는 해당 타입이 그 인터페이스의 메서드들을 모두 구현하면 되므로, 위의 Shape 인터페이스를 구현하기 위해서는 area(), perimeter() 2개의 메서드만 구현하면 된다. 예를 들어 Rect와 Circle 이라는 2개의 타입이 있을 때, Shape 인터페이스를 구현하기 위해서는 아래와 같이 각 타입별로 2개의 메서드를 구현해 주면 된다.

~~~go
//Rect 정의
type Rect struct {
    width, height float64
}
 
//Circle 정의
type Circle struct {
    radius float64
}
 
//Rect 타입에 대한 Shape 인터페이스 구현 
func (r Rect) area() float64 { return r.width * r.height }
func (r Rect) perimeter() float64 {
     return 2 * (r.width + r.height)
}
 
//Circle 타입에 대한 Shape 인터페이스 구현 
func (c Circle) area() float64 { 
    return math.Pi * c.radius * c.radius
}
func (c Circle) perimeter() float64 { 
    return 2 * math.Pi * c.radius
}
~~~

### 인터페이스 사용 

인터페이스를 사용하는 일반적인 예로 함수가 파라미터로 인터페이스를 받아들이는 경우를 들 수 있다. 함수 파라미터가 interface 인 경우, 이는 어떤 타입이든 해당 인터페이스를 구현하기만 하면 모두 입력 파라미터로 사용될 수 있다는 것을 의미한다. 
- 너무 당연한 이야기??

아래 예제에서 showArea() 함수는 Shape 인터페이스들을 파라미터로 받아들이고 있는데, 따라서 Rect와 Circle 처럼 Shape 인터페이스를 구현한 타입 객체들을 파라미터로 받을 수 있다. showArea() 함수 내에서 해당 인터페이스가 가진 메서드 즉 area() 혹은 perimeter()을 사용할 수 있다.

~~~go
func main() {
    r := Rect{10., 20.}
    c := Circle{10}
 
    showArea(r, c)
}
 
func showArea(shapes ...Shape) {
    for _, s := range shapes {
        a := s.area() //인터페이스 메서드 호출
        println(a)
    }
~~~

### 인터페이스 타입 

Go 프로그래밍을 하다보면 흔히 빈 인터페이스(empty interface)를 자주 접하게 되는데, 흔히 인터페이스 타입(interface type)으로도 불리운다. 예를 들어, 여러 표준패키지들의 함수 Prototype을 살펴보면, 아래와 같이 빈 interface 가 자주 등장함을 볼 수 있다. 빈 interface는 interface{} 와 같이 표현한다.

~~~go
func Marshal(v interface{}) ([]byte, error);
func Println(a ...interface{}) (n int, err error);
~~~
Empty interface는 메서드를 전혀 갖지 않는 빈 인터페이스로서, Go의 모든 Type은 적어도 0개의 메서드를 구현하므로, 흔히 Go에서 모든 Type을 나타내기 위해 빈 인터페이스를 사용한다. 

즉, 빈 인터페이스는 어떠한 타입도 담을 수 있는 컨테이너라고 볼 수 있으며, 여러 다른 언어에서 흔히 일컫는 Dynamic Type 이라고 볼 수 있다. (주: empty interface는 C#, Java 에서 object라 볼 수 있으며, C/C++ 에서는 void* 와 같다고 볼 수 있다)

-> Java Object를 굉장히 어렵게 설명하네... / Kotlin Any