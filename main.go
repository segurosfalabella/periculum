package main

import (
	"fmt"
	"log"

	"github.com/mrsangrin/periculum/application"
	"github.com/mrsangrin/periculum/service"
)

const filePATH = "apps.yml"

func checkHealthness(app application.Service) {
	caller := service.Caller{Endpoint: app.Endpoint}
	response, err := caller.Request()
	if err != nil {
		log.Println(err)
	}
	fmt.Printf("Status code %d", response.StatusCode)
}

func main() {
	var apps application.RemoteServices
	services := apps.GetApps(filePATH).Services
	for _, app := range services {
		checkHealthness(app)
	}
}
