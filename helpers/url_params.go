package helpers

import (
	"net/http"
	"strconv"
)

func ParseIDFromURL(path, url string, w http.ResponseWriter) (int, error) {
	idStr := path[len(url):] // ambil URL path dari r, trus di cut sepanjang "/api/categories/" dan diambil setelahnya
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return 0, err
	}

	return id, nil
}
