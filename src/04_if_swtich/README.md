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