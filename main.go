package main

import (
	"fmt"
	"net/http"

	"github.com/prakhar-agarwal-byte/go-url-shortner/urlshort"
)

func main() {
	mux := defaultMux()

	pathsToUrls := map[string]string{
		"/go": "https://www.google.com",
		"/fb": "https://www.facebook.com",
	}
	mapHandler := urlshort.MapHandler(pathsToUrls, mux)

	//Build the YAMLHandler using the mapHandler as the fallback
	yaml := `- path: /go
  url: https://www.google.com
- path: /fb
  url: https://www.facebook.com
`

	yamlHandler, err := urlshort.YAMLHandler([]byte(yaml), mapHandler)
	if err != nil {
		panic(err)
	}
	fmt.Println("Starting the server on :8080")
	http.ListenAndServe(":8080", yamlHandler)
}

func defaultMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", hello)
	return mux
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, world!")
}