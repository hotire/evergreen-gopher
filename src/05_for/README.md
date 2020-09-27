# For

Go는 반복문에 for 하나 밖에 없다. 다만, "초기값; 조건식; 증감"을 둘러싸는 괄호 ( )를 생략하는데, 괄호를 쓰면 에러가 난다.

http://golang.site/go/article/8-Go-%EB%B0%98%EB%B3%B5%EB%AC%B8

### 조건식만 쓰는 for 루프

Go에서 for 루프는 초기값과 증감식을 생략하고 조건식만을 사용할 수 있는데, 이는 마치 for 루프가 다른 언어의 while 루프와 같이 쓰이도록 한다

~~~go
package main

func main() {
    n := 1
    for n < 100 {
        n *= 2      
        //if n > 90 {
        //   break 
        //}     
    }
    println(n)
}
~~~

### for 문 - 무한루프

for 루프로 무한루프를 만들려면 "초기값; 조건식; 증감" 모두를 생략하면 된다

~~~go
package main
 
func main() {
    for {
        println("Infinite loop")        
    }
}
~~~
* 빠져나오기 위해서 Ctrl+C 


### for range 문

for range 문은 컬렉션으로 부터 한 요소(element)씩 가져와 차례로 for 블럭의 문장들을 실행한다. 이는 다른 언어의 foreach와 비슷한 용법이다.

~~~go
names := []string{"홍길동", "이순신", "강감찬"}
for index, name := range names {
    println(index, name)
}
~~~

### break, continue, goto 문

break문은 보통 단독으로 사용되지만, 경우에 따라 "break 레이블"과 같이 사용하여 지정된 레이블로 이동할 수도 있다. break의 "레이블"은 보통 현재의 for 루프를 바로 위에 적게 되는데, 이러한 "break 레이블"은 현재의 루프를 빠져나와 지정된 레이블로 이동하고, break문의 직속 for 루프 전체의 다음 문장을 실행하게 한다. 아래 예제는 언뜻 보기에 무한루프를 돌 것 같지만, 실제로는 OK를 출력하고 프로그램을 정상 종료한다. 이는 "break L1" 문이 for 루프를 빠져나와 L1 레이블로 이동한 후, break가 있는 현재 for 루프를 건너뛰고 다음 문장인 println() 으로 이동하기 때문이다.


~~~go
package main
 
func main() {
    i := 0
 
L1:
    for {
     
        if i == 0 {
            break L1
        }
    }
 
    println("OK")
}
~~~