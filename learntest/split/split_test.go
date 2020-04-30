package split

import (
	"fmt"
	"os"
	"reflect"
	"testing"
)

// -------------------- 测试组及覆盖率 ---------------------
/*
	测试组：避免反复定义测试函数
	覆盖率：
		将覆盖率的日志文件输出到文件： go test -cover -coverprofile=c.out
		使用浏览器以HTML方式打开上一步生成的文件：go tool cover -html=c.out

*/

// 测试集的Setup与Teardown
func setupTestCase(t *testing.T) func(t *testing.T) {
	t.Log("如有需要在此执行:测试之前的setup")
	return func(t *testing.T) {
		t.Log("如有需要在此执行:测试之后的teardown")
	}
}

// 子测试的Setup与Teardown
func setupSubTest(t *testing.T) func(t *testing.T) {
	t.Log("如有需要在此执行:子测试之前的setup")
	return func(t *testing.T) {
		t.Log("如有需要在此执行:子测试之后的teardown")
	}
}

func TestGroupSplit(t *testing.T) {
	type test struct {
		str  string
		sep  string
		want []string
	}

	//测试组初始化
	tests := map[string]test{
		"normal": test{"a:b:c", ":", []string{"a", "b", "c"},},
		"none":   test{"a:b:c", "*", []string{"a:b:c"},},
		"multi":  test{"abcdebcggbcjk", "bc", []string{"a", "de", "gg", "jk"},},
	}

	teardownTestCase := setupTestCase(t) // 测试之前执行setup操作
	defer teardownTestCase(t)            // 测试之后执行testdoen操作

	for k, v := range tests {
		//got := Split(v.str, v.sep)
		//if ok := reflect.DeepEqual(got, v.want); !ok {
		//	t.Fatalf("测试用例：%v 失败，期望得到：%#v，实际得到：%#v", k, v.want, got)
		//}
		t.Run(k, func(t *testing.T) {
			teardownTestCase := setupTestCase(t) // 测试之前执行setup操作
			defer teardownTestCase(t)            // 测试之后执行testdoen操作
			got := Split(v.str, v.sep)
			if !reflect.DeepEqual(got, v.want) {
				t.Fatalf("期望得到：%#v，实际得到：%#v", v.want, got)
			}
		})
	}
}

//基准测试
/*
	运行格式：
		go test -bench=函数名正则表达 -cpu=2 -benchmem
	其中：
		cpu=2 表示2核心运行
		-benchmem 表示显示内存信息
*/
func BenchmarkSplit(b *testing.B) {
	//b.Log("这是基准测试")
	for i := 0; i < b.N; i++ {
		Split("a:b:c", ":")
	}
}

//并行测试：运行命令和基准测试一样
//自动调用所有核心并行运行测试
func BenchmarkSplitParallel(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			Split("a:b:c", ":")
		}
	})
}

//Setup与TearDown
func TestMain(m *testing.M) {
	fmt.Println("write setup code here...") // 测试之前的做一些设置
	// 如果 TestMain 使用了 flags，这里应该加上flag.Parse()
	retCode := m.Run()                         // 执行测试
	fmt.Println("write teardown code here...") // 测试之后做一些拆卸工作
	os.Exit(retCode)                           // 退出测试
}


//示例函数
func ExampleSplit() {
	fmt.Println(Split("a:b:c",":"))
	// OutPut: [a b c]
}