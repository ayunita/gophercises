package urlshort

import (
	"encoding/json"
	"net/http"

	"github.com/boltdb/bolt"
)

// MapHandler will return an http.HandlerFunc (which also
// implements http.Handler) that will attempt to map any
// paths (keys in the map) to their corresponding URL (values
// that each key in the map points to, in string format).
// If the path is not provided in the map, then the fallback
// http.Handler will be called instead.
func MapHandler(pathsToUrls map[string]string, fallback http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		path := r.URL.Path
		if dest, ok := pathsToUrls[path]; ok {
			http.Redirect(w, r, dest, http.StatusFound)
			return
		}
		fallback.ServeHTTP(w, r)
	}
}

func DBHandler(db *bolt.DB, fallback http.Handler) (http.HandlerFunc, error) {
	pathsToUrls := make(map[string]string)
	var err error

	if err = db.View(func(tx *bolt.Tx) error {
		if err = tx.Bucket([]byte("PATHURLS")).ForEach(func(k, v []byte) error {
			pu, err := decode(v)
			if err != nil {
				return err
			}
			pathsToUrls[pu.Path] = pu.URL
			return nil
		}); err != nil {
			return err
		}
		return nil
	}); err != nil {
		return nil, err
	}

	return MapHandler(pathsToUrls, fallback), nil
}

func decode(v []byte) (*pathUrl, error) {
	var pu *pathUrl
	err := json.Unmarshal(v, &pu)
	if err != nil {
		return nil, err
	}
	return pu, nil
}

type pathUrl struct {
	Path string
	URL  string
}
