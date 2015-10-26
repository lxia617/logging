package main

import (
	"flag"
	"github.com/MISingularity/logging/be"
	"io"
	"log"
	"os"
	"time"
)

const (
	BI_SERVER_PORT = "50051"
)

func main() {
	fs := flag.NewFlagSet("dslogging", flag.ExitOnError)
	mh := fs.String("mongo-host", "127.0.0.1", "MongoDB host")
	mp := fs.String("mongo-port", "27017", "MongoDB port")
	if err := fs.Parse(os.Args[1:]); err != nil {
		log.Fatal(err)
	}
	logfile, err := os.OpenFile("logging_"+time.Now().Format("20060102_150405.log"), os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Println("Error: ", err)
	}
	log.SetOutput(io.MultiWriter(logfile, os.Stdout))
	defer logfile.Close()
	if err := be.InitDbConn(*mh, *mp); err != nil {
		log.Fatal(err)
	}
	be.InitGrpcServer(BI_SERVER_PORT)
}
