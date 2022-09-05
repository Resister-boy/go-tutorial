package main

import "fmt"

func main() {
	// var는 변수 선언을 의미
	// var [변수명] [변수 타입] = [변수 값]
	var a int = 10
	var msg string = "Hello Variable"
	// var로 선언된 변수는 수정 가능
	// 아래의 경우 위에서는 a = 10이었지만, 아래에서 a = 20으로 변환
	// 변수 수정은 같은 type에 한해 가능
	// const로 선언된 변수는 수정 불가능
	a = 20
	msg = "Good Morning"
	// fmt package에 있는 기능 Println() 실행
	// 변수 a, msg 출력
	fmt.Println(a, msg)
}