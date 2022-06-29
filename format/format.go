/*
 * @Description: 请输入....
 * @Author: Gavin
 * @Date: 2022-06-28 17:36:44
 * @LastEditTime: 2022-06-29 11:10:07
 * @LastEditors: Gavin
 */
package format

import "fmt"

type WebSite struct {
	Name string
}

func GetWebSite() {

	site := WebSite{Name: "zhounan"}
	fmt.Printf("site: %v\n", site)
	fmt.Printf("site: %#v\n", site)

}
