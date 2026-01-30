package helpers

import (
	"encoding/json"
	"net/http"
)

func EncodeToJSON(w http.ResponseWriter, data any) error {
	w.Header().Set("Content-Type", "application/json")
	return json.NewEncoder(w).Encode(data)
}
