package main

import "fmt"

// 定义全局变量
var num1 int = 10
var num2 int = 20

func f1() {
	// 在函数内部可以访问全局变量
	fmt.Println("全局变量num1:", num1) // 全局变量num: 10
	num2 := 30
	// 先在函数内部找局部变量 没有再去函数外找 一直找到全局变量num2: 30
	fmt.Println("先在函数内部找局部变量 没有再去函数外找 一直找到全局变量num2:", num2)
	// fmt.Println(y)  // 如果都没有就会报错
}

func f2() {
	// 在f2内部声明一个局部变量
	var num3 int = 20
	fmt.Println(num3)
}

func main() {
	f1()
	// fmt.Println(num3)  // 非法 f2()函数内定义的变量无法再f2()函数外如main函数中使用
	// if判断语句中的作用域
	if num2 > 0 {
		num4 := 100 // 在if判断语句中定义了一个变量
		fmt.Println(num4)
	}
	// fmt.Println(num4 == 0)  // 非法 if内定义的变量只能在if内部使用
	// for循环语句中的作用域
	for i := 0; i < 10; i++ { // 在for循环语句中定义了一个变量
		fmt.Print(i, " ")
	}
	fmt.Println()
	// fmt.Println(i)   // 在for循环语句中定义的变量只能在for循环语句内使用
	// 其他结构以此类推 与C/C++ Python一样

}
