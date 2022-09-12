package main

import "fmt"

// 1byte === 트랜지스터 8개, 0과 1의 상태를 통해 2진수로 표현 가능
// uint8 1byte 부호없는 정수 0 ~ 255(트랜지스터 8개로 담을 수 있는 가장 큰 숫자)
// uint16 2byte 부호없는 정수 0 ~ 65535
// uint32 4byte 부호없는 정수 0 ~ 4294967295
// uint64 8byte 부호없는 정수 0 ~ 18446744073709551615
// int8 1byte 부호있는 정수 -128 ~ 127
// int16 2byte 부호있는 정수 -32768 ~ 32767
// int32 4byte 부호있는 정수 -2147483648 ~ 2147483647
// int64 8byte 부호있는 정수 -9223372036854775808 ~ 9223372036854775807
// float32 4byte 실수 IEEE-754 32bit 실수
// float64 8byte 실수 IEEE-754 64bit 실수
// complex64 8byte 복소수 float32의 범위와 같음
// complex128 16byte 복소수 float64 범위와 같음
// byte uint8의 별칭 1byte를 나타낼 때 사용 0 ~ 255
// rune int32의 별칭 UTF-8로 문자 하나를 나타낼 때 사용 -2147483648 ~ 2147483647
// int 32bit 컴퓨터에서는 int32 / 64bit 컴퓨터에서는 int64
// uint 32bit 컴퓨터에서는 uint32/ 64bit 컴퓨터에서는 uint64

var m int = 10 // Package Global Variable

func main() {
	var a int16 = 1234
	var b int8 = int8(a)

	// ------------------------------------//

	var s int = 20 // Function Variable

	{
		var g int = 40
		fmt.Println(m, s, g)
	}

	// m + s + g = ? 를 연산할 경우 Error
	// 변수가 Block 안에서 선언된 경우 해당 Block 안에서만 호출할 수 있기 때문
	// 위의 경우 g 때문에 Error 발생

	fmt.Println(a, b) // a = 1234, b = -46 반환
	// Size가 큰 type의 변수를 Size가 작은 type으로 변환할 때에는 변수의 일부가 삭제
	// 위의 경우 int16은 2byte, int8은 1byte일 때, int16의 상위 1byte가 삭제된 것.

}
