/*
 * @Description: 请输入....
 * @Author: Gavin
 * @Date: 2022-06-24 23:43:26
 * @LastEditTime: 2022-06-27 17:42:46
 * @LastEditors: Gavin
 */
package main

import (
	"fmt"
	datatype "mypro/data_type"
)

func main() {
	b := datatype.Enter()
	if !b {
		fmt.Printf("b: %v\n", b)

	} else {

	}
	datatype.GetNumberAndAdvance()
	fmt.Printf("b: %v\n", b)
}
