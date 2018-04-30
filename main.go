package main

import (
	"io"
	"os"
)

func main() {
	file, err := os.Open(os.Args[1])
	check(err)
	io.Copy(os.Stdout, file)
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
