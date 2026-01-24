package main

import "fmt"

// Go不存在指针操作，只需要记住两个符号：&取地址符和*解引用符号（根据地址取值）

func main() {
	// &取地址符
	a := 18
	fmt.Println(&a)       // 得到a的地址
	p := &a               // 将a的地址储存到p中  指针
	fmt.Printf("%T\n", p) // *int类型指针 指向int类型地址的变量
	// *解引用符
	m := *p                     // p为指针存放着地址 通过地址取值
	fmt.Printf("%T %d\n", m, m) // int 18
	var b *int                  // 定义一个int类型的指针
	fmt.Println(b)              // <nil>
	/* 非法的  对指针a解引用后再赋值 本质上给其指向的地址存放值 但是指针b没有指向 是nil的
	*b = 100                  // 本质是给a指向的地址赋值
	fmt.Println(*b)           // 取地址中存放的值
	*/
	//使用new函数申请一个内存地址  new(存放的数据类型)
	var c *int
	c = new(int)
	fmt.Println(c, *c) // 0xc00000a110 0 非空 并且默认值为0
}
