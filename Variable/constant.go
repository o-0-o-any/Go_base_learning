/*
学习Go的常量
*/

package main

import "fmt"

// 常量声明
// 逐个声明
const PI = 3.1415926

// 批量声明 分别赋值
const (
	n1 = 100
	n2 = 200
	n3 = 300
)

// 批量声明 省略赋值  未赋值的常量的值与其上一个常量的值相同
const (
	b1 = 20
	b2
	b3 = 30
	b4
)

// Go语言的计数器-iota
const (
	c1 = iota // 0
	c2        // 1
	c3        // 2
	c4        // 3
)
const (
	a1 = iota // 0
	a2 = iota // 1
	a3 = iota // 2
	a4 = iota // 3
)

// 常见iota示例
// 使用匿名变量_跳过某些枚举值
const (
	d1 = iota
	d2
	_
	d3
)

// iota声明中插入其他值 与匿名变量跳过某些枚举值类似 这里理解成用自定义赋值覆盖了对应枚举值
const (
	e1 = iota // 0
	e2 = 100  // 100
	e3 = iota // 2
	e4        // 3
)

// 多个常量声明在同一行  iota计算是依据行数来看的 如果多个常量声明在同一行 那么他们的iota赋值是一样的
const (
	f1, f2 = iota, iota // 0 0
	f3, f4 = iota, iota // 1 1
	f5, f6 = iota, iota // 22
)

// 定义数量级 <<左移操作符 例如KB的语句 iota为1 1左移十位得到二进制数字10000000000转换成十进制就是2的10次方即1024
const (
	_  = iota             // 0 使用匿名函数跳过值0
	KB = 1 << (10 * iota) // 1024
	MB = 1 << (10 * iota) // 1024 * KB
	GB = 1 << (10 * iota) // 1024 * MB
	TB = 1 << (10 * iota) // 1024 * GB
	PB = 1 << (10 * iota) // 1024 * TB
)

func main() {
	//PI = 3.14  //不合法
	// 使用声明的常量
	fmt.Println(PI)
	fmt.Println(n1, n2, n3)
	fmt.Println(b1, b2, b3, b4)
	fmt.Println(c1, c2, c3, c4)
	fmt.Println(a1, a2, a3, a4)
	fmt.Println(d1, d2, d3)
	fmt.Println(e1, e2, e3, e4)
	fmt.Println(f1, f2, f3, f4, f5, f6)
	fmt.Println(KB, MB, GB, TB, PB)
}
