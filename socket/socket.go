/*
 * @Description: 请输入....
 * @Author: Gavin
 * @Date: 2022-10-10 18:29:01
 * @LastEditTime: 2022-10-11 11:05:38
 * @LastEditors: Gavin
 */
package socket

import (
	"fmt"
	"net"
	"os"
)

func Run() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s ip-addr\n", os.Args[0])
		os.Exit(1)
	}
	name := os.Args[1]
	addr := net.ParseIP(name)
	if addr == nil {
		fmt.Println("Invalid address")
	} else {
		fmt.Println("The address is ", addr.String())
	}
	os.Exit(0)

}
