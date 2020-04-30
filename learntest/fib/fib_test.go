package fib

import "testing"

//基准测试
func BenchmarkFib(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Fib(2)
	}
}

//内部调用性能比较调用
func benchmarkFib(b *testing.B, n int) {
	for i := 0; i < b.N; i++ {
		Fib(n)
	}
}

//
func BenchmarkFib2(b *testing.B) {
		benchmarkFib(b, 2)
}

func BenchmarkFib20(b *testing.B) {
	benchmarkFib(b, 20)
}

//跑不动
//func BenchmarkFib30(b *testing.B) {
//	benchmarkFib(b, 200)
//}