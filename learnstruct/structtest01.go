package learnstruct

import (
	"fmt"
	"os"
)

/*
	方法版学生管理系统
	写一个系统能查看\新增\删除学生
	有一个系统：
		1.保存了一些数据  ------------->  结构体字段
		2.有三个功能      ------------->  结构体方法
*/

type student struct {
	id   int64
	name string
}

type ManageSys struct {
	allStudent map[int64]student
}

func newStudent(id int64, name string) *student {
	return &student{
		id:   id,
		name: name,
	}
}

func (m *ManageSys) showMenu(){
	//1.打印菜单
	fmt.Println("--------------------------")
	fmt.Println("欢迎进入学生管理系统")
	fmt.Println(`
		1.查看所有学生
		2.新增学生
		3.删除学生
		4.退出
	`)
	fmt.Print("请输入你要做什么：")
}

func (m *ManageSys) showAllStudent() {
	//把所有的学生打印出来
	for k, v := range m.allStudent {
		fmt.Printf("学号：%d 姓名：%s\n", k, v.name)
	}
}

func (m *ManageSys) addStudent() {
	//向allStudent中添加一个学生
	//1.创建一个学生
	//1.1获取用户输入
	var (
		id   int64
		name string
	)
	fmt.Print("请输入学生学号：")
	fmt.Scanln(&id)
	fmt.Print("请输入学生姓名：")
	fmt.Scanln(&name)
	//1.2造学生
	addStu := newStudent(id, name)
	//2.追加到allStudent中
	for k, _ := range m.allStudent {
		if k == id {
			fmt.Println("学号重复，请重新输入")
			m.addStudent()
		}
	}
	m.allStudent[id] = *addStu
}

func (m *ManageSys) deleteStudent() {
	var deleteId int64
	fmt.Print("请输入删除学生学号：")
	fmt.Scanln(&deleteId)
	for k, _ := range m.allStudent {
		if k == deleteId {
			delete(m.allStudent, deleteId)
			fmt.Println("学号为", deleteId, "的学生已经删除")
		} else {
			fmt.Println("学号不存在")
		}
	}
}

func StructTest01() {
	sys := ManageSys{
		allStudent: make(map[int64]student),
	}
	for {
		//1.进入系统打印菜单
		sys.showMenu()
		//2.等待用户选择要做什么
		var choice int
		fmt.Scanln(&choice)
		if choice <= 4 {
			fmt.Printf("你选择了第 %d 种操作\n", choice)
			fmt.Println("--------------------------")
		}
		//3.执行对应的函数
		switch choice {
		case 1:
			sys.showAllStudent()
		case 2:
			sys.addStudent()
		case 3:
			sys.deleteStudent()
		case 4:
			os.Exit(1)
		default:
			fmt.Println("无效选择……")
		}
	}

}
