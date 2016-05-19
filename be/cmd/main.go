package main

import (
	"flag"
	"log"
	"os"
	"github.com/MISingularity/logging/be"
)

func main() {
	be.SetLogFile()
	fs := flag.NewFlagSet("dslogging", flag.ExitOnError)
	port := fs.String("port", "50051", "Listening port")
	//mh := fs.String("mongo-host", "127.0.0.1", "MongoDB host")
	mh := fs.String("mongo-host", "42.159.133.35", "MongoDB host")
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

	go be.ShowDataInBrowser()

	be.InitGrpcServer(*port)
}
