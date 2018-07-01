package main

import "github.com/DATA-DOG/godog"

func a(arg1 string) error {
	return godog.ErrPending
}

func iCheckArtifactsHealthness() error {
	return godog.ErrPending
}

func notifyToStatusPage() error {
	return godog.ErrPending
}

func FeatureContext(s *godog.Suite) {
	s.Step(`^A "([^"]*)"$`, a)
	s.Step(`^I check artifact\'s healthness$`, iCheckArtifactsHealthness)
	s.Step(`^notify to status page$`, notifyToStatusPage)
}
