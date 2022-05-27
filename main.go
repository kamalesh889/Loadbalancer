package main

import (
	"fmt"
	"net/http"

	"main.go/Router"
)

func main() {
	fmt.Println("Start load balancer")
	go Router.Healthchk()
	mux := Router.Rout()
	http.Handle("/", mux)
	http.ListenAndServe(":8080", nil)
}
