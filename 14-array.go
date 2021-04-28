package main

import "fmt"

func main() {
	// 类型 [n]T 表示拥有 n 个 T 类型的值的数组。
	var a [2]string
	a[0] = "hello"
	a[1] = "world"
	fmt.Println(a[0], a[1])
	fmt.Println(a)
}
