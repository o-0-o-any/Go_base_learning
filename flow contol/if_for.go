package main

import "fmt"

/*
条件判断的格式如下
if 表达式1 {    每一个左括号都有与if-else if-else在同一行

		分支1
	}else if 表达式2 {

		分支2
	}else {

		分支3
	}
*/
func main() {
	age := 20
	/*
		if age >= 18
		{

		}
	*/ //非法 左括号必须要与if-else if-else在同一行
	if age >= 18 {
		fmt.Println("你已成年 可以无限畅玩游戏!")
	} else if age < 9 {
		fmt.Println("你不能玩这个游戏!")
	} else {
		fmt.Println("你一天只能玩两个小时!")
	}
	// 作用域
	if age2 := 19; age2 >= 18 { // age2只在if判断语句中生效
		fmt.Println("成年了!")
	} else {
		fmt.Println("没成年!")
	}
	// fmt.Println(age2) // 不合法 age2声明在if里的作用域 在这里是未解析的变量名
	// 循环语句
	// 格式1 与C/C++中的for基本相同
	for i := 0; i <= 5; i++ {
		fmt.Print(i, " ") // 0 1 2 3 4 5
	}
	fmt.Println()
	// 初始语句可以省略 终止语句可以放在循环体中 但是分号不能省略
	j := 0
	for j <= 10 {
		j++
		fmt.Print(j, " ") // 0 1 2 3 4 5 6 7 8 9 10
	}
	fmt.Println()
	/*  死循环
	for {
		fmt.Println("1")
	}
	*/
	// for range循环
	s := "Hello World"
	for index, char := range s {
		fmt.Printf("%c(%d)", char, index) // H(0)e(1)l(2)l(3)o(4) (5)W(6)o(7)r(8)l(9)d(10)
	}
	fmt.Println()
	// 打印99乘法表
	for i := 1; i < 10; i++ {
		for j := 1; j <= i; j++ {
			fmt.Print(j, "x", i, "=", j*i, " ")
		}
		fmt.Println()
	}
}
