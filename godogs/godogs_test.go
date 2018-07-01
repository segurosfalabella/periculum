package main

import "github.com/DATA-DOG/godog"

func aRemoteService() error {
	return godog.ErrPending
}

func iCheckArtifactsHealthness() error {
	return godog.ErrPending
}

func notifyToStatusPage() error {
	return godog.ErrPending
}

func FeatureContext(s *godog.Suite) {
	s.Step(`^A remote service`, aRemoteService)
	s.Step(`^I check artifact\'s healthness$`, iCheckArtifactsHealthness)
	s.Step(`^notify to status page$`, notifyToStatusPage)
}
