/*
 * @Description: 请输入....
 * @Author: Gavin
 * @Date: 2022-07-05 14:14:18
 * @LastEditTime: 2022-07-05 14:51:54
 * @LastEditors: Gavin
 */
package goroutines

import (
	"fmt"
	"time"
)

//协程
func Run() {
	go showMessage("java")
	showMessage("galang")
	time.Sleep(time.Millisecond * 2000)
	fmt.Println("end")

}

//
func showMessage(msg string) {
	for i := 0; i < 5; i++ {
		fmt.Printf("msg:%v,%v\n", msg, i)
		time.Sleep(time.Millisecond * 100)
	}

}
