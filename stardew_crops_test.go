package main

import (
	"io/ioutil"
	"os"

	"github.com/DATA-DOG/godog"
	"github.com/DATA-DOG/godog/gherkin"
)

var cliOut []byte

func FeatureContext(s *godog.Suite) {
	var so, r, w *os.File

	s.BeforeStep(func(step *gherkin.Step) {
		so, r, w = stdOutUp()
	})

	s.AfterStep(func(step *gherkin.Step, err error) {
		cliOut = stdOutDown(so, r, w)
	})

	s.Step(`^info for (\w+) is requested$`, InfoCommand)
	s.Step(`^the output must match$`, TheOutputMustMatch)
}

func stdOutUp() (so, r, w *os.File) {
	so = os.Stdout
	r, w, _ = os.Pipe()
	os.Stdout = w

	return so, r, w
}

func stdOutDown(so, r, w *os.File) []byte {
	w.Close()
	out, _ := ioutil.ReadAll(r)
	os.Stdout = so

	return out
}
