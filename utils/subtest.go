package utils

import "fmt"
import "github.com/bwangelme/testmod"

func Sub() {
	fmt.Println("this is the testmod/utils package")
}

func Tag() {
	fmt.Println("go mod use tag to downlod ,so new contene --> new tag ,or fail")
	fmt.Println("你好")
	fmt.Println(testmod.Hi("赵眼界", "cn"))
}
