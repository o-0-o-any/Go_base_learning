package main

import "fmt"

/*
类型断言语法格式
x.(T)
*/

// 函数参数为空接口类型 函数功能用来类型断言
func assert(x interface{}) {
	fmt.Printf("%T\n", x)
	ans, ok := x.(string) // 如果不写ok这个值来接收返回的bool值就会引发panic然后程序终止 // panic: interface conversion: interface {} is int, not string
	if ok {
		fmt.Printf("%v\n", ans)
	} else {
		fmt.Println("猜错了!")
	}
}

// type switch
func assertSwitch(x interface{}) {
	fmt.Printf("%T\n", x)
	switch x.(type) {
	case string:
		fmt.Println("string:", x)
	case int:
		fmt.Println("int:", x)
	default:
		fmt.Println("other:", x)
	}
}

func main() {
	assert(100)
	/*
		int
		猜错了!
	*/
	assert("hello world")
	/*
		string
		hello world
	*/
	assertSwitch(true)
	/*
		bool
		other: true
	*/
	assert("Hello World")
	/*
		string
		Hello World
	*/
}
