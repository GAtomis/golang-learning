/*
 * @Description: 请输入....
 * @Author: Gavin
 * @Date: 2022-06-30 14:32:54
 * @LastEditTime: 2022-07-01 15:26:32
 * @LastEditors: Gavin
 */
package function

import "fmt"

func Run() {
	useClosure()
}
func f2(agrs ...int) {
	for _, v := range agrs {
		fmt.Printf("v: %v\n", v)
	}
}
func towRe() (string, int) {
	name := "张三"
	age := 12
	return name, age
}
func towRe2() (name string, age int) {
	name = "张三"
	age = 12
	return name, age
}
func f3(name string, age int, args ...int) {
	fmt.Printf("name: %v\n", name)
	fmt.Printf("age: %v\n", age)
	for _, v := range args {
		fmt.Printf("v: %v\n", v)
	}
}

func useSum(a int, b int) int {

	return a + b
}
func useAdd(a int, b int) int {
	return a + b
}

//函数的类型
func useFuncType() {
	type ft func(int, int) int
	var f1 ft
	f1 = useSum
	f1 = useAdd
	f1(1, 2)

}

//闭包 当函数内的使用有内部函数使用了当前函数的变量,当前函数的变量不会被回收 形成了闭包
func useClosure() {
	f1 := add()
	fmt.Printf("f1: %v\n", f1(10))
	fmt.Printf("f1: %v\n", f1(20))
	println("-----------")
	f1 = add()
	fmt.Printf("f1: %v\n", f1(200))
	fmt.Printf("f1: %v\n", f1(300))

}

type f1 func(int) int

//使用闭包函数
func add() f1 {
	var x int
	return func(a int) int {
		x += a
		return x
	}
}
