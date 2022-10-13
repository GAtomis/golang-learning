/*
 * @Description: 请输入....
 * @Author: Gavin
 * @Date: 2022-10-11 15:55:17
 * @LastEditTime: 2022-10-11 17:32:40
 * @LastEditors: Gavin
 */
package main

import (
	"context"
	"fmt"
	hello_grpc "mypro/gRPC/pb"
	"net"

	"google.golang.org/grpc"
)

type server struct {
	hello_grpc.UnimplementedHelloGRPCServer
}

func (s *server) SayHi(ctx context.Context, req *hello_grpc.Req) (res *hello_grpc.Res, err error) {

	fmt.Println(req.GetMessage())
	return &hello_grpc.Res{

		Message: "我是从服务端饭回的GRPC",
	}, nil

}

func main() {
	l, _ := net.Listen("tcp", ":8888")
	s := grpc.NewServer()
	hello_grpc.RegisterHelloGRPCServer(s, &server{})
	fmt.Println("开始创建")
	s.Serve(l)

}
