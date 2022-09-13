package main

// constant = 변하지 않는 수 
// var..가 아닌 const로 선언
// constant는 타입을 선언하지 않아도 ok
// constant는 동적 메모리를 사용하지 않음 
// type을 선언하지 않는 경우 constant가 사용될 때 결정
// const [변수명] [타입] = [변수값]

import (
	"fmt"
)

const Pig int = 0
const Cow int = 1
const Chicken int = 2

func PrintAnimal(animal int) {
	if (animal == Pig) {
		fmt.Println("꿀꿀")
	} else if animal == Cow {
		fmt.Println("음메")
	} else if animal == Chicken {
		fmt.Println("꼬끼오")
	} else {
		fmt.Println("...")
	}
}

func main() {
	const pi float64 = 3.1415926535897 // const 상수
	fmt.Println(pi)
	PrintAnimal(Pig) // 꿀꿀
	PrintAnimal(0) // 꿀꿀

	// iota로 열거 값 사용 가능 변수 값이 하나씩 증가
	const (
		Red int = iota
		Blue int = iota
		Green int = iota
	)
	fmt.Println(Red, Blue, Green) // 0, 1, 2 => 비트플래그에 활용
}