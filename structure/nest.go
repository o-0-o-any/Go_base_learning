package main

import "fmt"

// 定义多个结构体
type address struct {
	province string
	city     string
}
type worker struct {
	id     int
	name   string
	age    int
	gender string
	adr    address // 结构体嵌套
}
type company struct {
	name    string
	adr     address
	workers []worker // 结构体嵌套
}

func main() {
	// 嵌套结构体的初始化
	w1 := worker{
		id:     1,
		name:   "John Smith",
		age:    25,
		gender: "male",
		adr: address{
			province: "江苏",
			city:     "苏州",
		},
	}
	fmt.Println(w1) // {1 John Smith 25 male {江苏 苏州}}
	// 嵌套结构体中字段的访问 结构体变量名+.+嵌套结构体类型定义的变量名+.+嵌套结构体字段名
	fmt.Println(w1.adr.province, w1.adr.city) // 江苏 苏州

	c1 := company{
		name: "Go后端开发",
		adr: address{
			province: "南极洲",
			city:     "水下",
		},
		workers: []worker{
			w1,
			{
				id:     2,
				name:   "New York",
				age:    30,
				gender: "male",
				adr: address{
					province: "河南",
					city:     "商丘",
				},
			},
		},
	}
	fmt.Println(c1) // {Go后端开发 {江苏 苏州} [{1 John Smith 25 male {江苏 苏州}} {2 New York 30 male {河南 商丘}}]}
	// 嵌套结构体中字段的访问
	fmt.Println(c1.workers[0].name, c1.workers[1].name) // John Smith New York
	// 直接访问字段名 规则与全局/局部变量类似
	// 先在自己结构体中寻找有没有该字段名 如果没有继续向嵌套结构体中寻找
}
