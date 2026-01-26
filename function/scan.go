package main

import "fmt"

func main() {
	// 获取输入
	var (
		a int
		b string
		c bool
		d float32
	)
	/*
		fmt.Scan(&a, &b, &c, &d)          // 1 HelloWorld true 1.23  如果输入一个字符串中有空格 那么会被判定成两个变量的值
		fmt.Println("输入的内容:", a, b, c, d) // 输入的内容: 1 HelloWorld true 1.23
	*/
	fmt.Scanf("%d %s %t %f", &a, &b, &c, &d) // 1 HelloWorld true 1.23
	fmt.Println("输入的内容:", a, b, c, d)        // 输入的内容: 1 HelloWorld true 1.23
}
