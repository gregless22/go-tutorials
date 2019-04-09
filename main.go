package main

import (
	"fmt"
	"log"
	"net/http"
)

var (
	pf = fmt.Printf
	pl = fmt.Println
)

func main() {
	handleRequests()
}

func homePage(w http.ResponseWriter, r *http.Request) {
	pl("end point home page")
	fmt.Fprintf(w, "welcomto the home page")
}

func handleRequests() {
	http.Handle("/", isAuthorized(homePage))
	log.Fatal(http.ListenAndServe(":8090", nil))
}

func isAuthorized(endpoint func(http.ResponseWriter, *http.Request)) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		pl("Checking auth header is set")
		if val, ok := r.Header["Authorized"]; ok {
			fmt.Println(val)
			if val[0] != "true" {
				fmt.Println("Not Authorized!!")
				fmt.Fprintf(w, "Not Authorized!!")
			}
		}
		fmt.Println("Header is set! We can serve content!")
		endpoint(w, r)

	})
}
