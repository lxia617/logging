package main

import (
	"log"
	"misbi"
	"net/http"
)

func main() {
	go func() {
		log.Println("start http server")
		http.HandleFunc("/bi", misbi.BiFunc)
		http.HandleFunc("/logs", misbi.GetBiLogsFunc)
		if err := http.ListenAndServe(":8088", nil); err != nil {
			log.Fatal(err)
		}
	}()

	misbi.InitGrpcServer()
}
