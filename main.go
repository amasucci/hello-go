package main

import (
	"fmt"
	"log"
	"net/http"
)

func helloGo(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	fmt.Fprint(w, "Hello Go!")
}

func healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "All Good")
}

func main() {
	http.HandleFunc("/", helloGo)
	http.HandleFunc("/healthz", healthCheckHandler)
	log.Print("Listening on port 9090")
	log.Fatal(http.ListenAndServe(":9090", nil))
}
