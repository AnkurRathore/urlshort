package main

import (
	"fmt"
	"net/http"

	"github.com/AnkurRathore/urlshort"
)

func main() {

	mux := defaultMux()
	pathsToUrls := map[string]string{
		"/urlshort-godoc": "https://github.com/AnkurRathore/urlshort",
		"/yaml-godoc":     "https://godoc.org/gopkg.in/yaml.v2",
	}
	mapHandler := urlshort.MapHandler(pathsToUrls, mux)

	yaml := `
- path: /urlshort
  url: https://github.com/AnkurRathore/urlshort
- path: /godoc
  url: https://godoc.org/
`
	yamlHandler, err := urlshort.YAMLHandler([]byte(yaml), mapHandler)
	if err != nil {
		panic(err)
	}
	fmt.Println("Starting the server on :9000")
	http.ListenAndServe(":9000", yamlHandler)

}

func defaultMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", hello)
	return mux
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello,world!")
}
