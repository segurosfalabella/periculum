package main

import (
	"fmt"
	"log"

	"github.com/mrsangrin/periculum/service"
)

func main() {
	remoteServiceURL := "http://localhost:3001/health"
	caller := service.Caller{Endpoint: remoteServiceURL}
	response, err := caller.Request()

	if err != nil {
		log.Println("It's not possible fetch some status from api")
	}
	fmt.Printf("Status code %d", response.StatusCode)
}
