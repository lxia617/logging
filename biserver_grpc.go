package misbi

import (
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"log"
	"misbi/p"
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
	log.Println("[grpc] server api Bi() called")
	SaveBiLog(biLog)
	return &p.BiResult{false, "detaildetail"}, nil
}

func InitGrpcServer() {
	lis, err := net.Listen("tcp", ":8999")
	if err != nil {
		log.Fatal("failed to listen: ", err)
	}
	log.Println("grpc server start")
	grpcServer := grpc.NewServer()
	p.RegisterMisBiServer(grpcServer, newServer())
	if err := grpcServer.Serve(lis); err != nil {
		log.Panic(err)
	}
}
