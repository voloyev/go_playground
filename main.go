package main

import (
	"fmt"
	"os"
)

func main() {
	bytesNumber := 300
	file, err := os.Open(os.Args[1])
	check(err)
	bytes := make([]byte, bytesNumber)
	file.Read(bytes)
	fmt.Println(string(bytes))
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
