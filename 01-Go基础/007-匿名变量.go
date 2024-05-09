package main

import "fmt"

func main() {

	// 匿名变量 _ ,这个变量，会废弃掉

	a, _ := test(1, 2)

	fmt.Println(a)

}

func test(a int, b int) (int, int) {
	return a, b
}
