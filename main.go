package main

import (
	"fmt"
	"net/http"
)

const portNumber = ":8080"

func main() {

	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)

	fmt.Println(fmt.Sprintf("Staring application on port %s", portNumber[1:]))
	_ = http.ListenAndServe(portNumber, nil)
}
