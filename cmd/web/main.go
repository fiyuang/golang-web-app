package main

import (
	"fmt"
	"net/http"

	"github.com/fiyuang/golang-web-app/pkg/handlers"
)

const portNumber = ":9090"

func main() {
	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	fmt.Println(fmt.Sprintf("Server is listening on port %s", portNumber))
	_ = http.ListenAndServe(portNumber, nil)
}
