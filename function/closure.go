package main

import "fmt"

// 闭包

// 定义函数 函数类型变量作为函数参数
func func1(f func()) {
	fmt.Println("进入func1函数体 调用参数函数")
	f()
}
func func2(x, y int) {
	fmt.Println("进入func2函数体")
	fmt.Println(x + y)
}

// 定义func3函数对func2函数进行包装
// 将func(int int)类型的变量传入func3函数中 经过func3函数处理 返回一个func()类型的变量//
// 经过这样一个中间处理的过程 func1就可以顺利执行了
func func3(f func(int, int), m, n int) func() {
	temp := func() { // 定义匿名函数变量
		fmt.Println("func3中定义的匿名函数 用来调用传入的参数")
		f(m, n)
	}
	return temp
}

func main() {
	// func1为一个函数接口 传入的函数参数应该是func()类型
	// 但是func2函数是func(int int)类型 怎么将func2函数作为参数传入到func1函数中呢
	// func1(func2) // 这种直接传入肯定不行 类型不匹配  cannot use func2 (value of type func(x int, y int)) as func() value in argument to func1
	// 应该定义一个函数对func2函数进行包装
	// 通过包装func2的函数func3 传入func1函数中
	func1(func3(func2, 91, 77))
	/*
		进入func1函数体 调用参数函数   // 进入主函数
		func3中定义的匿名函数 用来调用传入的参数   // 进入中间的、对func2进行包装的函数
		进入func2函数体  // 进入func2函数
		168   // 成功完成func2函数变量传入func1函数的操作
	*/
}
