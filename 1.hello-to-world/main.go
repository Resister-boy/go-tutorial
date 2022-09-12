// Go는 무조건 package로 시작
// Go의 모든 코드는 package에 포함
// Go의 main.package는 프로그램의 시작점을 포함 & 의미하며, main.package는 유일
// Go는 main.package 하나와 여러 개의 다른 package들로 구성된다 할 수 있음

package main


// import는 다른 package를 가져옴을 의미
// 아래의 경우 fmt package를 Define하는 것이라 할 수 있음
// package는 자기 자신을 가져올 수 없으며, main.package는 가져올 수 없음

import (
	"fmt"
)

// main.package와 마찬가지로 func main()는 유일
// func main()은 main.package에서만 존재하며, func main()이 선언된 package가 main.package라고 할 수 있음.
// Go는 func main()에서 시작해서 funct main()에서 종료
func main() {

// import한 fmt package에 포함된 기능 Println()을 실행
	fmt.Println("Hello World");
}