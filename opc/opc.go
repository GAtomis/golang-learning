/*
 * @Description: 请输入....
 * @Author: Gavin
 * @Date: 2022-07-05 11:35:07
 * @LastEditTime: 2022-07-05 13:41:36
 * @LastEditors: Gavin
 */
package opc

type Pet interface {
	sleep()
	eat()
}

type Cat struct {
	name string
}
type Dog struct {
	name string
}
type Person struct {
}

func (cat Cat) eat() {

}
func (cat Cat) sleep() {

}
func (per Person) care(pet Pet) {
	pet.eat()
}

func Run() {
	per := Person{}
	cat := Cat{
		name: "zhangsan",
	}
	per.care(cat)

}
