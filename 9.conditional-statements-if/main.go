package main

import (
	"fmt"
)

// var cnt int = 0

// func IncreaseAndReturn() int {
// 	fmt.Println("IncreaseAndReturn(): ", cnt)
// 	cnt++
// 	return cnt
// }

func HasRichFriend() bool {
	return true
}

func GetFriendsCount() int {
	return 3
}

func main() {

	price := 35000

	if price >= 50000 {
		if HasRichFriend() {
			fmt.Println("앗 신발끈이 풀렸네")
		} else {
			fmt.Println("나눠내자")
		}
	} else if price >= 30000 {
		if GetFriendsCount() > 3 {
			fmt.Println("어이쿠 신발끈이...")
		} else {
			fmt.Println("사람도 얼마 없는데 나눠내자")
		}
	} else {
		fmt.Println("오늘은 내가 쏜다")
	}

// Short-Circuit
	// if true && IncreaseAndReturn() < 5 {
		// fmt.Println("1 증가")
	// }
	// temp := 33
	// if temp > 28 {
	// 	fmt.Println("Turn on")
	// } else if temp <= 3 {
	// 	fmt.Println("Turn off")
	// } else if temp <= 18 {
	// 	fmt.Println("Get out")
	// } else {
	// 	fmt.Println("Come in")
	// }
}

// && AND
// || OR
// != NOT	

