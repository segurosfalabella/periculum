package main

import (
	"log"

	"github.com/mrsangrin/periculum/service"
)

func main() {
	remoteServiceURL := "http://localhost:3001/health"
	service.Call(remoteServiceURL)
	log.Println("here in main")
}
