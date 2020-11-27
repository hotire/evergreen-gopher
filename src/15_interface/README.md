# Interface

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