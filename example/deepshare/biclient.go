package main

import (
	"log"
	"misbi/bi"
)

func main() {
	log.Println("init bi client")
	bi.Init("127.0.0.1", "8999")
	log.Println("add a bi log")
	bi.Bi()
}
