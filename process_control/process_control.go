/*
 * @Description: 请输入....
 * @Author: Gavin
 * @Date: 2022-06-29 14:26:10
 * @LastEditTime: 2022-06-29 15:09:27
 * @LastEditors: Gavin
 */
package process_control

import "fmt"

func Run() {
	a := 1
	b := 2
	if a > b {
		fmt.Printf("a大于b  ")
	} else {
		fmt.Printf("a小于b  ")
	}
	count := 10
	for i := 0; i < count; i++ {
		fmt.Printf("i: %v\n", i)
		if i == 9 {
			fmt.Println("*******")
		}
	}
	x := [...]int{1, 2, 3}
	for _, v := range x {
		fmt.Printf("v: %v\n", v)
	}

}
func ScanNumber() {
	var num int
	fmt.Println("请输入number")
	fmt.Scan(&num)
	if num%2 == 0 {
		fmt.Println("偶数")
	}

}
