/**
  @author:panliang
  @data:2022/7/17
  @note
**/
package server

import (
	"google.golang.org/grpc"
	"im-services/config"
	grpcAuth "im-services/server/grpc/auth"
	grpcMessage "im-services/server/grpc/message"
	"log"
	"net"
)

var RpcServer = grpc.NewServer()

func StartGrpc() {
	
	var auth grpcAuth.ImAuthHandlerServer
	var message grpcMessage.ImMessageServer

	grpcAuth.RegisterImAuthHandlerServer(RpcServer, auth)
	grpcMessage.RegisterImMessageServer(RpcServer, message)

	listener, err := net.Listen("tcp", config.Conf.Server.GrpcListen)
	if err != nil {
		log.Fatal("grpc服务启动失败", err)
	}
	_ = RpcServer.Serve(listener)

}
