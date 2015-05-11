package main

import (
	"log"
	"github.com/MiSingularity/logging/be"
	"net/http"
)

func main() {
	be.InitDbConn()
	go func() {
		log.Println("start http server")
		http.HandleFunc("/bi", be.BiFunc)
		if err := http.ListenAndServe(":8088", nil); err != nil {
			log.Fatal(err)
		}
	}()

	be.InitGrpcServer()
}
