package main

import "fmt"

// 定义三个函数
func funcA() {
	fmt.Println("函数funcA")
}
func funcB() {
	fmt.Println("函数funcB")
	defer func() {
		error := recover()
		fmt.Println("defer函数:释放必要资源如关闭数据库连接等... 错误类型:", error)
	}()
	panic("函数B内部出现程序崩溃的严重问题!")
	// 执行到panic时找defer语句 没有就程序崩溃退出
	// 若没有recover 执行完defer语句依然退出程序
	// 若有recover语句 通过recover将程序恢复回来继续往后执行
}
func funcC() {
	fmt.Println("函数funcC")
}

func main() {
	funcA()
	funcB()
	funcC()
	/*
		没有recover时的执行结果  捕获到错误后终止程序运行
		函数funcA
		函数funcB
		defer函数:释放必要资源如关闭数据库连接等...       // 触发panic后开始找执行过函数中的defer语句
		panic: 函数B内部出现程序崩溃的严重问题!

		goroutine 1 [running]:
		main.funcB()
		        D:/Goland/project/Go_base_learn/function/built_in_function.go:11 +0x59
		main.main()
		        D:/Goland/project/Go_base_learn/function/built_in_function.go:19 +0x50
	*/
	/*
		有recover的执行结果  捕获到错误后继续执行函数funcC
		函数funcA
		函数funcB
		defer函数:释放必要资源如关闭数据库连接等... 错误类型: 函数B内部出现程序崩溃的严 重问题!
		函数funcC
	*/
}
