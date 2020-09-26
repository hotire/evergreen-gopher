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