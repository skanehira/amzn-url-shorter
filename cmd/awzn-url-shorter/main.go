package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"os"
	"strings"
)

var version = "1.1.0"

var (
	errInvalidURL = errors.New("invalid amazon url")
)

func main() {
	name := "amzn-url-shorter"
	fs := flag.NewFlagSet(name, flag.ContinueOnError)
	fs.SetOutput(os.Stderr)
	fs.Usage = func() {
		fs.SetOutput(os.Stdout)
		fmt.Printf(`%[1]s - generate short amazon url

Version: %s

Usage:
 $ %[1]s url
 $ %[1]s < file
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
		urls, err := multishorter()
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}

		for _, url := range urls {
			fmt.Println(url)
		}
		return
	}

	url, err := shorter(args[0])
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	fmt.Println(url)
}

func multishorter() ([]string, error) {
	var urls []string
	var url string

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		url = scanner.Text()
		url, err := shorter(url)
		if err != nil {
			return nil, err
		}
		urls = append(urls, url)
	}

	return urls, nil
}

func shorter(url string) (string, error) {
	idx := strings.Index(url, "dp/")
	if idx == -1 {
		return "", errInvalidURL
	}
	idx += 3
	end := idx + 10

	if len(url) < end {
		return "", errInvalidURL
	}

	return "https://amazon.co.jp/dp/" + url[idx:end], nil
}
