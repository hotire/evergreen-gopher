# Struct 

http://golang.site/go/article/16-Go-%EA%B5%AC%EC%A1%B0%EC%B2%B4-struct

Go에서 struct는 Custom Data Type을 표현하는데 사용되는데, Go의 struct는 필드들의 집합체이며 필드들의 컨테이너이다. Go에서 struct는 필드 데이타만을 가지며, (행위를 표현하는) 메서드를 갖지 않는다.

Go 언어는 객체지향 프로그래밍(Object Oriented Programming, OOP)을 고유의 방식으로 지원한다. 즉, Go에는 전통적인 OOP 언어가 가지는 클래스, 객체, 상속 개념이 없다. 전통적인 OOP의 클래스(class)는 Go 언어에서 Custom 타입을 정의하는 struct로 표현되는데, 전통적인 OOP의 클래스가 필드와 메서드를 함께 갖는 것과 달리 Go 언어의 struct는 필드만을 가지며, 메서드는 별도로 분리하여 정의한다 (Go Method 에서 설명).

### Struct 선언

struct를 정의하기 위해서는 Custom Type을 정의하는데 사용하는 type 문을 사용한다. 예를 들어 name과 age 필드를 갖는 person 이라는 struct를 정의하기 위해서는 아래와 같은 type문을 사용할 수 있다. 만약 이 person 구조체를 패키지 외부에서 사용할 수 있게 하려면 (Go 패키지에서 설명하였듯이) struct명을 Person으로 변경하면 된다.

### Struct 객체 생성

선언된 struct 타입으로부터 객체를 생성하는 방법은 몇 가지 방법들이 있다. 위의 예제처럼 person{} 를 사용하여 빈 person 객체를 먼저 할당하고, 나중에 그 필드값을 채워넣는 방법이 있다. struct 필드를 엑세스하기 위해서는 . (dot)을 사용한다.

### 생성자(constructor) 함수

때로 구조체(struct)의 필드가 사용 전에 초기화되어야 하는 경우가 있다. 예를 들어, struct 의 필드가 map 타입인 경우 map을 사전에 미리 초기화해 놓으면, 외부 struct 사용자가 매번 map을 초기화 해야 된다는 것을 기억할 필요가 없다. 이러한 목적을 위해 생성자 함수를 사용할 수 있다. 생성자 함수는 struct를 리턴하는 함수로서 그 함수 본문에서 필요한 필드를 초기화한다.

아래 예제에서 생성자 함수 newDict()는 dict라는 struct의 map 필드를 초기화한 후 그 struct 포인터를 리턴하고 있다. 이어 main() 함수에서 struct 개체를 만들 때 dict 를 직접 생성하지 않고 대신 newDict() 함수를 호출하여 이미 초기화된 data 맵 필드를 사용하고 있다.

