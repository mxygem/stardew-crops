package utils

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"runtime"
	"strings"

	"github.com/sergi/go-diff/diffmatchpatch"
)

// Open opens a file based on the provided path
// and returns it as a trimmed string
func Open(openPath string) string {
	fp := getPath(openPath)
	f := open(fp)

	return strings.TrimSpace(string(f))
}

// OpenBytes opens a file and returns its bytes
func OpenBytes(openPath string) []byte {
	_, filename, _, ok := runtime.Caller(1)
	if !ok {
		fmt.Printf("unable to determine caller")
		os.Exit(1)
	}
	fp := path.Join(path.Dir(filename), openPath)

	return open(fp)
}

// STDOutUp sets up for capturing stdout
func STDOutUp() (so, r, w *os.File) {
	so = os.Stdout
	r, w, _ = os.Pipe()
	os.Stdout = w

	return so, r, w
}

// STDOutDown reads the captured stdout and returns it
func STDOutDown(so, r, w *os.File) []byte {
	w.Close()
	out, _ := ioutil.ReadAll(r)
	os.Stdout = so

	return out
}

// Diff returns the diff of two provided strings
func Diff(t1, t2 string) string {
	dmp := diffmatchpatch.New()
	diffs := dmp.DiffMain(t1, t2, false)

	return dmp.DiffPrettyText(diffs)
}

func AssertMatch(expected, actual string) error {
	if expected == actual {
		return nil
	}

	return fmt.Errorf("expected content and found content do not match!\n expected: %s\n found: %s", expected, actual)
}

func getPath(op string) string {
	callerPath := path.Dir(caller(3))

	return path.Join(callerPath, op)
}

func caller(level int) string {
	_, filename, _, ok := runtime.Caller(level)
	if !ok {
		fmt.Printf("unable to determine caller")
		os.Exit(1)
	}

	return filename
}

func open(path string) []byte {
	f, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Printf("unable to load file %s", err.Error())
		os.Exit(1)
	}

	return f
}
