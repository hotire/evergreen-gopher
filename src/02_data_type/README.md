# 데이터 타입 

Go 프로그래밍 언어는 다음과 같은 기본적인 데이타 타입들을 갖고 있다.

- bool

- string

- int int8 int16 int32 int64

- float32 float64 complex64 complex128

- byte / rune


## 문자열 

문자열 리터럴은 Back Quote(` `) 혹은 이중인용부호(" ")를 사용하여 표현할 수 있다.

- Back Quote (` `)로 둘러 싸인 문자열은 Raw String Literal이라 부르는데, 이 안에 있는 문자열은 별도로 해석되지 않고 Raw String 그대로의 값을 갖는다. 예를 들어, 문자열 안에 \n 이 있을 경우 이는 NewLine으로 해석되지 않는다. 또한, Back Quote은 복수 라인의 문자열을 표현할 때 자주 사용된다.
- 이중인용부호(" ")로 둘러 싸인 문자열은 Interpreted String Literal이라 부르는데, 복수 라인에 걸쳐 쓸 수 없으며, 인용부호 안의 Escape 문자열들은 특별한 의미로 해석된다. 예를 들어, 문자열 안에 \n 이 있을 경우 이는 NewLine으로 해석된다. 이중인용부호를 이용해 문자열을 여러 라인에 걸쳐 쓰기 위해서는 + 연산자를 이용해 결합하여 사용한다.


## 데이타 타입 변환 (Type Conversion)

하나의 데이타 타입에서 다른 데이타 타입으로 변환하기 위해서는 T(v) 와 같이 표현하고 이를 Type Conversion이라 부르는데, 여기서 T는 변환하고자 하는 타입을 표시하고, v는 변환될 값(value)을 지정한 것이다.

~~~go
func main() {
    var i int = 100
    var u uint = uint(i)
    var f float32 = float32(i)  
    println(f, u)
 
    str := "ABC"
    bytes := []byte(str)
    str2 := string(bytes)
    println(bytes, str2)
}
~~~