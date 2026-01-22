/*
Go学习的一个程序 打印Hello World
*/

// 声明编译单元
package main

// 导入
import "fmt"

// 函数外部不能写go语句 只能写标识符的声明
// fmt.Println("Hello World")  // 不合法

// 程序的入口函数
func main() {
	// 调用fmt中的Println函数 只有首字母大小的属性与方法可以被调用 类似Public属性 首字母小写的类似protected属性
	fmt.Println("Hello World")
}
