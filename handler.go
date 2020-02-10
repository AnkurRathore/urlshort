package urlshort

import "net/http"

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
