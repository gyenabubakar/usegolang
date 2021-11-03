package main

import "log"

func HandleError(err error, message string) {
	if err != nil {
		log.Fatal(message)
	}
}
