package main

import (
	"html/template"
	"os"
)

type Person struct {
	Name string
}

func main() {
	temp, err := template.ParseFiles("hello.gohtml")
	if err != nil {
		panic(err)
	}

	gyen := Person{Name: "Gyen Abubakar"}

	err = temp.Execute(os.Stdout, gyen)
	if err != nil {
		panic(err)
	}
}
