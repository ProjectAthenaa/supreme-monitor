package main

import (
	"github.com/ProjectAthenaa/sonic-core/protos/monitor"
	"google.golang.org/grpc"
	"log"
	"net"
	//monitor2 "supreme-monitor/monitor"
	"os"
)



func main(){
	var lis net.Listener

	if os.Getenv("DEBUG") == "1"{
		lis, _ = net.Listen("tcp", ":4000")
	}else
	{
		lis, _ = net.Listen("tcp", ":3000")
	}
	server := grpc.NewServer()

	monitor.RegisterMonitorServer(server, monitor.Server{})
	if err := server.Serve(lis); err != nil{
		log.Fatal(err)
	}

}