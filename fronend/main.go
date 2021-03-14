package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strings"
	"time"
)


var tmpl = template.Must(template.New("tmpl").ParseFiles("index.html"))

var url string  = "http://localhost:8000"

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if err := tmpl.ExecuteTemplate(w, "index.html", nil); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

	timeout := time.Duration(5 * time.Second)
	client := http.Client{
		Timeout: timeout,
	}

	var x string = r.FormValue("method")
	var z string = r.FormValue("body")

    request, err := http.NewRequest(x, url, strings.NewReader(z))
	if err != nil {
		log.Fatalln(err)
	}

    resp, err := client.Do(request)
	if err != nil {
		log.Fatalln(err)
	}
	
    defer resp.Body.Close()

	fmt.Fprintln(w, "HTTP METHOD:\n", x)
	fmt.Fprintln(w, "BODY DATA:\n", z)
	})

	log.Fatal(http.ListenAndServe(":9000", nil))
}

func getMethod(w http.ResponseWriter, r *http.Request) {
    if r.Method == "GET" {
        r.ParseForm()
        w.Write([]byte(r.Form["method"][0]))
    }
}