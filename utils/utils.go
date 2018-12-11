package utils

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/sergi/go-diff/diffmatchpatch"
)

// Open opens a file based on the provided path
// and returns it as a trimmed string
func Open(path string) string {
	fp, err := filepath.Abs(path)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(fp)

	f, err := ioutil.ReadFile(fp)
	if err != nil {
		fmt.Printf("unable to load file path = %s, Err = %s", path, err.Error())
		os.Exit(1)
	}

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

	f, err := ioutil.ReadFile(fp)
	if err != nil {
		fmt.Printf("unable to load file path = %s, Err = %s", fp, err.Error())
		os.Exit(1)
	}

	return f
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
