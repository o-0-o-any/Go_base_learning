package main

import "fmt"

/*
函数类型作为参数和返回值
函数也是一种标识符 也是有它自己的类型
*/
// 定义一个函数
func fu1() {
	fmt.Println("Hello World!")
}
func fu2(x int, y int) (z int) {
	z = x + y
	return
}
func fu3(x string) {
	fmt.Println(x)
}
func fu4() bool {
	return true
}

// 函数类型变量作为函数参数
func fu5(x func(int, int) int) string { // 函数类型变量作为函数参数 main函数中b为func(int, int) int类型变量
	return "Hello World!"
}

// 函数类型变量作为函数返回值
func fu6() (e func(string, string) int) {
	e = func(x, y string) int { // 给返回值赋值 赋的值类似函数的定义
		fmt.Println(x, y)
		return 0
	}
	return
}

func main() {
	a := fu1              // 定义一个f1()函数类型的变量a
	fmt.Printf("%T\n", a) // func()

	b := fu2              // 定义一个f2()函数类型的变量b
	fmt.Printf("%T\n", b) // func(int, int) int

	c := fu3              // 定义一个f3()函数类型的变量c
	fmt.Printf("%T\n", c) // func(string)

	d := fu4              // 定义一个f4()函数类型的变量d
	fmt.Printf("%T\n", d) // func() bool

	fmt.Println(fu5(b)) // Hello World!

	e := fu6()            // 用e接收函数的返回值 返回值是一个函数
	fmt.Printf("%T\n", e) // func(string, string) int
	e("Hello", ".go!")    // Hello .go!  调用返回的函数e  跟定义一个函数类似

}
