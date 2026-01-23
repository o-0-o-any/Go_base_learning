package main

import "fmt"

func main() {
	// 算数运算符
	n1 := 6
	n2 := 7
	var (
		n3 int
		n4 int
		n5 int
		n6 float64
		n7 int
	)
	// 加减乘除
	n3 = n1 + n2
	n4 = n1 - n2
	n5 = n1 * n2
	n6 = float64(n1 / n2)
	n7 = n2 % n1
	fmt.Println(n3, n4, n5, n6, n7) // 13 -1 42 0 1
	// 自加自减  属于单独的语句 并不是运算符
	n1++
	n2--
	fmt.Println(n1, n2) // 7 6
	// 关系运算符  只有相同类型的变量看可以比较关系
	var (
		b1 bool
		b2 bool
		b3 bool
		b4 bool
		b5 bool
		b6 bool
	)
	// n1 = 7 n2 = 6
	b1 = n1 == n2                       // false
	b2 = n1 != n2                       // true
	b3 = n1 > n2                        // true
	b4 = n1 < n2                        // false
	b5 = n2 >= n1                       // false
	b6 = n2 <= n1                       // true
	fmt.Println(b1, b2, b3, b4, b5, b6) // false true true false false true
	// 逻辑运算符
	var (
		b7  bool
		b8  bool
		b9  bool
		b10 bool
		b11 bool
		b12 bool
	)
	b7 = true
	b8 = false
	b9 = b7 && b8                  // false
	b10 = b7 || b8                 // true
	b11 = !b7                      // false
	b12 = !b8                      // true
	fmt.Println(b9, b10, b11, b12) // false true false true
	// 位运算符 针对二进制数字 比如 5 & 2  5的二进制数字为101 2的二进制数字为10 各位相比较 最后的结果为000 5 | 2 就是111
	fmt.Println(5 & 2)           // 000
	fmt.Println(5 | 2)           // 111
	fmt.Println(5 ^ 2)           // 111  对应位置上数不一样为1 一样为0
	fmt.Println(1 << 2)          // 4  1 << 2 1向左移2位即100 即是2的2次方 1左移几位就是2的几次方
	fmt.Println(float32(1 >> 2)) // 0.25  1 >> 2 1向右移2位即0.01 即是2的-2次方 1右移几位就是2的负几次方
	// 赋值运算符
	var (
		a1 int
		a2 int
	)
	a1 = 10             // =
	a2 = 20             // =
	fmt.Println(a1, a2) // 10 20
	a2 += a1            // 30 a2 = a2 + a1
	fmt.Println(a2)     // 30
	a2 -= a1            // 20 a2 = a2 - a1
	fmt.Println(a2)     // 20
	a2 *= a1            // 200 a2 = a2 * a1
	fmt.Println(a2)     // 200
	a2 /= a1            // 20 a2 = a2 / a1
	fmt.Println(a2)     // 20
	a3, a4 := 10, 9
	a3 %= a4 // 1  a3 = a3 % a4
	fmt.Println(a3)
	// 同样的 位运算符也可以这样使用 <<= >>= &= |= ^= 这里就不写了
}
