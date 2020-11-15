# Struct 

http://golang.site/go/article/16-Go-%EA%B5%AC%EC%A1%B0%EC%B2%B4-struct

Go에서 struct는 Custom Data Type을 표현하는데 사용되는데, Go의 struct는 필드들의 집합체이며 필드들의 컨테이너이다. Go에서 struct는 필드 데이타만을 가지며, (행위를 표현하는) 메서드를 갖지 않는다.

Go 언어는 객체지향 프로그래밍(Object Oriented Programming, OOP)을 고유의 방식으로 지원한다. 즉, Go에는 전통적인 OOP 언어가 가지는 클래스, 객체, 상속 개념이 없다. 전통적인 OOP의 클래스(class)는 Go 언어에서 Custom 타입을 정의하는 struct로 표현되는데, 전통적인 OOP의 클래스가 필드와 메서드를 함께 갖는 것과 달리 Go 언어의 struct는 필드만을 가지며, 메서드는 별도로 분리하여 정의한다 (Go Method 에서 설명).