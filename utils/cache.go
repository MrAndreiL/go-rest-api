package utils

import (
	"net/http"
	"strings"
)

func CacheSupport(key string, w http.ResponseWriter, r *http.Request) bool {
	entityTag := `"` + key + `"`
	w.Header().Set("Etag", entityTag)
	w.Header().Set("Cache-Control", "max-age=120")

	if match := r.Header.Get("If-None-Match"); match != "" {
		if strings.Contains(match, entityTag) {
			return true
		}
	}
	return false
}
