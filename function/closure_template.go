package main

import "fmt"

/*
闭包的应用 动态生成自定义函数
类似C++中的类模板
*/

// 根据传入的op数学运算符 动态生成专门用于处理对应运算的函数
func operator(op string) func(x, y int) int {
	if op == "+" {
		return func(x, y int) int {
			return x + y
		}
	} else if op == "-" {
		return func(x, y int) int {
			return x - y
		}
	} else if op == "*" {
		return func(x, y int) int {
			return x * y
		}
	} else if op == "/" {
		return func(x, y int) int {
			return x / y
		}
	} else {
		return func(x, y int) int {
			fmt.Print("本函数仅支持加减乘除四种运算!")
			return 0
		}
	}
}

func main() {
	sum := operator("+")       // 用来运算+的函数
	subtract := operator("-")  // 用来运算-的函数
	multiply := operator("*")  // 用来运算*的函数
	divide := operator("/")    // 用来运算/的函数
	other := operator("Hello") // else情况

	fmt.Println(sum(10, 20))      // 30
	fmt.Println(subtract(52, 17)) // 35
	fmt.Println(multiply(4, 8))   // 32
	fmt.Println(divide(18, 2))    // 9
	fmt.Println(other(12, 34))    // 本函数仅支持加减乘除四种运算! 0
}
