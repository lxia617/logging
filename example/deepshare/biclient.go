package main

import (
	"log"
	"misbi/bi"
	"misbi/p"
	"time"
)

func main() {
	log.Println("init bi client")
	bi.Init("127.0.0.1", "8999")
	log.Println("add a bi log")
	bi.Bi(&p.BiLog{
		ProjectName: "deepshare",
		ActionName:  "userlink",
		Timestamp:   time.Now().Unix(),
		Detail:      []byte("detail~~~~~~"),
	})
}
