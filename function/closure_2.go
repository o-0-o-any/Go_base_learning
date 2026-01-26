package main

import "fmt"

// 定义一个函数
func F1() func(int) int { // 函数作为返回值
	var x = 1
	return func(y int) int { // 返回一个匿名变量
		x += y
		return x
	}
}

func main() {
	// 接收F1函数的返回值 函数变量
	o1 := F1()
	fmt.Println(o1(1)) // 2
	fmt.Println(o1(2)) // 4
	fmt.Println(o1(3)) // 7
	// 再次定义一个函数变量 接收F1函数的返回值
	o2 := F1()
	fmt.Println(o2(1)) // 2
	fmt.Println(o2(2)) // 4
	fmt.Println(o2(3)) // 7
	/*
		闭包是一个函数，这个函数包好了它外部作用域的一个变量
		如上 o1接受的是一个函数 调用o1或o2时 不是只看o1或o2函数内部的东西 还需要看其外层的变量等 如F1函数中的x变量
		底层原理就是:函数可以作为返回值、函数内部查找变量的顺序 现在自己内部找 找不到再往外部找
		就像F1函数内部声明匿名函数并作为返回值让o1/o2变量接收 F1中(不包含匿名函数中的元素)的变量对于o1/o2来说就像是全局元素 o1/o2内部是局部元素
	*/
}
