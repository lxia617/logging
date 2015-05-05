package bi

import (
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"log"
	p "misbi/proto"
	"time"
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
	//	defer conn.Close()
	grpcClient = p.NewMisBiClient(conn)

}

// Bi
// add a bi log
func Bi() {
	rslt, err := grpcClient.Bi(context.Background(), &p.BiLog{
		ProjectName: "deepshare",
		ActionName:  "userlink",
		Timestamp:   time.Now().Unix(),
		Detail:      []byte("detaildetail")},
	)
	if err != nil {
		log.Println("error when calling Bi:", err)
	}
	log.Printf("grpc call result:%#v", rslt)
}
