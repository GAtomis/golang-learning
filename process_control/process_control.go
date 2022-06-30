/*
 * @Description: 请输入....
 * @Author: Gavin
 * @Date: 2022-06-29 14:26:10
 * @LastEditTime: 2022-06-30 14:18:35
 * @LastEditors: Gavin
 */
package process_control

import "fmt"

func Run() {
	m1 := createdMap()
	key := "name"
	value := m1[key]
	fmt.Printf("value: %v\n", value)

}
func control() {

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
	fmt.Printf("s2: %v\n", s2)
	var s3 = []int{1, 2, 3, 4, 5}
	//[i,n]i包含 n不包含
	fmt.Printf("s3从0～1: %v\n", s3[:1])
	fmt.Printf("s3从1～2: %v\n", s3[1:2])
	fmt.Printf("s3从2～end: %v\n", s3[2:])

}

//	创建切片
func CreatedSlice() []int {
	s1 := []int{1, 3, 2}
	fmt.Printf("s1: %v\n", cap(s1))
	fmt.Printf("s1: %v\n", len(s1))
	return s1
}

func DelAndAddSlice() {
	//添加
	s1 := []int{1, 2, 3}
	s1 = append(s1, 11)
	s1 = append(s1, 12)
	s1 = append(s1, 13)
	//拓展运算符
	s1 = append(s1, s1[0:3]...)
	fmt.Printf("s1: %v\n", s1)
	//删除
	s1 = delSlice[int](s1, 2)
	fmt.Printf("s1: %v\n", s1)
}

//实现一个删除切片||数组
func delSlice[T any](target []T, index int) []T {
	if index < len(target) {
		return append(target[:index], target[index+1:]...)
	}
	return target

}

//创建map
func createdMap() map[string]string {
	m1 := make(map[string]string)
	m1["name"] = "tom"
	m1["age"] = "20"
	for k, v := range m1 {
		fmt.Printf("%v,%v\n", k, v)

	}
	fmt.Printf("m1: %v\n", m1)

	return m1
}
