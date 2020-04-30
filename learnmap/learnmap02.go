package learnmap

import "fmt"

//-----------------delete()------------
/*
	使用 delete() 内建函数从 map 中删除一组键值对，格式如下：
		delete(map, key)
		其中：
			map：表示要删除键值对的 map
			key：表示要删除的键值对的键
 */

func PrintMap02(){
	m1 := make(map[string]int, 3)
	m1["俊林"] = 33
	m1["巧巧"] = 32
	m1["修凡"] = 2
	delete(m1,"巧巧")
	for key, value := range m1 {
		fmt.Println("key:", key, "value:", value)
	}


}