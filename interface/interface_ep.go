package main

import "fmt"

/*
接口实例
销售、行政、程序员都能计算月薪，我们可以把他们员工接口来处理
*/

// 定义接口 管理方法有对每个员工发应发的工资
type worker interface {
	giveMoney()
	// move()
}

// 定义员工 销售 行政 程序员三个结构体  模拟销售与行政与程序员继承员工的属性
type workerInfo struct {
	name  string
	age   int
	money int
}
type salesperson struct {
	info workerInfo
}
type administrativeStaff struct {
	info workerInfo
}
type programmer struct {
	info workerInfo
}

// 分别给三个结构体绑定对应的方法
func (s salesperson) giveMoney() {
	fmt.Println("销售应发工资:", s.info.money)
}
func (a administrativeStaff) giveMoney() {
	fmt.Println("行政应发工资:", a.info.money)
}
func (p programmer) giveMoney() {
	fmt.Println("程序员应发工资:", p.info.money)
}

// 接口类型变量作为函数参数 只要结构体中实现了对应的方法 就是worker类型 就可以作为参数传入这个函数中
func giveMoney(a worker) {
	a.giveMoney()
}
func main() {
	// 定义三个结构体变量
	var s1 = salesperson{
		info: workerInfo{
			name:  "张三",
			age:   20,
			money: 6900,
		},
	}
	var a1 = administrativeStaff{
		info: workerInfo{
			name:  "李四",
			age:   28,
			money: 7500,
		},
	}
	var p1 = programmer{
		info: workerInfo{
			name:  "张三",
			age:   20,
			money: 8900,
		},
	}
	giveMoney(s1) // 销售应发工资: 6900
	giveMoney(a1) // 行政应发工资: 7500
	giveMoney(p1) // 程序员应发工资: 8900
	/*
		结构体必须实现一个接口中的全部方法才可以算是这个接口类的
		# command-line-arguments
		.\interface_ep.go:70:12: cannot use s1 (variable of struct type salesperson) as worker value in argument to giveMoney: salesperson does not implement worker (missing method func2)
		.\interface_ep.go:71:12: cannot use a1 (variable of struct type administrativeStaff) as worker value in argument to giveMoney: administrativeStaff does not implement worker (missing method func2)
		.\interface_ep.go:72:12: cannot use p1 (variable of struct type programmer) as worker value in argument to giveMoney: programmer does not implement worker (missing method func2)
	*/
}
