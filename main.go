package main

import (
	"fmt"
	"math"
)

type areaError struct {
	err    string
	radius float64
}

func (e *areaError) Error() string {
	return fmt.Sprintf("radius %0.2f: %s", e.radius, e.err)
}

func main() {
	// f, err := os.Open("/test.txt")
	// if err, ok := err.(*os.PathError); ok {
	// 	fmt.Println("File at path", err.Path, "failed to open")
	// 	return
	// }
	// fmt.Println(f.Name(), "opened successfully")

	// addr, err := net.LookupHost("golangbot123.com")
	// if err, ok := err.(*net.DNSError); ok {
	// 	if err.Timeout() {
	// 		fmt.Println("operation timed out")
	// 	} else if err.Temporary() {
	// 		fmt.Println("temporary error")
	// 	} else {
	// 		fmt.Println("generic error", err)
	// 	}
	// 	return
	// }
	// fmt.Println(addr)

	// files, error := filepath.Glob("[")
	// if error != nil && error == filepath.ErrBadPattern {
	// 	fmt.Println(error)
	// 	return
	// }
	// fmt.Println("matched files", files)

	// 自定义错误
	// 可以使用 errors 包下的 New() 函数，以及 fmt 包下的 Errorf() 函数
	radius := -20.0
	area, err := circleArea(radius)
	if err != nil {
		if err, ok := err.(*areaError); ok {
			fmt.Printf("Radius %0.2f is less than zero", err.radius)
			return
		}
		fmt.Println(err)
		return
	}
	fmt.Printf("Area of circle %0.2f", area)
}

func circleArea(radius float64) (float64, error) {
	if radius < 0 {
		return 0, &areaError{"radius is negative", radius}
	}
	return math.Pi * radius * radius, nil
}
