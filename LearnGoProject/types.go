package main

import (
	"fmt"
	"unsafe"
)

/**
主要是练习golang语言中的基础数据类型：
1. boolean即：bool
2.int
3.string
*/
func main() {
	var flag bool
	fmt.Println("打印bool的默认值", flag)
	if flag {
		fmt.Println("flag is true")
	} else {
		fmt.Println("flag is false")
	}
	var flag1 bool = true
	flag2 := true
	const flag3 = true
	fmt.Println("不同的bool类型声明如下：", flag1, flag2, flag3)

	var (
		i1 int  //int either 32 or 64 bits
		i2 int8 //
		i3 int16
		i4 int32
		i5 int64
	)
	var (
		i6  uint //uint either 32 or 64 bitsa
		i7  uint8
		i8  uint16
		i9  uint32
		i10 uint64
		a   string = "我"
	)
	//unsafe.Sizeof: returns the size in bytes
	//1 汉字 = 2 byte = 16 bit
	fmt.Println("打印出汉字'我'实际所占存储空间长度：", unsafe.Sizeof(a))
	fmt.Println("打印出int实际所占存储空间长度：", unsafe.Sizeof(i1))
	fmt.Println("打印出int8实际所占存储空间大小：", unsafe.Sizeof(i2))
	fmt.Println("打印出int16实际所占存储空间大小：", unsafe.Sizeof(i3))
	fmt.Println("打印出int32实际所占存储空间大小：", unsafe.Sizeof(i4))
	fmt.Println("打印出int64实际所占存储空间大小：", unsafe.Sizeof(i5))
	fmt.Println("打印出uint实际所占存储空间大小：", unsafe.Sizeof(i6))
	fmt.Println("打印出uint8实际所占存储空间大小：", unsafe.Sizeof(i7))
	fmt.Println("打印出uint16实际所占存储空间大小：", unsafe.Sizeof(i8))
	fmt.Println("打印出uint32实际所占存储空间大小：", unsafe.Sizeof(i9))
	fmt.Println("打印出uint64实际所占存储空间大小：", unsafe.Sizeof(i10))
	//字 word 、字节 byte 、位 bit(来自英文bit，音译为“比特”，表示二进制位)
	//字长是指字的长度
	//1字=2字节(1 word = 2 byte)|一个字的字长为16
	//1字节=8位(1 byte = 8bit)|一个字节的字长是8
	//1 Byte = 8 Bits
	//1 KB = 1024 Bytes
	//1 MB = 1024 KB
	//1 GB = 1024 MB

	//string
	var str string = "hello , world"
	fmt.Println(str)
	str = "hello , \"world\"" //特殊字符转意
	fmt.Println(str)
	str = `hello , "world"`
	fmt.Println(str)
	str = `line1
	line2
	line3`
	fmt.Println(str)
	str = "hello,world,世界"
	fmt.Println([]byte(str))        //把字符串转换成byte
	fmt.Println(len(str))           //查看byte的大小
	fmt.Println(len(([]rune(str)))) //查看多少个字符
	//输出每一个字符
	for _, char := range str {
		//输出为range int32类型
		fmt.Println(char)
		//range类型是int32,需要转换成string
		fmt.Println(string(char))
	}
}