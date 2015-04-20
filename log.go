package main

import (
	"fmt"
	"os"
)

func debug(obj ...interface{}) {
	if debugOut {
		fmt.Println(obj...)
	}
}

func debugf(msg string, obj ...interface{}) {
	if debugOut {
		fmt.Printf(msg, obj...)
	}
}

func fatal(obj ...interface{}) {
	fmt.Fprintln(os.Stderr, obj...)
	os.Exit(1)
}
