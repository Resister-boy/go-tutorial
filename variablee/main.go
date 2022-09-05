// Properties of variable
// => name, value, memory, type

// memory address는 메모리 상 변수가 저장된 위치의 시작 지점을 의미
// type은 변수의 size를 의미 => int, float, string etc...
// 때문에 address와 type을 알 경우 변수가 저장된 공간을 추적할 수 있음. 

// Go에 존재하는 type
// int/uint => 0
// float => 0.0
// bool => false
// string => ""
// slice => nil
// array => nil
// struct => nil
// pointer => nil
// function => nil
// map => nil
// interface => nil
// channel => nil
// 별칭타입?
// nil은 정의되지 않은 memory address를 의미하는 Go의 키워드

package main

import "fmt"

func main() {
	var a int = 10
	b := 12 // 이 경우 Go Compiler가 type을 추론, 선언대입문이라고 함
	var c int // 이 경우 기본값 0으로 출력
	d := 3.14 // 이 경우 기본적으로 float64로 추론(Not float32)
	e := "Hello World" // type 추론을 통해 string type으로 선언

	fmt.Println(a, b, c, d, e)
	// Go에서 각 항목의 type은 반드시 일치해야 함 
	// int + float, int8 + int16도 안됨

	// Go에서 type 변환 => Go Compiler가 허용하는 범위 내에서 
	var f int = 1
	var g float64 = 1.1
	// f + g를 연산할 경우 f를 float64로 바꾸던지, g를 int로 바꿔야 함
	var h = int(g) // float을 int로 바꿀 경우 소수점 아래는 무조건 삭제(반올림X)
	fmt.Println(f + h) // 2
}