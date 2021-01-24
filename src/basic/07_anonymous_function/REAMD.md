# Anonymous Function

http://golang.site/go/article/10-Go-%EC%9D%B5%EB%AA%85%ED%95%A8%EC%88%98

함수명을 갖지 않는 함수를 익명함수(Anonymous Function)이라 부른다. 일반적으로 익명함수는 그 함수 전체를 변수에 할당하거나 다른 함수의 파라미터에 직접 정의되어 사용되곤 한다. 아래 예제는 main() 함수 안에서 익명함수가 선언되어 sum이라는 변수에 할당되고 있음을 보여준다.

~~~go
package main
 
func main() {
    sum := func(n ...int) int { //익명함수 정의
        s := 0
        for _, i := range n {
            s += i
        }
        return s
    }
 
    result := sum(1, 2, 3, 4, 5) //익명함수 호출
    println(result)
}
~~~

익명함수가 변수에 할당된 이후에는 변수명이 함수명과 같이 취급되며 "변수명(파라미터들)" 형식으로 함수를 호출할 수 있다.


###  일급함수

Go 프로그래밍 언어에서 함수는 일급함수로서 Go의 기본 타입과 동일하게 취급되며, 따라서 다른 함수의 파라미터로 전달하거나 다른 함수의 리턴값으로도 사용될 수 있다. 즉, 함수의 입력 파라미터나 리턴 파라미터로서 함수 자체가 사용될 수 있다. 

### type문을 사용한 함수 원형 정의

type 문은 구조체(struct), 인터페이스 등 Custom Type(혹은 User Defined Type)을 정의하기 위해 사용된다. type 문은 또한 함수 원형을 정의하는데 사용될 수 있다. 즉, 위 예제에서 func(x int, y int) int 함수 원형이 코드 상에 계속 반복됨을 볼 수 있는데, 이 경우 type 문을 정의함으로써 해당 함수의 원형을 간단히 표현할 수 있다.


~~~go
// 원형 정의
type calculator func(int, int) int
 
// calculator 원형 사용
func calc(f calculator, a int, b int) int {
    result := f(a, b)
    return result
}
~~~