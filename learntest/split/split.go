package split

import "strings"

// 按照要求切割字符串
//BenchmarkSplit-12        6365307         188 ns/op      112 B/op      3 allocs/op
//func Split(s ,sep string) (result []string) {
//	index := strings.Index(s,sep)
//	for index > 0 {
//		result = append(result, s[:index])
//		s = s[index+len(sep) :]
//		index = strings.Index(s,sep)
//	}
//	result = append(result, s)
//	return
//}

//性能优化版本
//BenchmarkSplit-12       13329393         87.3 ns/op       48 B/op      1 allocs/op
func Split(s, sep string) []string {
	count := strings.Count(s, sep)
	result := make([]string, 0, count+1)
	index := strings.Index(s, sep)
	for index > 0 {
		result = append(result, s[:index])
		s = s[index+len(sep):]
		index = strings.Index(s, sep)
	}
	result = append(result, s)
	return result
}
