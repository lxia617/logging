package fe

import (
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"log"
	"github.com/MiSingularity/logging/p"
)

var (
	grpcClient p.MisBiClient
)

// Init
// init a grpc client connect to server
func Init(host string, port string) {
	conn, err := grpc.Dial(host + ":" + port)
	if err != nil {
		log.Fatal("fail to dial:", err)
	}
	//TODO should close conn
	//	defer conn.Close()
	grpcClient = p.NewMisBiClient(conn)

}

// Bi
// add a bi log
func Bi(item *p.BiLog) {
	rslt, err := grpcClient.Bi(context.Background(), item)
	if err != nil {
		log.Println("error when calling Bi:", err)
	}
	log.Printf("grpc call result:%#v", rslt)
}
