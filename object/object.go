/*
 * @Description: 结构体
 * @Author: Gavin
 * @Date: 2022-07-01 18:13:14
 * @LastEditTime: 2022-07-03 09:51:36
 * @LastEditors: Gavin
 */
package object

import "fmt"

type Person struct {
	name, des string
	age, id   int
}

func Run() {
	getStructPointer()

}

//创建一个结构体
func createdStruct() Person {

	var tom Person

	var ashe struct {
		name string
		age  int
	}
	ashe.name = "爱惜"
	ashe.age = 123

	//方式一
	tom.id = 001
	tom.name = "Tom"
	tom.age = 12
	tom.des = "公司小职员"
	//方式二

	tom = Person{
		id:   10,
		name: "zhounan",
		des:  "GGAA",
		age:  123,
	}

	return tom
}

//结构体指针
func getStructPointer() (Person, *Person) {
	//方式二
	var p_person *Person
	tom := Person{
		id:   10,
		name: "zhounan",
		des:  "GGAA",
		age:  123,
	}
	//赋值内存地址
	p_person = &tom
	fmt.Printf("%v,%p\n", tom, p_person)
	fmt.Printf("p_person类型: %T\n", p_person)
	fmt.Printf("p_person取值: %v\n", *p_person)
	return tom, p_person

}
