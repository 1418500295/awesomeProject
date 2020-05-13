package main

import "fmt"
/**
常量用法
 */
func main() {
	const a  = 12
	const (
		c = 1
		b = 2
		d = 4
	)
	/**
	iota特殊常量，可认为可以被编译器修改
	ota 在 const关键字出现时将被重置为 0(const 内部的第一行之前)，
	const 中每新增一行常量声明将使 iota 计数一次(iota 可理解为 const 语句块中的行索引)。
	 */
	const (
		e = iota
		f = iota
		g
	)

	fmt.Print(a)
	fmt.Print(c)
	fmt.Print(e,f,g)

}
