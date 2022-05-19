package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/zohaibAsif/urlShortener/controller"
)

func main() {
	mux := DefaultMux()

	pathToUrls := map[string]string{
		"/urlshort-myfavsong": "https://www.youtube.com/watch?v=9udS0mpi1L4",
		"/yaml-godoc":         "https://godoc.org/gopkg.in/yaml.v2",
	}

	mapHandler := controller.MapHandler(pathToUrls, mux)

	yaml := `
- path: /urlshort
  url: https://github.com/gophercises/urlshort
- path: /urlshort-pixl
  url: https://github.com/zohaibAsif/pixl
`

	yamlHandler, err := controller.YamlHandler([]byte(yaml), mapHandler)
	if err != nil {
		log.Fatalf("ERROR(yamlHandler) :: %v", err)
	}

	fmt.Println("Starting server on :8081")
	err = http.ListenAndServe(":8081", yamlHandler)
	if err != nil {
		log.Fatalf("ERROR(listenAndServe) :: %v", err)
	}
}

func DefaultMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", Hello)
	return mux
}

func Hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello There!")
}
