package main

import (
	"github.com/MiSingularity/logging/be"
	"log"
	"os"
)

const (
	BI_SERVER_PORT = "50051"
)

func main() {
	logfile, err := os.OpenFile("logfile", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Println("Error: ", err)
	}
	log.SetOutput(logfile)
	defer logfile.Close()
	be.InitDbConn()
	be.InitGrpcServer(BI_SERVER_PORT)
}
