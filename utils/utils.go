package utils

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/sergi/go-diff/diffmatchpatch"
)

// Open opens a file based on the provided path
// and returns it as a string
func Open(path string) string {
	f, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Printf("unable to load file path = %s, Err = %s", path, err.Error())
		os.Exit(1)
	}

	return strings.TrimSpace(string(f))
}

func STDOutUp() (so, r, w *os.File) {
	so = os.Stdout
	r, w, _ = os.Pipe()
	os.Stdout = w

	return so, r, w
}

func STDOutDown(so, r, w *os.File) []byte {
	w.Close()
	out, _ := ioutil.ReadAll(r)
	os.Stdout = so

	return out
}

func Diff(t1, t2 string) string {
	dmp := diffmatchpatch.New()
	diffs := dmp.DiffMain(t1, t2, false)

	return dmp.DiffPrettyText(diffs)
}
