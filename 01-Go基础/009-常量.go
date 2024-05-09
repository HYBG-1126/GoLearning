package main

import "fmt"

func main() {

	const a int = 2
	var b = 3

	fmt.Print(a)
	fmt.Println(b)

	//a = 3
	b = 4
	fmt.Println(a)
	fmt.Println(b)
}
