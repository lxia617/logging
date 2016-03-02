package main

import (
	"log"
	"time"

	"github.com/MISingularity/logging/fe"
	"github.com/MISingularity/logging/p"
)

func main() {
	log.Println("init bi client")
	fe.Init("127.0.0.1", "50051")
	log.Println("add a bi log")
	fe.Bi(&p.BiLog{
		ProjectName: "deepshare",
		ActionName:  "userlink",
		Timestamp:   time.Now().Unix(),
		Detail:      []byte("detail~~~~~~"),
	})
	time.Sleep(time.Second)
	fe.Bi(&p.BiLog{
		ProjectName: "ime",
		ActionName:  "input",
		Timestamp:   time.Now().Unix(),
		Detail:      []byte("detail~~~~~~"),
	})
}
