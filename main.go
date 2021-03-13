package main

import (
    "fmt"
    "log"
    "net/http"
    "io/ioutil"
)

func main() {
    mux := http.NewServeMux()
    mux.HandleFunc("/", clientResponse)

    err := http.ListenAndServe(":8000", mux)
    fmt.Println("starting server...")
    log.Fatal(err)
}

func clientResponse(w http.ResponseWriter, r *http.Request) {

	b, err := ioutil.ReadAll(r.Body)
    if err != nil {
		return
	}

	fmt.Fprintf(w, "Method: %s, Body: %s", r.Method, b)
}
