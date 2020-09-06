package main

import "fmt"

func main() {
	/**
	切片是对数组的抽象，数组的长度是不可变的，切片本身没有任何数据，只是对现有数组的引用
	切片像一个结构体，包含三个元素
	1.指针，指向数组中 slice 指定的开始位置
	2.长度，即 slice 的长度
	3.最大长度，也就是 slice 开始位置到数组的最后位置的长度

	当多个切片共享相同的底层数组时，每个元素所做的更改将在数组中反映出来

	语法
		var identifier []type
	使用 make() 函数创建切片
		var slice1 []type = make([]type, len)
		slice1 := make([]type, len)
	*/
	a := [5]int{76, 77, 78, 79, 80}
	var b []int = a[1:4] // 前闭后开
	fmt.Println(b)

	darr := [...]int{57, 89, 90, 82, 100, 78, 67, 69, 59}
	dslice := darr[2:5] // 90, 82, 100
	fmt.Println("array before", darr)
	for i := range dslice {
		dslice[i]++
	}
	fmt.Println("array after", darr)

	// 多个切片共享相同的数组时，每个切片改变元素的值，数组中都会改变
	numa := [3]int{78, 79, 80}
	nums1 := numa[:] // 创建一个切片包含所有元素
	nums2 := numa[:]
	fmt.Println("array befire change 1 ", numa)
	nums1[0] = 100
	fmt.Println("array after modification to slice nums1", numa)
	nums2[1] = 101
	fmt.Println("array after modification to slice nums2", numa)

	/**
	切片
		长度：切片中元素的数量
		容量：从创建切片的索引开始的底层数组中元素的数量
	*/
	var slice2 = make([]int, 3, 5)
	printSlice(slice2)

	var numbers []int
	printSlice(numbers)

	// 允许追加空切片
	numbers = append(numbers, 0)
	printSlice(numbers)

	numbers = append(numbers, 1)
	printSlice(numbers)

	numbers = append(numbers, 2, 3, 4)
	printSlice(numbers)

	// 创建切片 numbers1 是之前切片的两倍容量
	numbers1 := make([]int, len(numbers), (cap(numbers))*2)

	// 拷贝 numbers 的内容到 numbers1
	copy(numbers1, numbers)
	printSlice(numbers1)

}

func printSlice(x []int) {
	fmt.Printf("len = %d cap = %d slice = %v\n", len(x), cap(x), x)
}
