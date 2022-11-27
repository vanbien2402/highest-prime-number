package main

import (
	"log"

	"github.com/vanbien2402/highest-prime-number/internal/api"
)

// @title My-application
// @version 1.0
// @description This is application which takes input number and return to user the highest prime number

// @contact.name Van Bien
// @contact.email vanbien2402@gmail.com

// @host 52.200.150.234:8080

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
