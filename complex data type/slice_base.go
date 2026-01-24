package main

import "fmt"

/*
切片定义
var 变量名 []切片类型
与数组不同的是 []内不写长度了
*/
func main() {
	// 定义一个切片
	var a []int                 // 定义一个存放int类型元素的切片
	var b []string              // 定义一个存放string类型元素的切片
	fmt.Printf("%T %T\n", a, b) // []int []string
	fmt.Println(a, b)           //  [] []  与数组不同 切片各个位置上是没有默认值的 因为他的长度是没有限制的
	// 切片的初始化方式与数组类似
	a = []int{1, 2, 3, 4}
	b = []string{"Hello", "World", ".go", "!!"}
	var c []int
	var d = []bool{false, true}    // var 切片变量名 = []元素类型{初始化值}
	e := []int{5, 6, 10, 9, 16, 2} // 短声明 切片变量名 := []元素类型{初始化值}
	fmt.Println(d, e)

	// fmt.Println(a == c)  // 非法 切片直接不能直接比较
	// 切片只能与nil进行比较 相当于Python中的None C中的NULL C++中的NULL与nullptr
	fmt.Println(a == nil, c == nil) // false true
	// 长度和容量
	fmt.Println(len(a), len(b), len(c)) // 4 4 0
	fmt.Println(cap(a), cap(b), cap(c)) // 4 4 0
	// 由数组得到切片
	// 定义数组并初始化  [1:]从下标为1到最后 [:]全部元素 [:4]从第一个元素到4
	var f = [...]int{1, 2, 3, 4, 5}
	g := f[0:3]                // 左包含右不包含
	fmt.Printf("%T %T ", f, g) // [5]int []int
	fmt.Println(g)             // [1 2 3 4 5]
	h := f[3:]
	// 切片的长度就是切片含有的元素个数 切片的容量是底层数组从切片操作的第一个元素到最后元素的个数
	//如[0:3]这个切片操作是从下标为0的元素开始 那容量就为5 [3:]这个切片操作是从下标为3的元素开始 到最后一个元素的个数是2 那容量就是2
	fmt.Println(len(g), cap(g)) // 3 5
	fmt.Println(len(h), cap(h)) // 2 2
	// 切片再切片
	var a1 = [...]int{1, 2, 3, 4, 5, 6, 7, 8}
	i := a1[3:]
	j := i[3:]
	fmt.Println(len(i), cap(i)) // 5 5
	fmt.Println(len(j), cap(j)) // 2 2
	// 切片是引用数组
	fmt.Println("a1修改前i:", i) // a1修改前i: [4 5 6 7 8]
	a1[3] = 9999
	fmt.Println("a1修改后i:", i) // a1修改后i: [9999 5 6 7 8]
	// 底层数组修改后 切片也随之修改 因此切片是一个引用类型 是一个指向切片操作的第一个元素的指针
	// make函数创造切片 切片名 := make([]切片元素类型, 切片长度, 切片容量)
	k := make([]int, 5, 10)                  //make创造之后会有默认值
	fmt.Println(k, len(k), cap(k), k == nil) // [0 0 0 0 0] 5 10 false
	// 判断切片是否为空
	var l []int
	m := []int{}
	n := make([]int, 3, 5)
	fmt.Println(l, len(l), cap(l), l == nil) // [] 0 0 true
	fmt.Println(m, len(m), cap(m), m == nil) // [] 0 0 false
	fmt.Println(n, len(n), cap(n), n == nil) // [0 0 0] 3 5 false
	// 用nil只能判断切片有没有指向一个底层数组 不能判断其元素是否为空
	// 要判断一个切片是否是空的 要用len()方法判断其元素个数是否为0 不能使用nil来判断
	// 切片的赋值拷贝
	o := []int{1, 2, 3}
	p := o // 切片o p 都指向数组{1, 2, 3} 通过任一个切片修改数组中的元素 两个切片的元素都会发生变化
	// 反应出切片的本质是引用底层数组
	fmt.Println(o, p) // [1 2 3] [1 2 3]
	o[1] = 999
	fmt.Println(o, p) // [1 999 3] [1 999 3]
	// 切片的遍历
	// 使用下标遍历
	for z := 0; z < len(o); z++ {
		fmt.Print(o[z], " ")
	} // 1 999 3
	fmt.Println()
	// 使用 for range 遍历
	for _, z := range o {
		fmt.Print(z, " ")
	} // 1 999 3
	fmt.Println()
}
