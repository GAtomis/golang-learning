/*
 * @Description: 请输入....
 * @Author: Gavin
 * @Date: 2022-10-13 17:40:24
 * @LastEditTime: 2022-10-17 15:07:15
 * @LastEditors: Gavin
 */
package assertion_reflect

import (
	"fmt"
	"reflect"
)

type User struct {
	Name string
	Age  int
	Sex  int
}

type Student struct {
	User
}

func (u User) SayName(name string) {
	fmt.Println("我的名称", name)

}
func Run() {
	u := User{
		Name: "周楠",
		Age:  123,
		Sex:  0,
	}
	checkb(u)

}
func check(v any) {

	switch v.(type) {
	case User:
		fmt.Println(v.(User).Name)
	case Student:
		fmt.Println("我是student")

	}

}

func checkb(inter any) {
	t := reflect.TypeOf(inter)
	v := reflect.ValueOf(inter)
	fmt.Println(t, v)

}
