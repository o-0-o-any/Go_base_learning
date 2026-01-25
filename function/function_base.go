package main

import "fmt"

/*
函数定义格式
func 函数名 (函数参数) (返回值类型) {
    函数体
}
返回值由返回值变量和其变量类型组成，也可以只写返回值的类型，多个返回必须用括号()包裹，每个返回值并用逗号,分隔
*/

// 函数定义
// 返回值写变量+变量类型 相当于提起定义了返回的变量
func sum1(a int, b int) (c int) {
	c = a + b
	return // 定义时声明了返回的变量 return后就不要加返回内容了 默认返回声明的变量
}

// 返回值仅写数据类型
func sum2(a int, b int) int {
	c := a + b
	return c
	// return a + b // 也可以返回一个语句
}

// 没有返回值
func sum3(a int, b int) {
	c := a + b
	println("a + b = ", c)
}

// 没有参数
func f4() {
	for i := 1; i < 10; i++ {
		for j := 1; j < i; j++ {
			fmt.Printf("%d%d=%d", j, i, j*i)
		}
		fmt.Println()
	}
}

// 没有参数但有返回值
func f5() string {
	return "Hello World!"
}

// 多个返回值
func f6() (string, string) {
	return "Hello", "World"
}

// 当参数中连续两个参数类型一样时 可以将前面那个参数的类型省略
func f7(x, y int, m, n string, o, p, q bool) {
	fmt.Printf("%T %T %T %T %T, %T, %T\n", x, y, m, n, o, p, q) //
}

// 可变长参数 相当于Python中的args参数 可变长参数必须放在后面
func f8(x string, y ...int) {
	fmt.Println(x)
	fmt.Println(y) // y的类型是[]int slice类型
}

// Go语言没有默认参数这个概念

// defer语句封装到函数
func defer_f9() {
	fmt.Println("进入函数")
	defer fmt.Println("函数即将返回时执行 第一个定义 最后一个执行")
	defer fmt.Println("函数即将返回时执行 第二个定义 第二个执行")
	defer fmt.Println("函数即将返回时执行 最后一个定义 第一个执行")
	fmt.Println("函数即将返回")
}

// defer的执行时机 在给返回值赋值操作之后 执行RET指令之前
func defer_f10() int {
	x := 10
	defer func() {
		x++ // 修改的是局部变量x不是返回值
	}()
	return x // 将当前x的值赋值给返回值(赋值操作) 即进行操作 返回值 = x = y = 10
}
func defer_f11() (x int) {
	x = 10
	defer func() {
		x++ // 这里修改的是返回值
	}()
	return // 返回值进行了自加 变成了11  返回值 = x++ = 11
}
func defer_12() (y int) {
	x := 10
	defer func() {
		x++ // 修改的是局部变量x不是返回值
	}()
	return x // 将当前x的值赋值给返回值(赋值操作)  即进行操作 返回值 = y = x = 10
}

func main() {
	// 函数调用 与Python C/C++一样
	res := sum1(10, 12)
	fmt.Println(res)

	fmt.Println(sum2(20, 99))

	fmt.Println(f6())

	f7(1, 2, "Hello", "World", false, true, true)

	f8("Hello World", 2026, 1, 25, 16, 12)

	// defer语句
	fmt.Println("START")
	defer fmt.Println("one")
	defer fmt.Println("two")
	defer fmt.Println("three")
	fmt.Println("0.5END")
	fmt.Println("END")
	/*
		START
		0.5END
		END
		three
		two
		one
	*/
	// 调用函数
	defer_f9()
	/*
			函数即将返回时执行 最后一个定义 第一个执行
		函数即将返回时执行 第二个定义 第二个执行
		函数即将返回时执行 第一个定义 最后一个执行
	*/
	fmt.Println(defer_f10()) // 10
	fmt.Println(defer_f11()) // 11
	fmt.Println(defer_12())  // 10
}
