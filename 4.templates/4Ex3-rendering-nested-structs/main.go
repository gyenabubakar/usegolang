package main

import (
	"html/template"
	"os"
)

type Dog struct {
	Name string
	Age  int
}

type Person struct {
	Name    string
	Friends []string
	Dog
}

func main() {
	temp, err := template.ParseFiles("hello.gohtml")
	if err != nil {
		panic(err)
	}

	gyen := Person{
		Name: "Gyen Abubakar",
		Friends: []string{
			"Sena",
			"De Graft",
			"Felix",
		},
		Dog: Dog{
			Name: "Honesty",
			Age:  12,
		},
	}

	err = temp.Execute(os.Stdout, gyen)
	if err != nil {
		panic(err)
	}
}
