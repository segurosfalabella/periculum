package main

import (
	"errors"

	"github.com/DATA-DOG/godog"
	"github.com/mrsangrin/periculum/godogs/drivers"
	"github.com/mrsangrin/periculum/service"
)

var caller = service.Caller{Endpoint: "http://localhost:3001/health"}

func aRemoteService() error {
	drivers.RunApp()
	return nil
}

func iCheckArtifactsHealthness() error {
	defer drivers.CloseServer()
	request, err := caller.Request()

	if err != nil {
		return err
	}
	if request.StatusCode <= 200 && request.StatusCode >= 300 {
		return errors.New("unhealthy")
	}
	return nil
}

func notifyToStatusPage() error {
	return godog.ErrPending
}

func FeatureContext(s *godog.Suite) {
	s.Step(`^A remote service`, aRemoteService)
	s.Step(`^I check artifact\'s healthness$`, iCheckArtifactsHealthness)
	s.Step(`^notify to status page$`, notifyToStatusPage)
}
