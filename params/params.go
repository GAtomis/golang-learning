/*
 * @Description: 请输入....
 * @Author: Gavin
 * @Date: 2022-06-27 11:36:27
 * @LastEditTime: 2022-06-27 14:58:13
 * @LastEditors: Gavin
 */
package params

import "fmt"

func getNameAndAge() (string, int) {
	return "我的队伍", 23
}

func Enter() {
	//结构匿名变量
	name, age := getNameAndAge()
	fmt.Printf("name: %v\n", name)
	fmt.Printf("age: %v\n", age)
	//第一个参数不使用,可以使用
	_, age2 := getNameAndAge()
	fmt.Printf("age: %v\n", age2)
	//常量
	const PI int = 10
	const CLASS_NAME = "类名"
	const (
		a1 = iota
		_
		a2 = iota
	)
	fmt.Printf("a1: %v\n", a1)
	fmt.Printf("a2: %v\n", a2)
}
