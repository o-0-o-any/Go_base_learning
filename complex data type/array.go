package main

import (
	"fmt"
)

/*
数组定义
var 数组名 [数组长度]数组元素类型
如果数组长度不同 即使数组元素类型相同 这两个数组也不可以进行赋值等操作
*/

func main() {
	// 定义两个数组
	var a [5]int
	var b [10]int
	// 打印数组类型 两个数组的数据类型不同
	fmt.Printf("%T %T", a, b) // [5]int [10]int
	// a = b  // 不合法 两个数组的数据类型不同
	// 数组的初始化
	var c [5]bool
	fmt.Println(a, b, c) // 如果不对数组赋值 那么都是默认值 int为0 bool为false string为""
	// 初始化数组
	// 初始化方式1
	var d [5]int
	d = [5]int{1, 2, 3, 4, 5}
	fmt.Println(d) // [1 2 3 4 5]
	// 初始化方式2
	e := [10]int{1, 2, 3, 4, 5}
	fmt.Println(e) // [1 2 3 4 5 0 0 0 0 0] 如果赋的值个数不足数组长度 那么没有赋值的位置默认值就是0
	// 初始化方式3 根据初始化值个数推断数组长度 中括号中写...
	f := [...]int{1, 2, 3, 4, 5, 6, 10}
	fmt.Println(len(f), f) // 7 [1 2 3 4 5 6 10]
	// 初始化方式4 下标从0开始 数组长度-1结束 类似Python中的字典 指定下标位置的值 未指定的为0
	g := [8]int{1: 8, 7: 6, 5: 3}
	fmt.Println(g) // [0 8 0 0 0 3 0 6]
	// 数组的遍历
	city := [...]string{"河南", "河北", "山东", "山西"}
	// 使用下标遍历 下标从0开始 从数组长度-1结束
	for i := 0; i < len(city); i++ {
		fmt.Print(city[i], " ")
	}
	fmt.Println()
	// for range 遍历
	for _, s := range city { // 使用匿名函数省略掉每个元素的下标
		fmt.Print(s, " ")
	}
	fmt.Println()
	// 多维数组
	var h [3][2]int
	h = [3][2]int{{1, 2}, {3, 4}, {5, 6}}
	fmt.Printf("%T ", h) // [3][2]int
	fmt.Println(h)       // [[1 2] [3 4] [5 6]]
	i := [2][3]string{
		{"河南", "商丘", "虞城县"},
		{"河南", "南阳", "宛城区"},
	}
	fmt.Printf("%T ", i) // [2][3]string
	fmt.Println(i)       // [[河南 商丘 虞城县] [河南 南阳 宛城区]]
	// 多维数组遍历
	// for range遍历
	for _, s2 := range i {
		fmt.Print(s2, " ")
	} // [河南 商丘 虞城县] [河南 南阳 宛城区] 打印的是多维数组内的数组元素而不是打印内置数组中的值
	// 如果要打印每个内置数组的元素 那就需要嵌套一个for循环继续打印一维数组
	fmt.Println()
	// 下标打印
	for j := 0; j < len(i); j++ {
		fmt.Print(i[j], " ")
	} // [河南 商丘 虞城县] [河南 南阳 宛城区] 同上
	fmt.Println()
	// 遍历二维数组中数组的元素  使用嵌套循环 先得到二维数组中的内置数组 然后再变量内置一维数组的元素
	for j := 0; j < len(i); j++ {
		for k := 0; k < len(i[j]); k++ {
			fmt.Print(i[j][k], " ")
		}
	} // 河南 商丘 虞城县 河南 南阳 宛城区
	fmt.Println()
	for _, s3 := range i {
		for _, s4 := range s3 {
			fmt.Print(s4, " ")
		}
	} // 河南 商丘 虞城县 河南 南阳 宛城区
	fmt.Println()
	// 数组是值类型  相同数组类型的可以相互赋值 可以通过下标修改数组元素
	j := [3]int{6, 3, 9}
	k := j
	k[1] = 12
	fmt.Println(j, k) // [6 3 9] [6 12 9]
}
