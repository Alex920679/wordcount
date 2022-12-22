package main

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"strings"
)

func main() {
	txt, err := readInput()
	if err != nil {
		fail(err)
	}
	txtSl := getSlice(txt)
	fmt.Println(len(txtSl))
}

func readInput() (string, error) {
	var txt string
	flag.StringVar(&txt, "p", "", "txt to match against")
	flag.Parse()
	if txt == "" {
		return txt, errors.New("missing pattern")
	}
	return txt, nil
}

func getSlice(txt string) []string {
	txtSl := strings.Split(txt, " ")
	return txtSl
}

func fail(err error) {
	fmt.Println(err)
	os.Exit(1)
}
