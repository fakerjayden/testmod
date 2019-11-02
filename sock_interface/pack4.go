package sock_interface

import "fmt"

// 会出现红色标记：使用enable go module
import "github.com/fakerjayden/testmod/third_party"

func Test_4() {
	fmt.Println("use go get -u to force update local")
	third_party.Test_5()
}
