/*
 * @Description: 属性和方法分离
 * @Author: Gavin
 * @Date: 2022-07-04 16:11:08
 * @LastEditTime: 2022-07-04 17:24:48
 * @LastEditors: Gavin
 */
package method

import "fmt"

type Person struct {
	name string
}

// 函数格式 接受者+方法名
func (per Person) eat() {
	fmt.Println("吃饭")
}
func (per Person) login(name string, password string) bool {
	if name == "zhounan" && password == "123" {
		return true
	}
	return false
}

func Run() {
	per := Person{
		name: "zhounan",
	}
	per.eat()
	fmt.Printf("per.login(\"zhounan\", \"123\"): %v\n", per.login("zhounan", "123"))

}
