package be

import (
	"fmt"
	"log"
	"net"

	"github.com/MISingularity/logging/p"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

type misBiServer struct {
	biLogs []*p.BiLog
}

func newServer() *misBiServer {
	s := new(misBiServer)
	return s
}

func (s *misBiServer) Bi(ctx context.Context, biLog *p.BiLog) (*p.BiResult, error) {
	writeLog(biLog)
	if err := SaveBiLog(biLog); err != nil {
		return &p.BiResult{false, err.Error()}, err
	}
	return &p.BiResult{true, ""}, nil
}

func (s *misBiServer) BiDeviceInfo(ctx context.Context, deviceInfo *p.DeviceInfo) (*p.BiResult, error) {
	writeDeviceInfo(deviceInfo)

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
