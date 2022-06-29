/*
 * @Description: 请输入....
 * @Author: Gavin
 * @Date: 2022-06-29 14:26:10
 * @LastEditTime: 2022-06-29 17:39:16
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
	} else {

		fmt.Println("奇数")
	}

}
func RunForByList() {
	var arr = [...]int{1, 2, 4}
	for i, v := range arr {
		fmt.Printf("i: %v\n", i)
		fmt.Printf("v: %v\n", v)
	}

}
func RunList() {
	// var a1 [3]int
	// var a2 [3]string
	a3 := [3]int{1, 0, 3}
	fmt.Printf("a3: %v\n", a3)
	a2 := [...]int{1, 2, 3}
	fmt.Printf("a2: %v\n", len(a2))
	a1 := [...]int{0: 1, 3: 5, 5: 7}
	fmt.Printf("a1: %v\n", a1)
	for i := 0; i < len(a1); i++ {
		a1[i] = i * 1000
	}
	fmt.Printf("a1: %v\n", a1)

}

//定义切片 没有固定长度
func RunSlice() {
	//声明一个切片
	var s1 []int
	fmt.Printf("s1: %v\n", s1)
	//通过make声明一个切片
	s2 := make([]int, 2)

}

func createdSlice()int[]{
	s1:=[]int{1,3,2}
		return s1
}
