/*
 * @Description: 请输入....
 * @Author: Gavin
 * @Date: 2022-07-15 15:40:17
 * @LastEditTime: 2022-07-15 15:53:23
 * @LastEditors: Gavin
 */
package waitgroup

import (
	"fmt"
	"sync"
)

//声明一个wg类型
var wg sync.WaitGroup

func showMsg(i int) {
	fmt.Printf("i: %v\n", i)
	defer wg.Done()
}

func Run() {

	for i := 0; i < 4; i++ {
		go showMsg(i)
		wg.Add(1)

	}
	wg.Wait()
	fmt.Println("结束")

}
