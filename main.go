package main

import (
	"fmt"
	"math"
)

type person struct {
	firstName string
	lastName  string
}

func main() {
	/**
	函数
		Go 语言至少有一个 main 函数
	语法
		func funcName(param type1, param type2)(output type1, output2 type2){
			return value1, value2
		}
	1:函数由 func 开始声明
	2：funcName：函数名称，函数名和参数列表一起构成了函数签名
	3：参数列表：指定的是参数类型、顺序、及参数个数。参数时可选的，函数也可以不包含参数
	4：返回类型，函数返回一列值
	5：不想声明返回变量，也可以直接就两个类型
	6：如果只有一个返回值且不声明返回值变量，那么你可以省略包括返回值的括号
	7：函数体：函数定义的代码集合

	参数
		可变参
			变量 arg 是一个 int 的 slice
			func myfunc(arg ...int) {}

	作用域
		变量可以使用的范围
	局部变量：一个函数内部定义的变量，就叫做局部变量
	全局变量：一个函数外部定义的变量，就叫做全局变量

	函数的本质
		函数也是 Go 语言中的一种数据类型，可以作为另外一个函数的参数，也可以作为另外一个函数的返回值

	defer 函数
		即延迟语句，延迟语句被用于执行一个函数调用，在这个函数之前，延迟语句返回
		可以在函数添加多个 defer 语句，当函数执行到最后，这些 defer 语句会按照逆序执行，最后该函数返回
		有很多调用 defer，那么 defer 是采用后进先出模式
		在离开所在的方法时，执行（报错的时候也会执行）
		并不局限于函数，延迟一个方法调用也是完全合法的
		延迟函数的参数在执行延迟语句时被执行，而不是在执行实际的函数调用时执行

		注：
			1：当外围函数中语句正常执行完毕时，只有其中所有的延迟函数都执行完毕，外围函数才会真正的结束执行
			2：当执行外围函数中的 return 语句时，只有其中所有的延迟函数都执行完毕后，外围函数才会真正返回
			3：当外围函数中的代码引发运行恐慌时，只有其中所有的延迟函数都执行完毕后，该运行时恐慌才会真正被扩展至调用函数
	*/
	var a int = 100
	var b int = 200
	var ret int

	ret = max(a, b)
	fmt.Printf("最大值： %d\n", ret)

	// 声明函数变量
	getSquareRoot := func(x float64) float64 {
		return math.Sqrt(x)
	}

	// 使用函数 参数传递，值传递
	fmt.Println(getSquareRoot(9))

	/**
	参数引用传递
		变量在内存中是存放于一定地址上的，修改变量实际是修改变量地址处的内存
	*/
	x := 3
	fmt.Println("x = ", x)
	x1 := add1(&x) // 调用 add1(&x) 传 x 的地址
	fmt.Println("x+1 = ", x1)
	fmt.Println("x = ", x)

	nums := []int{78, 109, 2, 563, 300}
	largest(nums)

	p := person{
		firstName: "John",
		lastName:  "Smith",
	}
	defer p.fullName()
	fmt.Printf("Welcome ")

	m := 5
	defer printA(m)
	m = 10
	fmt.Println("value of a before deferred function call", m)

}

func max(num1, num2 int) int {
	var result int

	if num1 > num2 {
		result = num1
	} else {
		result = num2
	}
	return result
}

/**
传指针使得多个函数能操作同一个对象
传指针比较轻量级（8 bytes），只是传内存地址，可以用指针传递体积较大的结构体，如果用参数传递的话，每次在 copy 上面就会花费较多的系统开销
Go 语言中 slice，map 这三种类型的实现机制类似指针，所以可以直接传递，而不用取地址后传递指针（注：若函数需改变 slice 的长度，则仍需要取地址传递指针）
*/
func add1(a *int) int {
	*a = *a + 1
	return *a
}

func finished() {
	fmt.Println("Finished finding largest")
}

func largest(nums []int) {
	defer finished()
	fmt.Println("Started finding largest")
	max := nums[0]
	for _, v := range nums {
		if v > max {
			max = v
		}
	}
	fmt.Println("Largest number in ", nums, " is ", max)
}

func (p person) fullName() {
	fmt.Printf("%s %s", p.firstName, p.lastName)
}

func printA(a int) {
	fmt.Println("value of a in deferred function", a)
}
