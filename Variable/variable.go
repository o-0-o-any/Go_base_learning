/*
学习Go的变量
*/
package main

import "fmt"

// 变量声明 可以在函数内外都可以 除作用域不同外 在函数内部声明的变量已经声明必须使用不然会报错 而在外部声明的在其他函数内不使用不会报错 这是Go 1.18版本之后优化的地方 不完全遵循声明的变量必须使用
// 逐个声明
var a int
var b string
var c bool

// 批量声明
var (
	d int    // int类型初始值0
	e string // string类型初始值""
	f bool   // bool类型类型初始值false
)

// g := 5  // 函数外部不能使用短变量声明

func main() {
	// var c int  //declared and not used: c 在函数内部只声明不使用报错
	// c = 20  //依然报错 赋值不属于是对变量操作
	// 对变量赋值
	d = 2
	e = "Hello World"
	f = true
	// 声明并赋值  var 变量名 变量类型 = 赋值内容
	var s1 int = 1
	var (
		s2 string = "史威化"
		s3 bool   = true
	)
	// 短变量声明 只能在函数内部使用
	c1 := 20
	c2 := "I am a student"
	c3 := true

	// 使用变量 比如打印变量的值
	fmt.Print(s1, s2, s3, "\n")                        // 与Println比较 这个方法少一个换行符
	fmt.Println(a, b, c, d, e, f)                      // 前三个未赋值打印初始化的值 这种打印方式类似Python中的print
	fmt.Printf("day: %d str: %s, str : %t\n", d, e, f) // 类似C语言的printf %d %s %t分别是int string bool对应的占位符
	fmt.Println(c1, c2, c3)
}
