/*
 * @Description: 结构体
 * @Author: Gavin
 * @Date: 2022-07-01 18:13:14
 * @LastEditTime: 2022-07-04 15:51:49
 * @LastEditors: Gavin
 */
package object

import "fmt"

type Person struct {
	name, des string
	age, id   int
}

func Run() {
	testFuncToPointer()

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

func createdPointerByNew() {

	tom := new(Person)
	//访问结构体可以省略 约等于（*tom.id=23）
	tom.id = 23
	tom.name = "zhounan"
	fmt.Printf("tom: %v\n", *tom)

}

//值传递是拷贝副本传递
func testFuncTovalue() {
	tom := Person{
		id:   23,
		name: "zhou",
		age:  23,
		des:  "F1赛车手",
	}
	fmt.Printf("tom: %v\n", tom)
	showPerson(tom)
	fmt.Printf("tom: %v\n", tom)
}
func testFuncToPointer() {
	tom := Person{
		id:   23,
		name: "zhou",
		age:  23,
		des:  "F1赛车手",
	}
	fmt.Printf("tom: %v\n", tom)
	showPersonPointer(&tom)
	fmt.Printf("tom: %v\n", tom)
}
func showPersonPointer(per *Person) {
	per.id = 123
	per.name = "zhouguanyu"
	fmt.Printf("per: %v\n", per)
}

func showPerson(person Person) {
	person.id = 123
	person.name = "zhouguanyu"
	fmt.Printf("person: %v\n", person)

}
