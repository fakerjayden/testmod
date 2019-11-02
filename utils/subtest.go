package utils

import "fmt"

// 竟然在git push时候会下载并更新go sum
import "github.com/bwangelme/testmod"

func Sub() {
	fmt.Println("this is the testmod/utils package")
}

func Tag() {
	fmt.Println("go mod use tag to downlod ,so new contene --> new tag ,or fail")
	fmt.Println("你好")
	fmt.Println(testmod.Hi("赵眼界"))
}
