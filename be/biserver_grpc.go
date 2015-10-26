package be

import (
	"fmt"
	"github.com/MISingularity/logging/p"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"log"
	"net"
)

type misBiServer struct {
	biLogs []*p.BiLog
}

func newServer() *misBiServer {
	s := new(misBiServer)
	return s
}

func (s *misBiServer) Bi(ctx context.Context, biLog *p.BiLog) (*p.BiResult, error) {
	log.Println("[grpc] server api Bi() called", biLog.ProjectName, biLog.ActionName, biLog.Timestamp, string(biLog.Detail))
	if err := SaveBiLog(biLog); err != nil {
		return &p.BiResult{false, err.Error()}, err
	}
	return &p.BiResult{true, ""}, nil
}

func InitGrpcServer(port string) {
	lis, err := net.Listen("tcp", ":"+port)
	if err != nil {
		fmt.Println(err)
		log.Fatal("failed to listen: ", err)
	}
	log.Println("grpc server started. listening on port", port)
	grpcServer := grpc.NewServer()
	p.RegisterMisBiServer(grpcServer, newServer())
	if err := grpcServer.Serve(lis); err != nil {
		fmt.Println(err)
		log.Panic(err)
	}
}
