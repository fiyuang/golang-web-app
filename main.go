package main

import (
	"errors"
	"fmt"
	"net/http"
)

const portNumber = ":9090"

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "This is the homepage")
}

func About(w http.ResponseWriter, r *http.Request) {
	sum := addValues(1, 2)
	_, _ = fmt.Fprintf(w, "The sum of 1 and 2 is %d", sum)
	// fmt.Fprint(w, "This is the about page")
}

func addValues(x, y int) int {
	// var sum int
	// sum = x + y
	// return sum, nil
	return x + y
}

func Divide(w http.ResponseWriter, r *http.Request) {
	f, err := divideValues(100.0, 0.0)
	if err != nil {
		fmt.Fprint(w, "cannot divide by zero")
		return
	}

	fmt.Fprintf(w, fmt.Sprintf("%f divided by %f is %f", 100.0, 0.0, f))
}

func divideValues(x, y float32) (float32, error) {
	if y <= 0 {
		err := errors.New("Cannot divide by zero")
		return 0, err
	}
	result := x / y
	return result, nil
}

func main() {
	// http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	// 	n, err := fmt.Fprintf(w, "Hello, World!")
	// 	if err != nil {
	// 		fmt.Println(err)
	// 	}
	// 	fmt.Println(fmt.Sprintf("Number of bytes written: %d", n))
	// })

	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)
	http.HandleFunc("/divide", Divide)

	fmt.Println(fmt.Sprintf("Server is listening on port %s", portNumber))
	_ = http.ListenAndServe(portNumber, nil)
}
