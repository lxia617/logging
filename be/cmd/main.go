package main

import (
	"flag"
	"log"
	"os"

	"github.com/MISingularity/logging/be"
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
	if fi, err := os.Stat("log"); os.IsNotExist(err) || !fi.IsDir() {
		os.Mkdir("log", 0755)
	}

	if err := be.InitDbConn(*mh, *mp); err != nil {
		log.Fatal(err)
	}
	be.InitGrpcServer(BI_SERVER_PORT)
}
