# 변수와 상수 

http://golang.site/go/article/4-Go-%EB%B3%80%EC%88%98%EC%99%80-%EC%83%81%EC%88%98

## 변수 

변수는 Go 키워드 var 를 사용하여 선언한다. var 키워드 뒤에 변수명을 적고, 그 뒤에 변수타입을 적는다

~~~go
var a int  = 1
~~~

변수(정수) a 를 선언 및 초기화(1) 또한 초기화 값을 통해 타입 추론이 가능하다. 

### Short Assignment Statement 

변수를 선언하는 또 다른 방식으로 Short Assignment Statement ( := ) 를 사용할 수 있다. 즉, var i = 1 을 쓰는 대신 i := 1 이라고 var 를 생략하고 사용할 수 있다. 하지만 이러한 표현은 함수(func) 내에서만 사용할 수 있으며, 함수 밖에서는 var를 사용해야 한다. Go에서 변수와 상수는 함수 밖에서도 사용할 수 있다.

## 상수 

상수는 Go 키워드 const 를 사용하여 선언한다. const 키워드 뒤에 상수명을 적고, 그 뒤에 상수타입, 그리고 상수값을 할당한다.

~~~go
const age int = 30
const name string = "hotire"

const (
    Visa = "Visa"
    Master = "MasterCard"
    Amex = "American Express"
)
~~~
이렇게 상수를 묶어서 지정도 가능하다. 

~~~go
const (
    Apple = iota // 0
    Grape        // 1
    Orange       // 2
)
~~~
iota를 사용하면 순차적으로 증가된 값을 사용할 수 있다. (enum ordinal 같은 느낌..??)



