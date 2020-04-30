package learnslice

import "fmt"

//----------------基于数组的切片---------------
/*
	s1 := a1[n1:n2]   切片s1基于数组a1，元素包括a1[n1],a1[n1+1],……,a1[n2-2],a1[n2-1]
	还支持如下方式：
		s2 := a1[n1:]
		s3 := a1[:n2]
		s4 := a1[:]
	还可以基于切片再次得到切片：操作类似于从数组得到切片
*/

func PrintSlice02() {
	//由数组得到切片
	a1 := [...]int{1, 4, 2, 4, 6, 2,23,4,5}
	s1 := a1[0:4] //得到一个由数组a1下标 0 到 4 的值组成的切片:左包含右不包含
	s2 := a1[1:]
	s3 := a1[:5]
	s4 := a1[:]
	s5 := a1[2:5]
	fmt.Println(s1, s2, s3, s4, s5)

	//切片的容量是指底层数组从切面第一个元素到最后一个元素的长度
	fmt.Println("s1切片长度：",len(s1),"\ts1切片容量：",cap(s1))
	fmt.Println("s2切片长度：",len(s2),"\ts2切片容量：",cap(s2))
	fmt.Println("s5切片长度：",len(s5),"\ts5切片容量：",cap(s5))

	//切片再切片
	a2 := [...]string{"拉萨","阿里","西安","汉中","成都","延安"}
	fmt.Printf("a2:%v type:%T len:%d cap:%d\n",a2,a2,len(a2),cap(a2))
	b := a2[1:3]
	fmt.Printf("b:%v type:%T len:%d cap:%d\n",b,b,len(b),cap(b))
	c := b[1:5]
	fmt.Printf("c:%v type:%T len:%d cap:%d\n",c,c,len(c),cap(c))
}
