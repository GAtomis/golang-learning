/*
 * @Description: 并发通道
 * @Author: Gavin
 * @Date: 2022-07-15 14:07:28
 * @LastEditTime: 2022-07-15 18:42:53
 * @LastEditors: Gavin
 */
package channel

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

//创造一个无缓冲的通道
var values = make(chan int)

func send() {
	rand.Seed(time.Now().UnixNano())
	value := rand.Intn(10)
	fmt.Printf("send: %v\n", value)
	time.Sleep(time.Second * 5)
	values <- value
}

func normal() {
	defer close(values)
	go send()
	fmt.Println("wait...,等待通道内载入数据")
	value := <-values
	fmt.Printf("receive: %v\n", value)
	fmt.Println("end....")
}
func Run() {
	useFor()
}

//通过循环检测通道是否打开
func myfun(mychnl chan string) {

	for v := 0; v < 4; v++ {
		mychnl <- "nhooo" + strconv.Itoa(v)
	}
	defer close(mychnl)
}

//关于通道的遍历
func useFor() {

	//创建通道
	c := make(chan string)

	// 使用 Goroutine
	go myfun(c)

	//当ok的值为为true时，表示通道已打开，可以发送或接收数据
	//当ok的值设置为false时，表示通道已关闭
	for {
		res, ok := <-c
		if ok == false {
			fmt.Println("通道关闭 ", ok)
			break
		}
		fmt.Println("通道打开 ", res, ok)
	}

}
