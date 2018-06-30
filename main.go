package main

import (
	"flag"
	"fmt"
	"io"
	"os"

	"github.com/umaumax/goecho"
)

const (
	ENV_PREFIX = "GOCAT_PREFIX"
	ENV_SUFFIX = "GOCAT_SUFFIX"
)

var (
	prefix string
	suffix string
)

func init() {
	flag.StringVar(&prefix, "prefix", "", "prefix string")
	flag.StringVar(&suffix, "suffix", "", "suffix string")
}

func main() {
	flag.Parse()
	args := flag.Args()

	if prefix == "" {
		prefix = os.Getenv(ENV_PREFIX)
	}
	if suffix == "" {
		suffix = os.Getenv(ENV_SUFFIX)
	}
	prefix = goecho.EscapeBackslash(prefix)
	suffix = goecho.EscapeBackslash(suffix)

	exitcode := 0
	var r io.Reader
	var w io.Writer
	w = os.Stdout
	if flag.NArg() == 0 {
		args = append(args, "-")
	}
	for _, filename := range args {
		var f *os.File
		var err error
		if filename == "-" {
			r = os.Stdin
		} else {
			f, err = os.Open(filename)
			if err != nil {
				fmt.Fprintf(os.Stderr, "gocat: %s: No such file or directory\n", filename)
				exitcode = 1
				continue
			}
			defer f.Close()
			r = f
		}
		fmt.Fprint(w, prefix)
		io.Copy(w, r)
		fmt.Fprint(w, suffix)
	}
	os.Exit(exitcode)
}
