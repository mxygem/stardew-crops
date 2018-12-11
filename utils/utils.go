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
// and returns it as a string
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
