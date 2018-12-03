package main

import (
	"fmt"
)

type myFile interface {
	Read() (string, error)
	Write(content *string) (string, error)
}

type defFile struct {
	name    string
	content string
}

func (d defFile) Read() (string, error) {
	return d.content, nil
}

func (d *defFile) Write(content string) (defFile, error) {
	d.name = content
	return *d, nil
}

func getName(ff *defFile) string {
	ffName, _ := ff.Read()
	return ffName
}

func main() {
	ff := defFile{"bob", "content"}
	fmt.Println(getName(&ff))

	ff.Write("nooooooo")
	fmt.Println(getName(&ff))
}
