package main

import "fmt"

func main() {
	var a string = "shen"
	var b int = 4
	var (
		c = 8
		d int = 9
		e int8 = 10
	)
	fmt.Println("一个string类型的变量a的值为："+a)
	fmt.Println("一个int类型的变量b的值为：",b)
	fmt.Println("一个int类型的变量c的值为：",c)
	fmt.Println("一个int类型的变量d的值为：",d)
	fmt.Println("一个int8类型的变量e的值为：",e)
}
