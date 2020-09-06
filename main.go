package main

import "fmt"

func main() {
	/**
	Map 无序的键值对集合，引用类型
	注意
		1：map 是无序的，不能通过 index 获取，必须通过 key 获取
		2：长度不固定，引用类型
		3：内置的 len 函数同样适用于 map，返回 map 拥有的 key 的数量
		4：map 的 key 可以是所有可比较的类型，如布尔型、整数型、浮点型、复杂型、字符串型
	*/
	var countryCapitalMap map[string]string
	countryCapitalMap = make(map[string]string)

	// map 插入 key-value 对，各个多家对应的首都
	countryCapitalMap["France"] = "Paris"
	countryCapitalMap["Italy"] = "Rome"
	countryCapitalMap["Japan"] = "Tokyo"
	countryCapitalMap["India"] = "New Delhi"

	for country := range countryCapitalMap {
		fmt.Println("Capital of", country, "is", countryCapitalMap[country])
	}

	// 查看元素在集合中是否存在
	captial, ok := countryCapitalMap["United States"]
	if ok {
		fmt.Println("Capital of United States is ", captial)
	} else {
		fmt.Println("Capital of United States is not present")
	}

	delete(countryCapitalMap, "France")
	fmt.Println("Entry for France is deleted")

	fmt.Println("删除元素后 map")

	for country := range countryCapitalMap {
		fmt.Println("Capital of ", country, "is ", countryCapitalMap[country])
	}

	m := make(map[string]int)
	m["a"] = 1
	x, ok := m["b"]
	fmt.Println(x, ok)
	x, ok = m["a"]
	fmt.Println(x, ok)

	// map 是引用类型，将映射分配给一个新变量时，它们都指向相同的内部数据结构，因此，一个的变化会反映另一个
	// map 不能使用 == 操作符进行比较。== 只能用来检查 map 是否为空。否则会报错
	personSalary := map[string]int{
		"steve": 12000,
		"jamie": 15000,
	}
	personSalary["mike"] = 9000
	fmt.Println("Original person salary", personSalary)
	newPersonSalary := personSalary
	newPersonSalary["mike"] = 18000
	fmt.Println("Person salary changed", personSalary)
}
