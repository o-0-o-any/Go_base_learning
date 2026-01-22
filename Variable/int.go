/*
学习Go的整型
*/
package main

import "fmt"

func main() {
	// 十进制
	var a int = 100
	fmt.Printf("%d\n", a) // 10
	fmt.Printf("%b\n", a) // %b为二进制的占位符 十进制转二进制  1100100
	fmt.Printf("%o\n", a) // 十进制转八进制
	fmt.Printf("%x\n", a) // 十进制转十六进制
	// 八进制 八进制以0开头
	var b int = 077
	fmt.Printf("%o\n", b) // %o为八进制的占位符
	// 十六进制 十六进制以0x开头
	var c int = 0xff
	fmt.Printf("%x\n", c) // ff
	fmt.Printf("%X\n", c) // FF
	// 查看变量类型
	fmt.Printf("%T %T %T\n", a, b, c) // int int int

	// 声明int8类型的变量 与Python中的类型转换类似
	d := int8(9)
	fmt.Printf("%T\n", d) // int8
}
