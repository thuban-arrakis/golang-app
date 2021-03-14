package main

import (
    "fmt"
    "log"
    "net/http"
    "io/ioutil"
    "os"
)

func main() {
    mux := http.NewServeMux()
    mux.HandleFunc("/", clientResponse)
    

    err := http.ListenAndServe(":8000", mux)
    fmt.Println("starting server...")
    log.Fatal(err)
    log.SetOutput(os.Stderr)
    log.SetOutput(os.Stdout)
}

func clientResponse(w http.ResponseWriter, r *http.Request) {

    w.Header().Set("Access-Control-Allow-Origin", "*")

    b, err := ioutil.ReadAll(r.Body)
    if err != nil {
        return
    }

    fmt.Fprintf(w, "Method: %s, Body: %s", r.Method, b)
}
