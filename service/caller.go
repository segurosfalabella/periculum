package service

import (
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
