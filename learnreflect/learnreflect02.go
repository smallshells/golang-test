package learnreflect

import (
	"fmt"
	"reflect"
)

// --------------------------------- 反射 -----------------------------
/*
	reflect.ValueOf() 返回的是 reflect.Value 类型，其中包含了原始值的值信息。
	reflect.Value 与原始值之间可以互相转换。
	reflect.Value 类型提供的获取原始值的方法如下：
		Inerface() interface{}  将值以interfa{}类型返回，可以通过类型断言转换为指定类型
		Int()      int64    	将值以int类型返回，所有有符号整型均可以此方式返回
		Uint()     uint64   	将值以uint类型返回，所有无符号整型均可以此方式返回
		Float()    float64  	将值以双精度（float64）类型返回，所有浮点数（float32、float64）均可以此方式返回
		Bool() 	   bool  		将值以bool类型返回
		Bytes()	   []bytes  	将值以字节数组[]bytes 类型返回
		String()   string  		将值以字符串类型返回
*/

func reflectValue(x interface{}) {
	v := reflect.ValueOf(x)
	k := v.Kind()
	switch k {
	case reflect.Int64:
		// v.Int()从反射中获取整形的原始值，然后通过int64()强制类型转换
		fmt.Printf("type is int64, value is %d\n", int64(v.Int()))
	case reflect.Float32:
		// v.Float()从反射中获取整形的原始值，然后通过float32()强制类型转换
		fmt.Printf("type is int64, value is %f\n", float32(v.Float()))
	case reflect.Float64:
		// v.Float()从反射中获取整形的原始值，然后通过float64()强制类型转换
		fmt.Printf("type is int64, value is %f\n", float64(v.Float()))
	}
}

func TestReflect02() {
	var a float32 = 3.14
	var b int64 = 100
	reflectValue(a)
	reflectValue(b)
	//将int类型的原始值转换成reflect.Value类型
	c := reflect.ValueOf(10)
	fmt.Printf("type c :%T, value c :%v\n", c, c)
}
