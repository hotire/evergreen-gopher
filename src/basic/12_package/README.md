## Go 패키지

http://golang.site/go/article/15-Go-%ED%8C%A8%ED%82%A4%EC%A7%80

Go는 패키지(Package)를 통해 코드의 모듈화, 코드의 재사용 기능을 제공한다. Go는 패키지를 사용해서 작은 단위의 컴포넌트를 작성하고, 이러한 작은 패키지들을 활용해서 프로그램을 작성할 것을 권장한다.

표준 라이브러리 패키지들은 GOROOT/pkg 안에 존재한다. GOROOT 환경변수는 Go 설치 디렉토리를 가리키는데, 보통 Go 설치시 자동으로 추가된다. 즉, 윈도우즈에서 Go를 설치했을 경우 디폴트로 C:\go 에 설치되며, GOROOT는 C:\go를 가리킨다.

Go에 사용하는 표준패키지는 https://golang.org/pkg 에 자세히 설명되어 있다. Go 프로그래밍에서 표준패키지를 자주 불러 사용하므로 이 링크를 자주 참조하게 될 것이다.

## Main 패키지 

일반적으로 패키지는 라이브러리로서 사용되지만, "main" 이라고 명명된 패키지는 Go Compiler에 의해 특별하게 인식된다. 패키지명이 main 인 경우, 
컴파일러는 해당 패키지를 공유 라이브러리가 아닌 실행(executable) 프로그램으로 만든다. 
main 패키지 안의 main() 함수가 프로그램의 시작점 즉 Entry Point가 된다. 패키지를 공유 라이브러리로 만들 때에는, main 패키지나 main 함수를 사용해서는 안된다.

## 패키지 Import

다른 패키지를 프로그램에서 사용하기 위해서는 import 를 사용하여 패키지를 포함시킨다. 예를 들어, Go의 표준 라이브러리인 fmt 패키지를 사용하기 위하여, import "fmt" 와 같이 해당 패키지를 포함시킬 것을 선언해 준다. Import 후에는 아래 예제처럼 fmt 패키지의 Println() 함수를 호출하여 사용할 수 있다.

~~~go
package main
 
import "fmt"
 
func main(){
 fmt.Println("Hello")
}
~~~


## 패키지 Scope

패키지 내에는 함수, 구조체, 인터페이스, 메서드 등이 존재하는데, 이들의 이름(Identifier)이 첫문자를 대문자로 시작하면 이는 public 으로 사용할 수 있다. 즉, 패키지 외부에서 이들을 호출하거나 사용할 수 있게 된다. 반면, 이름이 소문자로 시작하면 이는 non-public 으로 패키지 내부에서만 사용될 수 있다.


## 패키지 init 함수와 alias
개발자가 패키지를 작성할 때, 패키지 실행시 처음으로 호출되는 init() 함수를 작성할 수 있다. 즉, init 함수는 패키지가 로드되면서 실행되는 함수로 별도의 호출 없이 자동으로 호출된다.

~~~go
package testlib
 
var pop map[string]string
 
func init() {   // 패키지 로드시 map 초기화
    pop = make(map[string]string)
}
~~~

경우에 따라 패키지를 import 하면서 단지 그 패키지 안의 init() 함수만을 호출하고자 하는 케이스가 있다. 
이런 경우는 패키지 import 시 _ 라는 alias 를 지정한다. 아래는 other/xlib 패키지를 호출하면서 _ alias를 지정한 예이다.

~~~go
package main
import _ "other/xlib"
~~~

## 사용자 정의 패키지 생성
개발자는 사용자 정의 패키지를 만들어 재사용 가능한 컴포넌트를 만들어 사용할 수 있다. 사용자 정의 라이브러리 패키지는 일반적으로 폴더를 하나 만들고 그 폴더 안에 .go 파일들을 만들어 구성한다. 하나의 서브 폴더안에 있는 .go 파일들은 동일한 패키지명을 가지며, 패키지명은 해당 폴더의 이름과 같게 한다. 즉, 해당 폴더에 있는 여러 *.go 파일들은 하나의 패키지로 묶인다.

~~~go
package testlib
 
import "fmt"
 
var pop map[string]string
 
func init() {
    pop = make(map[string]string)
    pop["Adele"] = "Hello"
    pop["Alicia Keys"] = "Fallin'"
    pop["John Legend"] = "All of Me"
}
 
// GetMusic : Popular music by singer (외부에서 호출 가능)
func GetMusic(singer string) string {
    return pop[singer]
}
 
// 내부에서만 호출 가능
func getKeys() {  
    for _, kv := range pop {
        fmt.Println(kv)
    }
}
~~~

- 예제로 간단한 패키지를 만들기 위해 /src 폴더 안에 (임의의 폴더명으로) 24lab.net/testlib 폴더를 생성한 후에 다음 코드를 music.go 라는 파일에 저장한다. 여기서 패키지명은 폴더명과 동일하게 testlib로 정해주어야 한다. 패키지 폴더 안에 여러 파일들이 있을 경우에도, 동일하게 testlib 패키지명을 사용한다.

Optional : 사이즈가 큰 복잡한 라이브러리 같은 경우, "go install" 명령을 사용하여 라이브러리를 컴파일하여 Cache할 수 있는데, 이렇게 하면 다음 빌드시 빌드타임을 크게 줄일 수 있다. Go 패키지를 빌드하고 /pkg 폴더에 인스톨하기 위해서 "go install" 명령어를 아래 그림과 같이 testlib 폴더안에서 실행할 수 있다. 이 명령어가 실행되면, testlib.a 라는 파일이 /pkg/windows_amd64/24lab.net 안에 생성된다 (주: 여기서 windows_amd64는 머신에 따라 변경된다. 만약 main 패키지에 go install 명령을 수행하면 /bin 폴더에 실행파일을 생성한다.).


24lab.net/testlib 패키지가 있는 위치를 찾기 위해 GOROOT와 GOPATH의 경로를 사용하는데, GOROOT와 GOPATH에 있는 각 루트폴더의 src 밑에 24lab.net/testlib 폴더를 순서대로 찾게 된다. 즉, GOPATH가 C:\GoApp;C:\GoSrc 인경우, 지정된 라이브러리를 찾기 위해 다음과 같은 폴더를 순차적으로 검색하게 된다.