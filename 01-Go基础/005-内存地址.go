package main

import "fmt"

func main() {

	var name string = "xiaohong"

	// &是取地址符
	fmt.Printf("name: %s,内存地址: %p", name, &name)

	fmt.Println()

	name = "xiaogang"

	fmt.Printf("name:%s,内存地址：%p", name, &name)
}
