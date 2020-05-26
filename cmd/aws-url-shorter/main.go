package main

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"strings"
)

var version = "1.0.0"

var (
	ErrInvalidURL = errors.New("invalid aws url")
)

func main() {
	name := "aws-url-shorter"
	fs := flag.NewFlagSet(name, flag.ContinueOnError)
	fs.SetOutput(os.Stderr)
	fs.Usage = func() {
		fs.SetOutput(os.Stdout)
		fmt.Printf(`%[1]s - generate short aws url

Version: %s

Usage:
 $ %[1]s url
`, name, version)
	}

	if err := fs.Parse(os.Args[1:]); err != nil {
		if err == flag.ErrHelp {
			return
		}
		os.Exit(1)
	}

	args := fs.Args()
	if len(args) == 0 {
		fs.Usage()
		return
	}

	url, err := shorter(args[0])
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
	fmt.Println(url)
}

func shorter(url string) (string, error) {
	idx := strings.Index(url, "dp/")
	if idx == -1 {
		return "", ErrInvalidURL
	}
	idx += 3
	end := idx + 10

	if len(url) < end {
		return "", ErrInvalidURL
	}

	return "https://amazon.co.jp/dp/" + url[idx:end], nil
}
