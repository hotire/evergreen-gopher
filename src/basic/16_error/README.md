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

os.Open() 함수는 func Open(name string) (file *File, err error) 과 같은 함수 원형을 갖는 것으로 첫번째는 File 포인터를 두번째는 error 인터페이스를 리턴한다. 그래서 이경우 두번째 error를 체크해서 nil 이면 에러가 없는 것이고, nil 이 아니면 err.Error() 로부터 해당 에러를 알 수 있다. 아래 예제는 파일을 오픈하는데 에러가 발생하면 에러메시지를 출력하고 빠져나가는 예이다 (주: log.Fatal() 은 메시지를 출력하고 os.Exit(1)을 호출하여 프로그램을 종료한다).

~~~go
package main
 
import (
    "log"
    "os"
)
 
func main() {
    f, err := os.Open("C:\\temp\\1.txt")
    if err != nil {
        log.Fatal(err.Error())
    }
    println(f.Name())
}
~~~