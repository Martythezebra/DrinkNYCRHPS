package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("get pogg")

	const PORT_NUMBER = ":3000"

	mainFileServer := http.FileServer(http.Dir("./webby"))

	http.Handle("/", mainFileServer)

	potentialError := http.ListenAndServe(PORT_NUMBER, nil)

	if potentialError != nil {
		log.Fatal(potentialError)
	}
}
