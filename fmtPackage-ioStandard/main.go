package main

import (
	"fmt"
	"bufio"
	"os"
)
// Print() => 입력 값 출력
// Println() => 입력 값 출력 및 개행
// Printf() => Fotmat에 맞게 입력 값 출력

// Sacn() => 값 입력
// Scanln() => 한 줄 값 입력
// Scanf() => Format에 맞게 값 입력

// func solution(a, b) {}
// Go에서 사용되는 argument는 항상 값으로 사용되며, 주소 값은 입력되지 않음
// argument의 메모리 주소는 &a와 같은 방식으로 argument 앞에 & 입력

func main() {	
	var a int = 10
	var b int = 20
	var c float64 = 30.0

	fmt.Print("a:", a, " b:", b, " c:", c)
	fmt.Println("a:", a, " b:", b, " c:", c) // 개행: 다음 줄로 넘어감, 반환 값 사이에 공백 추가
	fmt.Printf("a: %d b: %d c: %f\n", a, b, c) // %d: 서식 문자(Formatter)
	// %v: type에 맞는 기본 형태
	// %t: type 데이터 type
	// %d: Decimal 정수
	// %f: Float 실수
	// %c: Character 문자
	// %s: String 문자열
	// \n: 개행 문자

	stdin := bufio.NewReader(os.Stdin) // stdin: Standard Input
	
	var d int
	var e int

	number, err := fmt.Scanln(&d, &e) 
	if err != nil {
		fmt.Println(err)
		stdin.ReadString('\n')
	} else {
		fmt.Println(number, d, e) // number = input.length, d & e = input value
	}
}