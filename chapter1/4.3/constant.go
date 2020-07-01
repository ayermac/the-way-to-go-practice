/**
 * 常量定义
 */
package main

import (
	"fmt"
)


// 常量pi
const Pi = 3.14159
// 显示定义
const abc string = "abc"
// 隐式定义
//const b = "abc"
// 常量的值必须是能够在编译时就能够确定的；你可以在其赋值表达式中涉及计算过程，但是所有用于计算的值必须在编译期间就能获得。
const Ln2= 0.693147180559945309417232121458176568075500134360255254120680009
const Log2E= 1/Ln2 // this is a precise reciprocal
const Billion = 1e9 // float constant
const hardEight = (1 << 100) >> 97

//const (
//	Monday, Tuesday, Wednesday = 1, 2, 3
//	Thursday, Friday, Saturday = 4, 5, 6
//)

const (
	Sunday = iota
	Monday
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
)

const (
	a = iota// 0
	b = iota// 1
	c = iota// 2
)

func main()  {
	fmt.Println(Pi, b, Log2E, Billion, hardEight)
	fmt.Println(Sunday, Monday, Tuesday, Wednesday, Thursday, Friday, Saturday)
	fmt.Println(a, b, c)
}
