package learnvar

//--------------变量定义--------------

/*
//此函数有两个返回值，如果只想接受一个可以像这样：
//  x, _ := LearnVar()
下划线"_" 被称为匿名变量，多用于占位，表示忽略值
*/

func LearnVar() (int,string) {
	return 10, "Foo"
}
