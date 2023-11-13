package main

import (
	"fmt"
	"github.com/profemzy/freshGoProject/pkg/handlers"
	"net/http"
)

const portNumber = ":8080"

func main() {

	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	fmt.Println(fmt.Sprintf("Staring application on port %s", portNumber[1:]))
	_ = http.ListenAndServe(portNumber, nil)
}
