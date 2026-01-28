package main

import "fmt"

// 匿名字段  没有名字的字段

// 定义一个结构体
type test1 struct {
	str string
	int // 匿名字段 没有字段名
	// int  // 非法 不能有两个类型相同的匿名字段
}

// 匿名字段主要用于字段比较少并且没有重复类型的场景 不常用

func main() {
	t1 := test1{
		"匿名字段",
		1,
	}
	// 匿名字段默认将类型作为其名字 通过类型名访问匿名字段
	fmt.Println(t1.str, t1.int) // 匿名字段 1
}
