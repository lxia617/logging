package fe

import (
	"github.com/MISingularity/logging/p"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"log"
)

var (
	grpcClient p.MisBiClient
)

// Init
// init a grpc client connect to server
func Init(host string, port string) error {
	conn, err := grpc.Dial(host + ":" + port)
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
