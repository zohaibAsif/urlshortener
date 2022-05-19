package controller

import (
	"log"
	"net/http"

	"gopkg.in/yaml.v2"
)

func MapHandler(data map[string]string, fallback http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		path := r.URL.Path
		if dest, ok := data[path]; ok {
			http.Redirect(w, r, dest, http.StatusFound)
			return
		}
		fallback.ServeHTTP(w, r)
	}
}

func YamlHandler(data []byte, fallback http.Handler) (http.HandlerFunc, error) {
	var pathToUrls []pathToUrl

	err := yaml.Unmarshal(data, &pathToUrls)
	if err != nil {
		log.Fatalf("ERROR(yaml.Unmarshal) :: %v", err)
	}

	pathToUrlsMap := make(map[string]string)

	for _, pu := range pathToUrls {
		pathToUrlsMap[pu.Path] = pu.Url
	}

	// fmt.Println(pathToUrlsMap)

	return MapHandler(pathToUrlsMap, fallback), nil
}

type pathToUrl struct {
	Path string `yaml:"path"`
	Url  string `yaml:"url"`
}
