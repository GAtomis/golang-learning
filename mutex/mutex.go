/*
 * @Description: 请输入....
 * @Author: Gavin
 * @Date: 2022-07-15 16:40:41
 * @LastEditTime: 2022-07-15 17:18:39
 * @LastEditors: Gavin
 */
package mutex_test

import (
	"fmt"
	"sync"
	"time"
)

var i int = 100

//等待组
var wt sync.WaitGroup

//声明锁 试着去除锁和加上锁保持代码执行上文协程锁定,从而保证副作用不会受到影响,保持同步
var lock sync.Mutex

func add() {
	lock.Lock()
	i += 1
	fmt.Printf("i++: %v\n", i)
	time.Sleep(time.Microsecond * 10)
	lock.Unlock()
	defer wt.Done()

}

func sub() {
	lock.Lock()
	i -= 1
	fmt.Printf("i--: %v\n", i)
	time.Sleep(time.Microsecond * 2)
	lock.Unlock()
	defer wt.Done()
}
func Run() {
	for i := 0; i < 100; i++ {
		wt.Add(1)
		go add()
		wt.Add(1)
		go sub()
	}
	wt.Wait()
	fmt.Printf("i: %v\n", i)
}
