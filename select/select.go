/*
 * @Description: 请输入....
 * @Author: Gavin
 * @Date: 2022-07-16 21:28:48
 * @LastEditTime: 2022-07-16 21:57:21
 * @LastEditors: Gavin
 */
package select_test

import (
	"fmt"
	"time"
)

var num = make(chan int)
var leng = make(chan string)

func Run() {

	go func() {
		num <- 24
		leng <- "周冠宇"
		defer close(num)
		defer close(leng)
	}()
	for {
		select {
		case r := <-leng:
			fmt.Printf("leng: %v\n", r)
		case r := <-num:
			fmt.Printf("num: %v\n", r)
		default:
			fmt.Println("默认行为")

		}
		time.Sleep(time.Second * 2)

	}

}
