package learnconst

import "fmt"

//-----------常量----------

//常量定义好之后不能修改
//在程序运行期间不会改变
const pi = 3.1415926

//批量声明常量
const (
	statusOK = 200
	notFound = 404
)

//批量声明常量，下面默认和第一个同值
const (
	n1 = 100
	n2 //相当于 n2 = 100
	n3 //相当于 n2 = 100
)

//iota 在const关键字出现时被重置为0
//const中每新增一行常量声明将使iota计数一次（iota可以理解为const语句块中的行索引）
//使用iota可以简化定义，在定义枚举时很有用
const (
	a1 = iota // a1 = 0
	a2        //相当于a2 = iota = 1
	a3        //相当于a3 = iota = 2
	a4        //相当于a4 = iota = 3
)

//关于iota的例题1
const (
	b1 = iota // 0
	b2        // 1
	_         // 2
	b3        // 3
)

//关于iota的例题2
const (
	c1 = iota // 0
	c2 = 100  // 100
	c3 = iota // 2
	c4        // 3
)
const c5 = iota // 0

//关于iota的例题3
const (
	d1, d2 = iota + 1, iota + 2 //1, 2
	d3, d4 = iota + 1, iota + 2 //2, 3
)

//关于iota的例题4:定义数量级
const (
	_  = iota
	KB = 1 << (10 * iota)
	MB = 1 << (10 * iota)
	GB = 1 << (10 * iota)
	TB = 1 << (10 * iota)
	PB = 1 << (10 * iota)
)

//关于iota的例题4
const (
	a, b = iota + 1, iota + 2 //1,2
	c, d                      //2,3
	e, f                      //3,4
)

func PrintConst() {
	fmt.Println("pi:", pi, "statusOK:", statusOK, "notFound:", notFound, "n1:", n1, "n2:", n2, "n3:", n3, "a1:", a1, "a2:", a2, "a3:", a3, "a4:", a4)
	fmt.Println("KB:", KB, "MB:", MB, "KB:", GB, "KB:", GB, "TB:", TB, "PB:", PB)
	fmt.Println("b1:", b1, "b2:", b2, "b3:", b3, "c1:", c1, "c2:", c2, "c3:", c3, "c4:", c4, "d1:", d1, "d2:", d2, "d3:", d3, "d4:", d4)
	fmt.Println("a:", a, "b:", b, "c:", c, "d:", d, "e:", e, "f:", f)
}
