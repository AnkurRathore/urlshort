package urlshort

import (
	"net/http"

	yaml "gopkg.in/yaml.v2"
)

// Map handler will return an http.Handlerfunc that
// will attempt to map any paths to their corresponding
// URL. if the path is not provided in the map, than the fallback
// http.Handler will be called instead

func MapHandler(pathToUrls map[string]string, fallback http.Handler) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		path := r.URL.Path
		// if we can match a path
		//  redirect to it
		if dest, ok := pathToUrls[path]; ok {
			http.Redirect(w, r, dest, http.StatusFound)

			return
		}

		// else
		fallback.ServeHTTP(w, r)
	}
}

func YAMLHandler(yamlBytes []byte, fallback http.Handler) (http.HandlerFunc, error) {
	pathUrls, err := parseYaml(yamlBytes)

	if err != nil {
		return nil, err
	}
	pathToUrls := buildMap(pathUrls)

	return MapHandler(pathToUrls, fallback), nil

}

func buildMap(pathUrls []pathUrl) map[string]string {
	pathsToUrls := make(map[string]string)

	for _, pu := range pathUrls {
		pathsToUrls[pu.Path] = pu.URL
	}

	return pathsToUrls
}

func parseYaml(data []byte) ([]pathUrl, error) {
	var pathUrls []pathUrl
	err := yaml.Unmarshal(data, &pathUrls)
	if err != nil {
		return nil, err
	}
	return pathUrls, err
}

type pathUrl struct {
	Path string `yaml::`
	URL  string
}
