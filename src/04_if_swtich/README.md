# If && Switch

## If

if 문에서 조건식을 사용하기 이전에 간단한 문장(Optional Statement)을 함께 실행할 수 있다. 

~~~go
if val := i * 2; val < max {
    println(val)
}

val++ // Scope error
~~~
- 변수 val는 if문 블럭 (혹은 if-else 블럭 scope) 안에서만 사용할 수 있다는 것이다.


## Switch

### switch 뒤에 expression이 없을 수 있음

다른 언어는 switch 키워드 뒤에 변수나 expression 반드시 두지만, Go는 이를 쓰지 않아도 된다. 이 경우 Go는 switch expression을 true로 생각하고 첫번째 case문으로 이동하여 검사한다

~~~go
func grade(score int) {
    switch {
    case score >= 90:
        println("A")
    case score >= 80:
        println("B")
    case score >= 70:
        println("C")
    case score >= 60:
        println("D")
    default:
        println("No Hope")
    }
}  
~~~

swtich 뒤에 조건변수 / Expression를 적지 않는다. if / else if / else 처럼 사용하는 방식이다. 