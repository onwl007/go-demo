package main

import "fmt"

func main() {
	a := [...]float64{67.7, 89.8, 21, 78}
	for i := 0; i < len(a); i++ {
		fmt.Printf("%d th element of a is %.2f\n", i, a[i])
	}

	sum := float64(0)
	for i, v := range a {
		fmt.Printf("%d the elemrnt of a is %.2f\n", i, v)
		sum += v
	}
	fmt.Println("\nsum of all elements of a ", sum)

	arr := [...]string{"USA", "CHINA", "INDIA", "GERMANY", "FRANCE"}
	arr1 := arr
	arr1[0] = "SINGAPORE"
	fmt.Println("arr is ", arr)
	fmt.Println("arr1 is ", arr1)
}
