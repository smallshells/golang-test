package learnreflect

import (
	"fmt"
	"reflect"
)

// --------------------------------- 反射 -----------------------------
/*
	类型太多了，类型断言猜不全，使用反射就能直接拿到接口值的动态类型和动态值
	任意接口值在反射中都可以理解为由reflect.Type 和 reflect.Value 两部分组成
	reflect包提供了 reflect.TypeOf 和 reflect.ValueOf 两个函数来获取任意对象的 Type 和 Value
	应用：各种web框架、配置文件解析库、ORM框架
	反射是把双刃剑：
		优点：让代码更加灵活
		缺点：降低了代码运行效率
	TypeOf:
		在Go语言中，使用reflect.TypeOf() 函数可以获得任意值的类型对象（reflect.Type）,程序通过类型对象
		可以访问任意值的类型信息
	在反射中关于类型划分为两种：类型（Type） 和 种类（kind）。
	因为在Go语言中我们可以使用type关键字构造很多自定义类型，而种类（kind）就是指底层的类型。
	在反射中，当需要区分指针、结构体等品种的类型时，就会用到 种类（kind）。
*/

func reflectType01(x interface{}) {
	t := reflect.TypeOf(x)
	fmt.Printf("type:%v\n", t)
}

func reflectType02(x interface{}) {
	t := reflect.TypeOf(x)
	fmt.Printf("type:%v kind:%v\n", t.Name(), t.Kind())
}

type myInt int64

func TestReflect01() {
	var a float32 = 3.14
	reflectType01(a)
	var b int64 = 100
	reflectType01(b)
	reflectType01(map[string]int{"jun": 23})

	var c *float32
	var d myInt
	var e rune
	reflectType02(c) //指针类型type为空
	reflectType02(d)
	reflectType02(e)

	type person struct {
		name string
		age  int
	}
	type book struct{ title string }
	var p = person{
		name: "hello",
		age:  20,
	}
	var k = book{title: "《go语言核心编程》"}
	reflectType02(p)
	reflectType02(k)
	reflectType02([]int{1, 2, 3})   //复合类型type为空
	reflectType02(map[string]int{}) //复合类型type为空
}
