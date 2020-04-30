package learnmap

import "fmt"

//------------- map -----------------
/*
	map 是一种无序的基于 key-value 的数据结构。
	Go语言中的map是引用类型，必须初始化才能使用。
	map 定义语法如下：
		map[KeyType]ValueType
		其中：
			KeyType：表示键的类型
			ValueType：表示键对应的值得类型
	map类型的变量默认初始值为 nil， 需要使用 make() 函数来分配内存。语法为：
		make(map[KeyType]ValueType, [cap])
		其中： cap 表示 map 的容量，该参数不是必须的，但是应该在初始化 map 的时候就为其指定一个合适的容量

	map 的遍历用 for range
*/

func PrintMap01() {
	var m1 map[string]int
	fmt.Println(m1 == nil) //还没有在内存中开辟空间，不能初始化赋值

	//map的初始化
	m2 := make(map[string]int, 3)
	m2["俊林"] = 33
	m2["巧巧"] = 32
	m2["修凡"] = 2
	fmt.Println(m2)

	m3 := map[string]int{"hhh":23,"jjj":32}
	fmt.Println(m3)


	//map 遍历
	for key, value := range m2 {
		fmt.Println("key:", key, "value:", value)
	}

	//只想遍历 key 值，如下：
	for key := range m2 {
		fmt.Println("key:", key)
	}

	//只想遍历 value 值，如下：
	for _, value := range m2 {
		fmt.Println("value:", value)
	}

}
