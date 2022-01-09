package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%s %s %s", r.Method, r.URL, r.Proto)

	for k, v := range r.Header {
		fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
	}

	for k, v := range r.URL.Query() {
		fmt.Fprintf(w, "Query[%q] = %q\n", k, v[0])
	}

	fmt.Fprintf(w, "Host = %q\n", r.URL.Host)
	fmt.Fprintf(w, "Remote addr = %q\n", r.RemoteAddr)

	if err := r.ParseForm; err != nil {
		log.Print(err)
	}

	for k, v := range r.Form {
		fmt.Fprintf(w, "Form[%q] = %q\n", k, v)
	}

}
