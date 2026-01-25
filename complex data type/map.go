package main

import "fmt"

/*
map的定义语法
map[键的类型KetType]值的类型ValueType
*/

func main() {
	// 声明一个map
	var a map[string]int
	/*  不合法 定义的a为nil 需要用make为其开辟内存空间
	a["大一"] = 19
	a["大二"] = 20
	a["明"] = 21
	fmt.Println(a)
	goroutine 1 [running]:
	main.main()
	        D:/Goland/project/Go_base_learn/complex data type/map.go:13 +0x2c
	*/
	a = make(map[string]int, 2) // 要估算好map容量 避免程序运行期间再动态扩容
	a["高等数学"] = 88
	a["Python程序设计"] = 85
	fmt.Println(a) // map[Python程序设计:85 高等数学:88]
	// 使用键找到值
	fmt.Println(a["高等数学"]) // 88
	fmt.Println(a["大学物理"]) // 0 不存则不会报错 会返回一个默认值
	// map的遍历
	for key, value := range a { // key value分别接收map的键与值
		fmt.Print(key, ":", value, " ")
	} // Python程序设计:85 高等数学:88
	fmt.Println()
	// 如果只需要遍历键 可以只写个key变量 也不用匿名变量了
	for key := range a {
		fmt.Print(key, " ")
	} // 高等数学 Python程序设计
	fmt.Println()
	// 如果只需要遍历值 那就需要用到匿名变量了
	for _, value := range a {
		fmt.Print(value, " ")
	} // 88 85
	fmt.Println()
	/*  delete()函数删除map中的键值对
	delete()函数使用格式如下
	delete(map变量, 要删除的键名)
	*/
	delete(a, "高等数学")
	fmt.Println(a)    // map[Python程序设计:85]
	delete(a, "大学物理") // 如果删除的对象不存在也不会报错
	fmt.Println(a)    // map[Python程序设计:85]
	// 元素为map类型的切片
	// 定义变量
	/*
		前面学的切片本质是元素为int类型的切片 定义为
		a = make([]int, 0, 5) 这是定义了一个元素为int类型的切片 长度为0 容量为5
		推理 定义元素为map类型的切片如下 map与slice组合
		b = make([]map[string]int, 0, 5) 这是定义了一个元素为map类型的切片 长度为0 容量为5
	*/
	b := make([]map[string]int, 5)
	/*  不合法 b := make([]map[string]int, 5)仅是对切片进行了申请内存 没有对其内部的map数据进行申请内存
	b[0]["高等数学"] = 85
	b[0]["Linux"] = 79
	b[1]["大学物理"] = 70
	b[1]["数据库"] = 70
	*/
	b[0] = make(map[string]int, 2) // 为嵌入切片内部的map分配内存
	b[0]["高等数学"] = 85
	b[0]["Linux"] = 79
	b[1] = make(map[string]int, 2)
	b[1]["大学物理"] = 70
	b[1]["数据库"] = 70
	fmt.Println(b) //  [map[Linux:79 高等数学:85] map[大学物理:70 数据库:70] map[] map[] map[]]
	/*
			前面学的map是键为string/int类型 值为string/int类型 定义为
			c := make(map[string]int, 3)
		推理 声明一个键为string类型 值为切片类型即[]int或[]string类型 定义为
		c := make(map[string][]int, 3)
	*/
	c := make(map[string][]string, 3)
	// 使用键获取值 用切片给值赋值
	c["河南"] = []string{"商丘", "郑州", "南阳"}
	c["江苏"] = []string{"南京", "苏州"}
	fmt.Println(c) // map[江苏:[南京 苏州] 河南:[商丘 郑州 南阳]]
}
