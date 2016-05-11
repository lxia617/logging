package fe

import (
	"log"

	"github.com/MISingularity/logging/p"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

var (
	grpcClient p.MisBiClient
)

// Init
// init a grpc client connect to server
func Init(addr string) error {
	conn, err := grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		log.Fatal("fail to dial:", err)
		return err
	}
	//TODO should close conn
	//	defer conn.Close()
	grpcClient = p.NewMisBiClient(conn)
	return nil
}

// Bi
// add a bi log
func Bi(item *p.BiLog) error {
	rslt, err := grpcClient.Bi(context.Background(), item)
	if err != nil {
		log.Println("error when calling Bi:", err)
		return err
	}
	log.Printf("grpc call result:%#v", rslt)
	return nil
}
