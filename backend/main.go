package main

import (
    "fmt"
    "log"
    "net/http"
    "io/ioutil"
    "os"
    "time"
)

func main() {
    fmt.Println("starting server...")

    mux := http.NewServeMux()
    mux.HandleFunc("/", clientResponse)
    
    err := http.ListenAndServe(":8000", mux)
    log.Fatal(err)    
}

func clientResponse(w http.ResponseWriter, r *http.Request) {

    w.Header().Set("Access-Control-Allow-Origin", "*")

    b, err := ioutil.ReadAll(r.Body)
    if err != nil {
        return
    }

    fmt.Println(time.Now(), "recieved data to handle with method", r.Method)
    
    log.SetOutput(os.Stderr)
    log.SetOutput(os.Stdout)

    fmt.Fprintf(w, "Method: %s, Body: %s", r.Method, b)
}
