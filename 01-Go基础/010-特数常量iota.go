package main

import "fmt"

func main() {

	//只有const出现时

	const (
		a = iota
		b
		c
	)

	fmt.Println(b)

	//const中不定义只声明，会继承上一个值,iota相当于索引
	const (
		aa = iota
		bb
		cc
		dd = "haha" //iota 3
		ee          //iota 4
		ff = 100    //iota 5
		gg          //iota 6
		hh = iota   //iota 7
		ii
	)

	fmt.Println(aa, bb, cc, dd, ee, ff, gg, hh, ii)

	//重新开始const就是新的iota
	const (
		jj = iota
		kk
	)

	fmt.Println(jj, kk)

}
