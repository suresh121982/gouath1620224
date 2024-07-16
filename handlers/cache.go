package handlers

import (
	"io/ioutil"
	"net/http"
)

var cache map[string][]byte

// CacheHandler handles caching of data
func CacheHandler(w http.ResponseWriter, r *http.Request) {
	// Your caching logic here
	// Example: Cache response of an API call
	data := []byte("data to cache")
	cache["cachedData"] = data
	err := ioutil.WriteFile("cache.txt", data, 0644)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

// ExampleHandler handles fetching cached data
func ExampleHandler(w http.ResponseWriter, r *http.Request) {
	// Example: Fetch cached data
	data, ok := cache["cachedData"]
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	w.Write(data)
}
