# Error
http://golang.site/go/article/19-Go-%EC%97%90%EB%9F%AC%EC%B2%98%EB%A6%AC

Go는 내장 타입으로 error 라는 interface 타입을 갖는다. Go 에러는 이 error 인터페이스를 통해서 주고 받게 되는데, 이 interface는 다음과 같은 하나의 메서드를 갖는다. 개발자는 이 인터페이스를 구현하는 커스텀 에러 타입을 만들 수 있다.

~~~go
type error interface {
    Error() string
}
~~~

### Go 에러처리

Go 함수가 결과와 에러를 함께 리턴한다면, 이 에러가 nil 인지를 체크해서 에러가 없는지를 체크할 수 있다