package main

import (
	"log"
	
	"github.com/vanbien2402/highest-prime-number/internal/api"
)

func main() {
	svr, err := api.NewServer()
	if err != nil {
		log.Println("init server failed", err)
		panic(err)
	}
	if err = svr.Start(); err != nil {
		log.Println("start server failed", err)
		panic(err)
	}
}
