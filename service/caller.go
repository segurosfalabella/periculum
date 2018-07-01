package service

import (
	"fmt"
	"log"
	"net/http"
)

// Remote ...
type Remote interface {
	Request()
}

// Caller ...
type Caller struct {
	Endpoint string
}

// Response ...
type Response struct {
	StatusCode int
	Message    string
}

// Request ...
func (c *Caller) Request() (Response, error) {
	res, err := http.Get(c.Endpoint)
	if err != nil {
		return Response{}, err
	}
	return Response{StatusCode: res.StatusCode, Message: res.Status}, err
}

// Call ...
func Call(Enpoint string) {
	caller := Caller{Endpoint: Enpoint}
	response, err := caller.Request()

	if err != nil {
		log.Println("It's not possible fetch some status from api")
	}

	fmt.Printf("Status code %d", response.StatusCode)
}
