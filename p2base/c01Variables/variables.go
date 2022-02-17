package c01Variables

/*
golang variables
*/

// 通用的声明方式
var a int = 10

//block块多变量声明
var (
	b1 int    = 10
	b2 int8   = 6
	b3 string = "hello"
	b4 rune   = 'A'
	b5 bool   = true
)

// 一行声明多个变量
var x, y, z = 0, 0, 0

//海象运算符只能在方法内部使用,局部变量
