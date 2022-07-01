/*
 * @Description: 指针
 * @Author: Gavin
 * @Date: 2022-07-01 17:21:15
 * @LastEditTime: 2022-07-01 17:25:34
 * @LastEditors: Gavin
 */
package pointer

import "fmt"

func Run() {
	var ip *int
	fmt.Printf("ip: %T\n", ip)
	fmt.Printf("ip: %v\n", ip)

	//取指针
	var i int = 100
	ip = &i
	//指针地址
	fmt.Printf("ip: %v\n", ip)
	//指针取值
	fmt.Printf("ip: %v\n", *ip)
}
