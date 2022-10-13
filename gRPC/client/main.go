/*
 * @Description: 请输入....
 * @Author: Gavin
 * @Date: 2022-10-11 15:55:17
 * @LastEditTime: 2022-10-11 17:19:57
 * @LastEditors: Gavin
 */
package main

import (
	"context"
	"fmt"
	hello_grpc "mypro/gRPC/pb"

	"google.golang.org/grpc"
)

func main() {
	cc, err := grpc.Dial("localhost:8888", grpc.WithInsecure())
	if err != nil {
		fmt.Println(err)
	}
	defer cc.Close()
	client := hello_grpc.NewHelloGRPCClient(cc)
	req, _ := client.SayHi(context.Background(), &hello_grpc.Req{
		Message: "我从客户端来",
	})
	fmt.Println(req.GetMessage())

}
