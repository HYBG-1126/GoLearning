package main

import "fmt"

//基本数据类型
//布尔型
//数值型
//整数:int,int8,int16,int32,int64
// byte 类似 uint8
// rune 类似 int32
// uint 类似32或者64
// int 与unit一样大小
// uintptr 无符号整型，放置一个指针
//浮点:float21,float64,complex64,complex128
//字符串 string
//派生数据类型

func main() {

	// bool 类型
	var isflag bool

	fmt.Println(isflag)

	fmt.Printf("isflag:%T,isflap:%t\n", isflag, isflag)

	// 数值类型
	var num int64
	fmt.Printf("num:%T,num:%d\n", num, num)

	//浮点型，默认6位置小数
	var money float64
	fmt.Printf("money:%T,money:%.2f", money, money)

	// 字符串类型
	//双引号
	var name string = "xiaoming"
	fmt.Println(name)
	var sex string = "nan" + "nan"
	fmt.Println(sex)

}
