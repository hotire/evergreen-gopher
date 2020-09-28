# 함수 

함수는 여러 문장을 묶어서 실행하는 코드 블럭의 단위이다. Go에서 함수는 func 키워드를 사용하여 정의한다. func 뒤에 함수명을 적고 괄호 ( ) 안에 그 함수에 전달하는 파라미터들을 적게 된다. 함수 파라미터는 0개 이상 사용할 수 있는데, 각 파라미터는 파라미터명 뒤에 int, string 등의 파라미터 타입을 적어서 정의한다. 함수의 리턴 타입은 파라미터 괄호 ( ) 뒤에 적는다.

~~~go 
package main
func main() {
    msg := "Hello"
    say(msg)
}
 
func say(msg string) {
    println(msg)
}
~~~~
