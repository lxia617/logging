package main

import (
	"log"
	"time"

	"strconv"

	"flag"

	"os"

	"fmt"

	"github.com/MISingularity/logging/fe"
	"github.com/MISingularity/logging/p"
)

func main() {
	log.Println("init bi client")

	fs := flag.NewFlagSet("logging_test", flag.ContinueOnError)
	serverAddr := fs.String("server-addr", "42.159.133.35:50051", "Specify the server url to connect to")
	runs := fs.Int("runs", 1, "How many loops")
	interval := fs.Int("interval", 1, "Interval between two grpc call(in Seconds)")
	if err := fs.Parse(os.Args[1:]); err != nil {
		log.Panic(err)
	}
	log.Println("arguments are:")
	fs.VisitAll(func(flag *flag.Flag) {
		fmt.Println("			", flag.Name, ":", flag.Value)
	})
	if err := fe.Init(*serverAddr); err != nil {
		log.Panic(err)
	}
	for i := 0; i < *runs; i++ {
		log.Println("--------------------", i)
		log.Println("add a bi log for deepshare")
		fe.Bi(&p.BiLog{
			ProjectName: "deepshare",
			ActionName:  "userlink",
			Timestamp:   time.Now().Unix(),
			Detail:      []byte("detail~~~~~~" + strconv.Itoa(i)),
		})
		time.Sleep(time.Second * time.Duration(*interval))
		log.Println("add a bi log for ime")
		fe.Bi(&p.BiLog{
			ProjectName: "ime",
			ActionName:  "input",
			Timestamp:   time.Now().Unix(),
			Detail:      []byte("detail~~~~~~" + strconv.Itoa(i)),
		})
		time.Sleep(time.Second * time.Duration(*interval))
	}

}
