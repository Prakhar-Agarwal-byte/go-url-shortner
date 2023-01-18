package urlshort

import (
	"net/http"

	"gopkg.in/yaml.v3"
)

func MapHandler(pathsToUrls map[string]string, fallback http.Handler) http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request) {
		path := r.URL.Path
		if dest, ok := pathsToUrls[path]; ok {
			http.Redirect(w, r, dest, http.StatusFound)
			return
		}
		fallback.ServeHTTP(w,r)
	}
}

func YAMLHandler(data []byte, fallback http.Handler) (http.HandlerFunc, error) {
	paths, err := parseYAML(data)
	if err != nil {
		return nil, err
	}
	pathsToUrls := map[string]string{}
	for _, path := range paths {
		pathsToUrls[path.Path] = path.Url
	}
	return MapHandler(pathsToUrls, fallback), nil
}

func parseYAML(data []byte) ([]pathURL, error) {
	var paths []pathURL
	err := yaml.Unmarshal(data, &paths)
	if err != nil {
		return nil, err
	}
	return paths, nil
}

type pathURL struct {
	Path string `yaml:"path,omitempty"`
	Url  string `yaml:"url,omitempty"`
}
