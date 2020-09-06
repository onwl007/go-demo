package main

import "fmt"

type name int8
type first struct {
	a int
	b bool
	name
}

func main() {
	/**
	变量
		是一种使用方便的占位符，用于引用计算机内存地址
	指针
		是存储另一个变量内存地址的变量
	声明指针
		var var_name *var_type
		var ip *int 指向整型
		var fp *float32 指向浮点型
	& 取址符
	获取指针的值，语法：*a

	不要将一个指向数组的指针传递给函数，使用切片

	指针的指针
		一个指针变量存放的又是另一个指针变量的地址，则称这个指针变量为指向指针的指针变量
		var ptr **int
	*/
	var a int = 20
	var ip *int

	ip = &a // 指针变量的存储地址
	fmt.Printf("a 变量的地址是：%x\n", &a)

	// 指针变量的存储地址
	fmt.Printf("ip 变量的存储地址：%x\n", ip)

	// 使用指针访问值
	fmt.Printf("*ip 变量的值：%d\n", *ip)

	var s = first{1, false, 2}
	var b *first = &s
	fmt.Println(s.b, s.a, s.name, &s, b.a, &b, (*b).a)

	m := 255
	n := &m
	fmt.Println("address of m is ", n)
	fmt.Println("value of m is ", *n)
	*n++
	fmt.Println("new value of m is ", m)

	x := 58
	fmt.Println("value of x before function call is ", x)
	y := &x
	change(y)
	fmt.Println("value of x after function call is ", x)

	var c int
	var ptr *int
	var pptr **int
	c = 3000
	// 指针 ptr 地址
	ptr = &c
	// 指向指针 ptr 地址
	pptr = &ptr
	// 获取 pptr 的值
	fmt.Printf("变量 c = %d\n", c)
	fmt.Printf("指针变量 *ptr = %d\n", *ptr)
	fmt.Printf("指向指针的指针变量 **ptr = %d\n", **pptr)

	var e int = 100
	var f int = 200
	fmt.Printf("交换前 e 的值 : %d\n", e)
	fmt.Printf("交换前 f 的值 : %d\n", f)

	/* 调用函数用于交换值
	 * &a 指向 a 变量的地址
	 * &b 指向 b 变量的地址
	 */
	swap(&e, &f)
	fmt.Printf("交换后 e 的值 : %d\n", e)
	fmt.Printf("交换后 f 的值 : %d\n", f)
}

func change(val *int) {
	*val = 55
}

func swap(x, y *int) {
	var temp int
	temp = *x
	*x = *y
	*y = temp
}
