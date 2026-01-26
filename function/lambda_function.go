package main

import "fmt"

/*
匿名函数
var 变量名 = func(参数 参数类型, ......){
	函数体
}
如果只是调用一次函数 简写函数立即执行函数 多嵌入函数内部
func(参数 参数类型, ......) {

} (......)  // 这个括号中传参 如果无惨就不用写 但是括号一定要有
*/

func main() {
	// 匿名函数定义
	fuc1 := func(x, y string) {
		fmt.Println(x + y)
	}
	fuc1("Hello", "World") // HelloWorld

	// 简写函数成立立即执行
	func(x, y int) {
		fmt.Println(x + y)
	}(10, 11) // 21
	func() {
		fmt.Println("Hello, World")
	}() // 即使无惨也要写这个括号      Hello, World
}
