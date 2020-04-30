package learnchannel

import (
	"fmt"
	"math/rand"
	"time"
)

//----------------- 练习题：生产者消费之模型 --------------------
/*
	使用goroutine和channel实现一个简易的生产者模型
*/

type item struct {
	id  int64
	num int64
}

type result struct {
	item *item
	sum  int64
}

func producer(itemChan chan *item) {
	//1.生成随机数
	var id int64
	for {
		id++
		number := rand.Int63()
		tmp := &item{
			id:  id,
			num: number,
		}
		//2.把随机数发送到通道中
		itemChan <- tmp
	}
}

//计算一个数字每个位的和
func calcSum(num int64) (sum int64) {
	for num > 0 {
		sum += num % 10
		num /= 10
	}
	return
}

func consumer(itemChan chan *item, resultChan chan *result) {
	for tmp := range itemChan {
		//(*item).num  //item.num
		sum := calcSum(tmp.num)
		//根据 result结构体
		retObj := &result{
			item: tmp,
			sum:  sum,
		}
		resultChan <- retObj
	}
}

func startWorker(n int, itemChan chan *item, resultChan chan *result) {
	for i := 0; i < n; i++ {
		go consumer(itemChan, resultChan)
	}
}

func printResult(resultChan chan *result) {
	for ret := range resultChan {
		fmt.Printf("id : %v , num : %v , sum : %v\n", ret.item.id, ret.item.num, ret.sum)
		time.Sleep(time.Second)
	}
}

func TestChannel04() {
	itemChan := make(chan *item, 100)
	resultChan := make(chan *result, 100)

	go producer(itemChan)
	startWorker(20, itemChan, resultChan)
	printResult(resultChan)
}
