/*
 * @Description: 请输入....
 * @Author: Gavin
 * @Date: 2022-07-15 16:08:01
 * @LastEditTime: 2022-07-15 16:39:55
 * @LastEditors: Gavin
 */
package runtime

import (
	"fmt"
	"runtime"
)

func showMsg() {
	for i := 0; i < 5; i++ {
		fmt.Printf("i: %v\n", i)
	}
}
func Run() {
	go showMsg()

	//直接结束当前携程
	runtime.Goexit()
	//让出当前执行顺序
	runtime.Gosched()

	fmt.Println("end")
}
