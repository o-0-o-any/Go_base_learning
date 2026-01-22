/*
学习Go的浮点型与布尔型
*/
package main

import "fmt"

func main() {
	// 声明浮点数
	f1 := 1.23456
	// 打印f1类型
	fmt.Printf("%T\n", f1) //float64 默认Go语言中的小数都是float64类型
	// 指定小数为float32类型 强制转换
	f2 := float32(1.23456)
	fmt.Printf("%T\n", f2)
	// 赋值
	// f1 = f2 // 不合法 float32与float64不能相互赋值
	// 打印浮点数的值
	fmt.Printf("%f %f\n", f1, f2)
	// 声明复数
	var c1 complex64
	c1 = 1 + 2i
	var c2 complex128
	c2 = 3 + 4i
	// 打印类型
	fmt.Printf("%T %T\n", c1, c2) // complex64 complex128
	// 赋值
	// c1 = c2 // 不合法 complex64与complex128不能相互赋值
	// 打印复数
	fmt.Println(c1, c2)
	// 声明布尔类型
	var b1 bool
	b1 = true
	var b2 bool // 默认值为false
	// 打印数据类型
	fmt.Printf("%T %T\n", b1, b2)
	// 打印数据的值
	fmt.Println(b1, b2)
}
