# Map

http://golang.site/go/article/14-Go-%EC%BB%AC%EB%A0%89%EC%85%98---Map

Map은 키(Key)에 대응하는 값(Value)을 신속히 찾는 해시테이블(Hash table)을 구현한 자료구조이다. Go 언어는 Map 타입을 내장하고 있는데, "map[Key타입]Value타입" 과 같이 선언할 수 있다. 예를 들어 정수를 키로하고 문자열을 값으로 하는 맵 변수 idMap을 선언하기 위해서는 다음과 같이 할 수 있다.

다른 언어에서 사용하는 Map...

~~~go
var idMap map[int]string
~~~