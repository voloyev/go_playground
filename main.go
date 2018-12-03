package main

import (
	"fmt"
)

type MyFile interface {
	Read() (string, error)
	Write(content string) (string, error)
}

// TODO rewrite interface for reference too

type defFile struct {
	name    string
	content string
}

func (d defFile) Read() (string, error) {
	return d.content, nil
}

func (d defFile) Write(content string) (string, error) {
	d.name = content
	return "success", nil
}

func getName(ff *defFile) string {
	ff_name, _ := ff.Read()
	return ff_name
}
func main() {
	ff := defFile{"bob", "content"}
	fmt.Println(getName(&ff))

	ff.Write("nooooooo")
	fmt.Println(getName(&ff))
}
