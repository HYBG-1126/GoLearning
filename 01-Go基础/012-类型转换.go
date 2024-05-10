package main

import "fmt"

func main() {

	//go语言中所有的转换都要显示转换

	var a int = 5

	fmt.Printf("a:%T", a)

	//奖a转换成字符串类型
	c := string(a)
	fmt.Printf("a:%T", c)

	// 整型无法转换成bool类型
}
